IMG="pocketnow-new"
VERSION="latest"

@echo ${IMG}
@echo ${VERSION}

-podman manifest exists localhost/${IMG}:latest && podman manifest rm localhost/${IMG}:latest

podman build --arch=amd64 -t ${IMG}:${VERSION}.amd64 -f Dockerfile
podman build --arch=arm64 -t ${IMG}:${VERSION}.arm64 -f Dockerfile

podman manifest create ${IMG}:${VERSION}
podman manifest add ${IMG}:${VERSION} containers-storage:localhost/${IMG}:${VERSION}.amd64
podman manifest add ${IMG}:${VERSION} containers-storage:localhost/${IMG}:${VERSION}.ar
