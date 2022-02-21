DOCKER ?= docker
BUILDIMAGE = yum-repo
DOCKERFILE = $(CURDIR)/docker/Dockerfile

docker-build:
	$(DOCKER) build \
		--tag $(BUILDIMAGE) \
		--file $(DOCKERFILE) .