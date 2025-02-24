_RULES_MK_MINIMUM_VERSION=202502241730

_RULES_MK_VARS_NAME := snoop
_RULES_MK_VARS_DESCRIPTION := OpenStack internal event sniffer.
_RULES_MK_VARS_COPYRIGHT := 2025 © Andrea Funtò
_RULES_MK_VARS_LICENSE := MIT
_RULES_MK_VARS_LICENSE_URL := https://opensource.org/license/mit/
_RULES_MK_VARS_VERSION_MAJOR := 0
_RULES_MK_VARS_VERSION_MINOR := 0
_RULES_MK_VARS_VERSION_PATCH := 1
_RULES_MK_VARS_VERSION=$(_RULES_MK_VARS_VERSION_MAJOR).$(_RULES_MK_VARS_VERSION_MINOR).$(_RULES_MK_VARS_VERSION_PATCH)
_RULES_MK_VARS_MAINTAINER=dihedron.dev@gmail.com
_RULES_MK_VARS_VENDOR=dihedron.dev@gmail.com
_RULES_MK_VARS_PRODUCER_URL=https://github.com/dihedron/
_RULES_MK_VARS_DOWNLOAD_URL=$(_RULES_MK_VARS_PRODUCER_URL)$(_RULES_MK_VARS_NAME)
_RULES_MK_VARS_METADATA_PACKAGE=$$(grep "module .*" go.mod | sed 's/module //gi')/metadata
#_RULES_MK_VARS_DOTENV_VAR_NAME=

_RULES_MK_FLAG_ENABLE_CGO=0
_RULES_MK_FLAG_ENABLE_GOGEN=0
_RULES_MK_FLAG_ENABLE_RACE=0
#_RULES_MK_FLAG_STATIC_LINK=1
#_RULES_MK_FLAG_ENABLE_NETGO=1
#_RULES_MK_FLAG_STRIP_SYMBOLS=1
#_RULES_MK_FLAG_STRIP_DBG_INFO=1
#_RULES_MK_FLAG_FORCE_DEP_REBUILD=1
#_RULES_MK_FLAG_OMIT_VCS_INFO=1

include rules.mk

.PHONY: clean-cache ## remove all cached build entries
clean-cache:
	@go clean -x -cache
