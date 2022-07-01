
export DOCKER_IMAGE_NAME := "$(shell basename $(PWD)):v1.0.0-rc"



.PHONY: bootstrap
bootstrap:
	@$(shell which go) mod tidy &&\
	$(shell which go) mod vendor

.PHNOY: docker-build
docker-build:
	@$(shell which bash) $(PWD)/local/scripts/build-docker.sh

.PHONY: show-docker-image
show-docker-image:
	@echo $(DOCKER_IMAGE_NAME)
