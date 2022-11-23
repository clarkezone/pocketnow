#Stage 1 build and test
#docker.io prefix required by podman
# use podman build . --build-arg BUILD_VERSION="jikjikjik" --build-arg BUILD_HASH="0001100"
FROM docker.io/golang:alpine as builder
RUN mkdir /build
WORKDIR /build
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
run apk --no-cache add gcc build-base git
run go build

# generate clean, final image for end users
FROM alpine:3.11.3
COPY --from=builder /build/geo .

# executable
ENTRYPOINT [ "./geo" ]
