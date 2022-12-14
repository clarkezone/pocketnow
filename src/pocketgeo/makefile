MAIN_BRANCH := main
HEAD_BRANCH := HEAD
ifeq ($(strip $(VERSION_HASH)),)
# hash of current commit
VERSION_HASH := $(shell git rev-parse --short HEAD)
# tag matching current commit or empty
HEAD_TAG := $(shell git tag --points-at HEAD)
#name of branch
BRANCH_NAME := $(shell git rev-parse --abbrev-ref HEAD)
endif

VERSION_STRING := $(BRANCH_NAME)
#if we are on main and there is a tag pointing at head, use that for version else use branch name as version
ifeq ($(BRANCH_NAME),$(MAIN_BRANCH))
$(info "match main")
ifneq ($(strip $(HEAD_TAG)),)
VERSION_STRING := $(HEAD_TAG)
$(info    $(VERSION_STRING))
endif
endif

#if we are on HEAD and there is a tag pointing at head, use that for version else use branch name as version
ifeq ($(BRANCH_NAME),$(HEAD_BRANCH))
$(info match head)
ifneq ($(strip $(HEAD_TAG)),)
VERSION_STRING := $(HEAD_TAG)
$(info    $(version_string))
endif
endif


BINDIR    := $(CURDIR)/bin
PLATFORMS := linux/amd64/rk-Linux-x86_64 darwin/amd64/rk-Darwin-x86_64 windows/amd64/rk.exe linux/arm64/rk-Linux-arm64 darwin/arm64/rk-Darwin-arm64
BUILDCOMMANDDEBUG := go build -gcflags "all=-N -l" -tags "osusergo netgo static_build"
BUILDCOMMAND := go build -trimpath -ldflags "-s -w -X github.com/clarkezone/geocache/pkg/config.VersionHash=${VERSION_HASH} -X github.com/clarkezone/geocache/pkg/config.VersionString=${VERSION_STRING}" -tags "osusergo netgo static_build"
temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))
label = $(word 3, $(temp))

UNAME := $(shell uname)
ifeq ($(UNAME), Darwin)
SHACOMMAND := shasum -a 256
else
SHACOMMAND := sha256sum
endif

.DEFAULT_GOAL := build

install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
	go install github.com/uw-labs/strongbox@latest
	go install github.com/mgechev/revive@latest

# export $(cat ~/.geocache.env | xargs)
.PHONY: test
test:
	go test -p 4 -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: integration
integration:
	go test -tags="integration" --count=1 -v -timeout 15m ./pkg/kubelayer/...
#	go test -tags="integration" --count=1 -v -timeout 15m ./pkg/jobmanager/...

# This requires
.PHONY: e2etest
e2etest:
# TODO these need to wait before they can be used in here
	go test -tags="integration" --count=1 -v -timeout 15m ./cmd/ -run TestSetupEnvironment
	go test -tags="integration" --count=1 -v -timeout 15m ./cmd/ -run TestCreateJobForClone
	go test -tags="integration" --count=1 -v -timeout 15m ./cmd/ -run TestFullE2eTestWithWebhook

.PHONY: dep
dep:
	go mod tidy

.PHONY: latest
latest:
	echo ${VERSION_STRING} > bin/latest

.PHONY: lint
lint:
	revive $(shell go list ./...)
	go vet $(shell go list ./...)
	golangci-lint run

.PHONY: manifestexists
manifestexists:
	$(eval IMG := "pocketgeotemp")
	$(eval VERSION := "latest")

	@echo ${IMG}
	@echo ${VERSION}

# - means continue on error if command returns 0 it means manifest exists hence delete it
	-podman manifest exists localhost/pocketgeotemp:latest && podman manifest rm localhost/pocketgeotemp:latest

.PHONY: buildimage
buildimage:
	$(eval IMG := "pocketgeotemp")
	$(eval VERSION := "latest")

	@echo ${IMG}
	@echo ${VERSION}

	-podman manifest exists localhost/$(IMG):latest && podman manifest rm localhost/pocketgeotemp:latest

	podman build --arch=amd64 -t ${IMG}:${VERSION}.amd64 -f Dockerfile
	podman build --arch=arm64 -t ${IMG}:${VERSION}.arm64 -f Dockerfile

	podman manifest create ${IMG}:${VERSION}
	podman manifest add ${IMG}:${VERSION} containers-storage:localhost/${IMG}:${VERSION}.amd64
	podman manifest add ${IMG}:${VERSION} containers-storage:localhost/${IMG}:${VERSION}.arm64

.PHONY: pushimage
pushimage:
	$(eval IMG := "pocketgeotemp")
	$(eval VERSION := "latest")

	@echo ${IMG}
	@echo ${VERSION}

	podman manifest push ${IMG}:${VERSION} docker://registry.dev.clarkezone.dev/${IMG}:${VERSION}

.PHONY: precommit
precommit:
	pre-commit run --all-files

.PHONY: build
build:
	$(BUILDCOMMAND) -o ${BINDIR}/pocketgeo

.PHONY: builddlv
builddlv:
	$(BUILDCOMMANDDEBUG) -o ${BINDIR}/pocketgeo

.PHONY: release
build-all: $(PLATFORMS)

$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) CGO_ENABLED=0 $(BUILDCOMMAND) -o "bin/$(label)"
	$(SHACOMMAND) "bin/$(label)" > "bin/$(label).sha256"
