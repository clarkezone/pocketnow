# podman can't build dotnet7 images?
.PHONY: buildimage
buildimage:
	docker build . -t pocketnow:latest && docker tag pocketnow:latest registry.dev.clarkezone.dev/pocketnow:latest

.PHONY: pushimage
pushimage:
	docker push registry.dev.clarkezone.dev/pocketnow:latest
