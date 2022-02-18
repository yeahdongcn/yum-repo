ifeq ($(shell which git),)
  $(warning "git not found, no commit hash will be bundled into binaries")
  GIT_HASH :=
  GIT_STATE :=
  GIT_TAG :=
  VERSION_SUFFIX := unknown
else
  GIT_HASH := $(shell git rev-parse --short HEAD)
  GIT_STATE := $(shell test -n "`git status --porcelain --untracked-files=no`" && echo "dirty" || echo "clean")
  GIT_TAG := $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
  ifeq ($(GIT_STATE),dirty)
    VERSION_SUFFIX := $(GIT_HASH).dirty
  else
    VERSION_SUFFIX := $(GIT_HASH)
  endif
endif

VERSION := $(shell cat VERSION)
IMAGE_VERSION := $(VERSION)-$(VERSION_SUFFIX)
VERSION_LDFLAGS = -X yum-repo/version.Version=$(VERSION)
VERSION_LDFLAGS += -X yum-repo/version.GitHash=$(GIT_HASH)
VERSION_LDFLAGS += -X yum-repo/version.GitState=$(GIT_STATE)
VERSION_LDFLAGS += -X yum-repo/version.ReleaseStatus=$(RELEASE_STATUS)

.PHONY: version
version:
	@echo "VERSION: $(VERSION)"
	@echo "GIT_HASH: $(GIT_HASH)"
	@echo "GIT_STATE: $(GIT_STATE)"
	@echo "IMAGE_VERSION: $(IMAGE_VERSION)"
