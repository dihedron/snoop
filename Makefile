_RULES_MK_MINIMUM_VERSION=202502241730

_APPLICATION_NAME := snoop
_APPLICATION_DESCRIPTION := OpenStack internal event sniffer.
_APPLICATION_COPYRIGHT := 2025 © Andrea Funtò
_APPLICATION_LICENSE := MIT
_APPLICATION_LICENSE_URL := https://opensource.org/license/mit/
_APPLICATION_VERSION_MAJOR := 0
_APPLICATION_VERSION_MINOR := 0
_APPLICATION_VERSION_PATCH := 1
_APPLICATION_VERSION=$(_APPLICATION_VERSION_MAJOR).$(_APPLICATION_VERSION_MINOR).$(_APPLICATION_VERSION_PATCH)
_APPLICATION_MAINTAINER=dihedron.dev@gmail.com
_APPLICATION_VENDOR=dihedron.dev@gmail.com
_APPLICATION_PRODUCER_URL=https://github.com/dihedron/
_APPLICATION_DOWNLOAD_URL=$(_APPLICATION_PRODUCER_URL)$(_APPLICATION_NAME)
_APPLICATION_METADATA_PACKAGE=$$(grep "module .*" go.mod | sed 's/module //gi')/metadata
#_APPLICATION_DOTENV_VAR_NAME=

_RULES_MK_FLAG_ENABLE_CGO=0
_RULES_MK_FLAG_ENABLE_GOGEN=0
_RULES_MK_FLAG_ENABLE_RACE=0
#_RULES_MK_FLAG_STATIC_LINK=1
#_RULES_MK_FLAG_ENABLE_NETGO=1
#_RULES_MK_FLAG_STRIP_SYMBOLS=1
#_RULES_MK_FLAG_STRIP_DBG_INFO=1
#_RULES_MK_FLAG_FORCE_DEP_REBUILD=1
#_RULES_MK_FLAG_OMIT_VCS_INFO=1

include golang.mk
include nfpm.mk
include help.mk
include piped.mk

.PHONY: clean-cache ## remove all cached build entries
clean-cache:
	@go clean -x -cache

.PHONY: test
test:
	go test ./...