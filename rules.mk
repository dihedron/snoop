#
# This value is updated each time a new feature is added
# to the rules.mk targets and build rules file.
#
_RULES_MK_CURRENT_VERSION := 202502241730
ifeq ($(_RULES_MK_MINIMUM_VERSION),)
	_RULES_MK_MINIMUM_VERSION := 0
endif

#
# test if minimum rules.mk version requirement is met
#
ifneq ($(shell test $(_RULES_MK_CURRENT_VERSION) -ge $(_RULES_MK_MINIMUM_VERSION); echo $$?),0)
	@echo "minimum rules.mk version requirement not met (expected at least $(_RULES_MK_MINIMUM_VERSION), got $(_RULES_MK_CURRENT_VERSION))" && exit 1
endif

#
# default application metadata
#
_RULES_MK_VARS_NAME ?= my-app
_RULES_MK_VARS_DESCRIPTION ?= <Provide your description here>
_RULES_MK_VARS_COPYRIGHT ?= <20XX> © <your name>
_RULES_MK_VARS_LICENSE ?= MIT
_RULES_MK_VARS_LICENSE_URL ?= https://opensource.org/license/mit/
_RULES_MK_VARS_VERSION_MAJOR ?= 0
_RULES_MK_VARS_VERSION_MINOR ?= 0
_RULES_MK_VARS_VERSION_PATCH ?= 1
_RULES_MK_VARS_VERSION ?= $(_RULES_MK_VARS_VERSION_MAJOR).$(_RULES_MK_VARS_VERSION_MINOR).$(_RULES_MK_VARS_VERSION_PATCH)
_RULES_MK_VARS_MAINTAINER ?= <your-email>@gmail.com
_RULES_MK_VARS_VENDOR ?= <your-email>@gmail.com
_RULES_MK_VARS_PRODUCER_URL ?= https://github.com/<your-github-username>/
_RULES_MK_VARS_DOWNLOAD_URL ?= $(_RULES_MK_VARS_PRODUCER_URL)$(_RULES_MK_VARS_NAME)
_RULES_MK_VARS_METADATA_PACKAGE ?= $$(grep "module .*" go.mod | sed 's/module //gi')/metadata
_RULES_MK_VARS_DOTENV_VAR_NAME ?= $$(echo $(_RULES_MK_VARS_NAME) | tr '[:lower:]' '[:upper:]' | tr '-' '_')_DOTENV

#
# default feature flag values
#
_RULES_MK_FLAG_TIDY_DEPS ?= 1
_RULES_MK_FLAG_ENABLE_CGO ?= 1
_RULES_MK_FLAG_ENABLE_GOGEN ?= 1
_RULES_MK_FLAG_ENABLE_RACE ?= 1
_RULES_MK_FLAG_STATIC_LINK ?= 0
_RULES_MK_FLAG_ENABLE_NETGO ?= 0
_RULES_MK_FLAG_STRIP_SYMBOLS ?= 0
_RULES_MK_FLAG_STRIP_DBG_INFO ?= 0
_RULES_MK_FLAG_FORCE_DEP_REBUILD ?= 0
_RULES_MK_FLAG_OMIT_VCS_INFO ?= 0

#
# Set this flag to 1 to enable automatic dependency tidying.
#
ifneq ($(_RULES_MK_FLAG_TIDY_DEPS),1)
	_RULES_MK_FLAG_TIDY_DEPS := 0
else # neet to enable CGO
	_RULES_MK_FLAG_TIDY_DEPS := 1
endif

#
# In order to enable race detector, the _RULES_MK_FLAG_ENABLE_RACE
# must be set to 1; any other value disables race detector;
# note that the race detector requires CGO to be enabled.
#
ifneq ($(_RULES_MK_FLAG_ENABLE_RACE),1)
	_RULES_MK_FLAG_ENABLE_RACE := 0
else # neet to enable CGO
	_RULES_MK_FLAG_ENABLE_CGO := 1
endif

#
# In order to enable CGO, the _RULES_MK_FLAG_ENABLE_CGO must be
# set to 1; any other value disables CGO.
#
ifneq ($(_RULES_MK_FLAG_ENABLE_CGO),1)
	_RULES_MK_FLAG_ENABLE_CGO := 0
endif

#
# In order to enable go generate, the _RULES_MK_FLAG_ENABLE_GOGEN
# must be set to 1; any other value disables go generate.
#
ifneq ($(_RULES_MK_FLAG_ENABLE_GOGEN),1)
	_RULES_MK_FLAG_ENABLE_GOGEN := 0
endif

#
# In order to statically link the generated binary against libc
# and other libraries (both with and without CGO, see this
# thread https://github.com/golang/go/issues/26492), set this
# value to 1; any other value will produce dynamically linked
# binaries.
#
ifneq ($(_RULES_MK_FLAG_STATIC_LINK),1)
	_RULES_MK_FLAG_STATIC_LINK := 0
endif

#
# In order to use the pure Go network stack implementation (which
# does not require linking against libc), set this to 1; any other
# value uses the native platform's network stack implementation (and
# requires linking against system C libraries).
#
ifneq ($(_RULES_MK_FLAG_ENABLE_NETGO),1)
	_RULES_MK_FLAG_ENABLE_NETGO := 0
endif

#
# Set this flag to 1 if you want to reduce the executable size by
# stripping all the symbols. You will not be able to run go tool nm
# against the binary.
#
ifneq ($(_RULES_MK_FLAG_STRIP_SYMBOLS),1)
	_RULES_MK_FLAG_STRIP_SYMBOLS := 0
endif

#
# Set this flag to 1 if you want to reduce the executable size by
# stripping all the GDB debug information; you will not be able to
# debug the resulting application.
#
ifneq ($(_RULES_MK_FLAG_STRIP_DBG_INFO),1)
	_RULES_MK_FLAG_STRIP_DBG_INFO := 0
endif

#
# Set this flag to 1 if you want to force the rebuild of all dependencies
# even if they are up-to-date. This can be useful when changing the value
# of CGO, in order to make sure that all object files (.a) are compiled
# with the desired settings.
#
ifneq ($(_RULES_MK_FLAG_FORCE_DEP_REBUILD),1)
	_RULES_MK_FLAG_FORCE_DEP_REBUILD := 0
endif

#
# Set this flag to 1 to omit VCS information from the binary.
#
ifneq ($_RULES_MK_FLAG_OMIT_VCS_INFO), 1)
	_RULES_MK_FLAG_OMIT_VCS_INFO := 0
endif

#
# TARGETS
#

.DEFAULT_GOAL := compile

SHELL := /bin/bash

platforms="$$(go tool dist list)"
module := $$(grep "module .*" go.mod | sed 's/module //gi')
ifeq ($(_RULES_MK_VARS_METADATA_PACKAGE),)
	package := $(module)/commands/version
else
	package := $(_RULES_MK_VARS_METADATA_PACKAGE)
endif

now := $$(date --rfc-3339=seconds)

-include .piped

ifeq ($(piped),1)
black:=
red:=
green:=
yellow:=
blue:=
magenta:=
cyan:=
white:=
bold:=
reset:=
else
black:=\033[30m
red:=\033[31m
green:=\033[32m
yellow:=\033[33m
blue:=\033[34m
magenta:=\033[35m
cyan:=\033[36m
white:=\033[37m
bold:=\033[1m
reset:=\033[0m
endif

#
# Linux x86-64 build settings
#
linux/amd64: GOAMD64 ?= v3

#
# Windows x86-64 build settings
#
windows/amd64: GOAMD64 ?= v3

.PHONY: compile
compile: linux/amd64 ;

.PHONY: release
release: quality compile deb rpm apk

.PHONY: show-build-vars
show-build-vars: ## show actual build variables values
	@echo -e "Build Variables:"
	@echo -e " - _RULES_MK_VARS_NAME             : $(green)$(_RULES_MK_VARS_NAME)$(reset)"
	@echo -e " - _RULES_MK_VARS_DESCRIPTION      : $(green)$(_RULES_MK_VARS_DESCRIPTION)$(reset)"
	@echo -e " - _RULES_MK_VARS_COPYRIGHT        : $(green)$(_RULES_MK_VARS_COPYRIGHT)$(reset)"
	@echo -e " - _RULES_MK_VARS_LICENSE          : $(green)$(_RULES_MK_VARS_LICENSE)$(reset)"
	@echo -e " - _RULES_MK_VARS_LICENSE_URL      : $(green)$(_RULES_MK_VARS_LICENSE_URL)$(reset)"
	@echo -e " - _RULES_MK_VARS_VERSION_MAJOR    : $(green)$(_RULES_MK_VARS_VERSION_MAJOR)$(reset)"
	@echo -e " - _RULES_MK_VARS_VERSION_MINOR    : $(green)$(_RULES_MK_VARS_VERSION_MINOR)$(reset)"
	@echo -e " - _RULES_MK_VARS_VERSION_PATCH    : $(green)$(_RULES_MK_VARS_VERSION_PATCH)$(reset)"
	@echo -e " - _RULES_MK_VARS_VERSION          : $(green)$(_RULES_MK_VARS_VERSION)$(reset)"
	@echo -e " - _RULES_MK_VARS_MAINTAINER       : $(green)$(_RULES_MK_VARS_MAINTAINER)$(reset)"
	@echo -e " - _RULES_MK_VARS_VENDOR           : $(green)$(_RULES_MK_VARS_VENDOR)$(reset)"
	@echo -e " - _RULES_MK_VARS_PRODUCER_URL     : $(green)$(_RULES_MK_VARS_PRODUCER_URL)$(reset)"
	@echo -e " - _RULES_MK_VARS_DOWNLOAD_URL     : $(green)$(_RULES_MK_VARS_DOWNLOAD_URL)$(reset)"
	@echo -e " - _RULES_MK_VARS_METADATA_PACKAGE : $(green)$(_RULES_MK_VARS_METADATA_PACKAGE)$(reset)"
	@echo -e " - _RULES_MK_VARS_DOTENV_VAR_NAME  : $(green)$(_RULES_MK_VARS_DOTENV_VAR_NAME)$(reset)"

%: ## replace % with one or more <goos>/<goarch> combinations, e.g. linux/amd64, to build it
	@[ -t 1 ] && piped=0 || piped=1 ; echo "piped=$${piped}" > .piped
#	@echo ""
	@echo -e "Build Flags:"
ifeq ($(_RULES_MK_FLAG_TIDY_DEPS),1)
	@echo -e " - tidy dependencies               : $(green)enabled$(reset)"
	@go mod tidy
else
	@echo -e " - tidy dependencies               : $(yellow)disabled$(reset)"
endif
ifeq ($(_RULES_MK_FLAG_OMIT_VCS_INFO),1)
	@echo -e " - stamp binary with VCS info      : $(yellow)no$(reset)"
	$(eval cvsflags=-buildvcs=false)
else
	@echo -e " - stamp binary with VCS info      : $(green)yes$(reset)"
endif
ifeq ($(_RULES_MK_FLAG_ENABLE_GOGEN),1)
	@echo -e " - go generate                     : $(green)enabled$(reset)"
	@go generate ./...
else
	@echo -e " - go generate                     : $(yellow)disabled$(reset)"
endif
ifeq ($(_RULES_MK_FLAG_ENABLE_CGO),1)
	@echo -e " - CGO dependencies                : $(green)enabled$(reset)"
else
	@echo -e " - CGO dependencies                : $(yellow)disabled$(reset)"
endif
ifeq ($(_RULES_MK_FLAG_ENABLE_NETGO),1)
	@echo -e " - network stack                   : $(green)pure go$(reset)"
else
	@echo -e " - network stack                   : $(yellow)native$(reset)"
endif
ifeq ($(_RULES_MK_FLAG_STRIP_SYMBOLS),1)
	@echo -e " - strip symbols                   : $(yellow)yes$(reset)"
	$(eval strip_symbols=-s)
else
	@echo -e " - strip symbols                   : $(green)no$(reset)"
endif
ifeq ($(_RULES_MK_FLAG_STRIP_DBG_INFO),1)
	@echo -e " - strip debug info                : $(yellow)yes$(reset)"
	$(eval strip_dbg_info=-w)
else
	@echo -e " - strip debug info                : $(green)no$(reset)"
endif
ifeq ($(_RULES_MK_FLAG_ENABLE_CGO),1)
	$(eval linkmode=-linkmode 'external')
endif
ifeq ($(_RULES_MK_FLAG_STATIC_LINK),1)
	@echo -e " - linking                         : $(green)static$(reset)"
	$(eval static=-extldflags '-static')
ifeq ($(_RULES_MK_FLAG_ENABLE_CGO),1)
	$(eval linkmode=-linkmode 'external')
endif
else
	@echo -e " - linking                         : $(yellow)dynamic$(reset)"
endif
ifeq ($(_RULES_MK_FLAG_FORCE_DEP_REBUILD),1)
	@echo -e " - build cache                     : $(yellow)disabled$(reset)"
	$(eval recompile=-a)
else
	@echo -e " - build cache                     : $(green)enabled$(reset)"
endif
ifeq ($(_RULES_MK_FLAG_ENABLE_RACE),1)
	@echo -e " - race detector                   : $(green)enabled$(reset)"
	$(eval race=-race)
else
	@echo -e " - race detector                   : $(yellow)disabled$(reset)"
endif
	@echo -e " - metadata package                : $(green)$(package)$(reset)"
	@$(MAKE) show-build-vars
	@for platform in "$(platforms)"; do \
		if test "$(@)" = "$$platform"; then \
			echo -e "PLATFORM: $(green)$(@)$(reset)"; \
			echo -e "PACKAGES:"; \
			mkdir -p dist/$(@); \
			GOOS=$(shell echo $(@) | cut -d "/" -f 1) \
			GOARCH=$(shell echo $(@) | cut -d "/" -f 2) \
			GOAMD64=$(GOAMD64) \
			CGO_ENABLED=$(_RULES_MK_FLAG_ENABLE_CGO); \
			go build -v \
			$(cvsflags) \
			$(race) \
			$(recompile) \
			-ldflags="\
			$(strip_dbg_info) \
			$(strip_symbols) \
			$(linkmode) \
			$(static) \
			-X '$(package).Name=$(_RULES_MK_VARS_NAME)' \
			-X '$(package).Description=$(_RULES_MK_VARS_DESCRIPTION)' \
			-X '$(package).Copyright=$(_RULES_MK_VARS_COPYRIGHT)' \
			-X '$(package).License=$(_RULES_MK_VARS_LICENSE)' \
			-X '$(package).LicenseURL=$(_RULES_MK_VARS_LICENSE_URL)' \
			-X '$(package).BuildTime=$(now)' \
			-X '$(package).VersionMajor=$(_RULES_MK_VARS_VERSION_MAJOR)' \
			-X '$(package).VersionMinor=$(_RULES_MK_VARS_VERSION_MINOR)' \
			-X '$(package).VersionPatch=$(_RULES_MK_VARS_VERSION_PATCH)' \
			-X '$(package).Vendor=$(_RULES_MK_VARS_VENDOR)' \
			-X '$(package).Maintainer=$(_RULES_MK_VARS_MAINTAINER)' \
			-X '$(package).RulesMkVersion=$(_RULES_MK_CURRENT_VERSION)' \
			-X '$(package).DotEnvVarName=$(_RULES_MK_VARS_DOTENV_VAR_NAME)'" \
			-o dist/$(@)/ . && echo -e "RESULT: $(green)OK$(reset)" || echo -e "RESULT: $(red)KO$(reset)";\
		fi; \
	done
	@rm -f .piped

.PHONY: quality
quality: ## perform static analysis on the code
	@[ -t 1 ] && piped=0 || piped=1 ; echo "piped=$${piped}" > .piped
	@echo -e "Performing $(green)quality checks$(reset)"
ifeq (, $(shell which govulncheck))
	@go install golang.org/x/vuln/cmd/govulncheck@latest
endif
ifeq (, $(shell which gosec))
	@go install github.com/securego/gosec/v2/cmd/gosec@latest
endif
ifeq (, $(shell which shadow))
	@go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
endif
ifeq (, $(shell which staticcheck))
	@go install honnef.co/go/tools/cmd/staticcheck@latest
endif
	@echo -e "Running $(green)govulncheck$(reset)..."
	@govulncheck -show verbose ./...
	@echo -e "Running $(green)shadow$(reset)..."
	@-shadow ./...
	@echo -e "Running $(green)staticcheck$(reset)..."
	@staticcheck ./...
	@echo -e "Running $(green)gosec$(reset)..."
	@-gosec ./...
	@echo -e "$(green)Quality checks$(reset) done!"
	@rm -f .piped

.PHONY: compress
compress: ## compress all the executables with UPX (good quality)
	@[ -t 1 ] && piped=0 || piped=1 ; echo "piped=$${piped}" > .piped
ifeq (, $(shell which upx))
	@echo -e "Need to $(green)install UPX$(reset) first..."
	@sudo apt install upx
endif
	@for binary in `find dist/ -type f -regex '.*$(_RULES_MK_VARS_NAME)[\.exe]*'`; do \
		upx -9 $$binary; \
	done;
	@rm -f .piped

.PHONY: extra-compress
extra-compress: ## compress all the executables with UPX (best quality, slooow!)
	@[ -t 1 ] && piped=0 || piped=1 ; echo "piped=$${piped}" > .piped
ifeq (, $(shell which upx))
	@echo-e  "Need to $(green)install UPX$(reset) first..."
	@sudo apt install upx
endif
	@for binary in `find dist/ -type f -regex '.*$(_RULES_MK_VARS_NAME)[\.exe]*'`; do \
		upx --brute $$binary; \
	done;
	@rm -f .piped

.PHONY: clean
clean: ## remove all build artifacts
	@[ -t 1 ] && piped=0 || piped=1 ; echo "piped=$${piped}" > .piped
	@echo -e "$(green)Cleaning up$(reset) directory..."
	@rm -rf dist
	@rm -rf fetch/server.key fetch/server.crt
	@rm -f .piped

.PHONY: install
install: ## [deprecated] install to a PREFIX (default: /usr/local/bin)
	@[ -t 1 ] && piped=0 || piped=1 ; echo "piped=$${piped}" > .piped
ifneq ($(shell id -u), 0)
	@echo -e "$(red)You must be root to perform this action.$(reset)"
else
ifneq (x86_64, $(shell uname -m))
	@echo -e "$(red)You must be running on x86_64 Linux to perform this action.$(reset)"
endif
ifeq ($(PREFIX),)
	$(eval PREFIX="/usr/local/bin")
endif
ifeq ($(PLATFORM),)
	$(eval PLATFORM=linux/amd64)
endif
	@echo -e "Installing $(green)$(PLATFORM)/$(_RULES_MK_VARS_NAME)$(reset) to $(PREFIX)/$(_RULES_MK_VARS_NAME)..."
	@cp dist/$(PLATFORM)/$(_RULES_MK_VARS_NAME) $(PREFIX)
	@chmod 755 $(PREFIX)/$(_RULES_MK_VARS_NAME)
endif
	@rm -f .piped

.PHONY: uninstall
uninstall: ## [deprecated] remove from a PREFIX (default: /usr/local/bin)
	@[ -t 1 ] && piped=0 || piped=1 ; echo "piped=$${piped}" > .piped
ifneq ($(shell id -u), 0)
	@echo -e "$(red)You must be root to perform this action.$(reset)"
else
ifneq (x86_64, $(shell uname -m))
	@echo -e "You must be running on x86_64 Linux to perform this action."
endif
ifeq ($(PREFIX),)
	$(eval PREFIX="/usr/local/bin")
endif
	@echo "Uninstalling $(PREFIX)/$(_RULES_MK_VARS_NAME)..."
	@rm -rf $(PREFIX)/$(_RULES_MK_VARS_NAME)
endif
	@rm -f .piped

.PHONY: deb
deb: ## package in DEB format the given PLATFORM (default: linux/amd64)
	@[ -t 1 ] && piped=0 || piped=1 ; echo "piped=$${piped}" > .piped
ifeq (, $(shell which nfpm))
	@echo -e "Need to $(green)install nFPM$(reset) first..."
	@go install github.com/goreleaser/nfpm/v2/cmd/nfpm@latest
endif
ifeq ($(PLATFORM),)
	$(eval PLATFORM=linux/amd64)
endif
	$(eval GOOS=$(shell echo $(PLATFORM) | cut -d '/' -f 1))
	$(eval GOARCH=$(shell echo $(PLATFORM) | cut -d '/' -f 2))
	@echo -e "Creating $(green)DEB$(reset) package for $(green)$(_RULES_MK_VARS_NAME)$(reset) version $(green)$(_RULES_MK_VARS_VERSION)$(reset) (for platform $(green)$(PLATFORM)$(reset))..."
	@NAME=$(_RULES_MK_VARS_NAME) VERSION=$(_RULES_MK_VARS_VERSION) GOOS=$(GOOS) GOARCH=$(GOARCH) PLATFORM=$(PLATFORM) nfpm package --packager deb --target dist/$(PLATFORM)/
	@rm -f .piped
# @echo -e "PLATFORM: $(PLATFORM)"
# @echo -e "GOOS: $(GOOS)"
# @echo -e "GOARCH: $(GOARCH)"
# @echo -e "_RULES_MK_VARS_NAME: $(_RULES_MK_VARS_NAME)"
# @echo -e "_RULES_MK_VARS_VERSION: $(_RULES_MK_VARS_VERSION)"

.PHONY: rpm
rpm: ## package in RPM format the given PLATFORM (default: linux/amd64)
	@[ -t 1 ] && piped=0 || piped=1 ; echo "piped=$${piped}" > .piped
ifeq (, $(shell which nfpm))
	@echo -e "Need to $(green)install nFPM$(reset) first..."
	@go install github.com/goreleaser/nfpm/v2/cmd/nfpm@latest
endif
ifeq ($(PLATFORM),)
	$(eval PLATFORM=linux/amd64)
endif
	$(eval GOOS=$(shell echo $(PLATFORM) | cut -d '/' -f 1))
	$(eval GOARCH=$(shell echo $(PLATFORM) | cut -d '/' -f 2))
	@echo -e "Creating $(green)RPM$(reset) package for $(green)$(_RULES_MK_VARS_NAME)$(reset) version $(green)$(_RULES_MK_VARS_VERSION)$(reset) (for platform $(green)$(PLATFORM)$(reset))..."
	@NAME=$(_RULES_MK_VARS_NAME) VERSION=$(_RULES_MK_VARS_VERSION) GOOS=$(GOOS) GOARCH=$(GOARCH) PLATFORM=$(PLATFORM) nfpm package --packager rpm --target dist/$(PLATFORM)/
	@rm -f .piped

.PHONY: apk
apk: ## package in APK format the given PLATFORM (default: linux/amd64)
	@[ -t 1 ] && piped=0 || piped=1 ; echo "piped=$${piped}" > .piped
ifeq (, $(shell which nfpm))
	@echo -e "Need to $(green)install nFPM$(reset) first..."
	@go install github.com/goreleaser/nfpm/v2/cmd/nfpm@latest
endif
ifeq ($(PLATFORM),)
	$(eval PLATFORM=linux/amd64)
endif
	$(eval GOOS=$(shell echo $(PLATFORM) | cut -d '/' -f 1))
	$(eval GOARCH=$(shell echo $(PLATFORM) | cut -d '/' -f 2))
	@echo -e "Creating $(green)APK$(reset) package for $(green)$(_RULES_MK_VARS_NAME)$(reset) version $(green)$(_RULES_MK_VARS_VERSION)$(reset) (for platform $(green)$(PLATFORM)$(reset))..."
	@NAME=$(_RULES_MK_VARS_NAME) VERSION=$(_RULES_MK_VARS_VERSION) GOOS=$(GOOS) GOARCH=$(GOARCH) PLATFORM=$(PLATFORM) nfpm package --packager apk --target dist/$(PLATFORM)/
	@rm -f .piped

.PHONY: container
container: ## create a Docker container to run containerised builds
	@docker build -t golang-1.23.1-with-tools .

.PHONY: docker-prompt
docker-prompt: ## run a bash in the container to run builds
	$(eval USER=$(shell id -u))
	$(eval GROUP=$(shell id -g))
	@docker run -it \
	--rm \
	--volume /etc/passwd:/etc/passwd:ro \
	--volume /etc/group:/etc/group:ro \
	--volume "$(PWD)":/usr/src/ \
	--user $(USER):$(GROUP) \
	-w /usr/src/ \
	golang-1.23.1-with-tools \
	/bin/bash

.PHONY: help
help: ## show help message
	@echo
	@echo "    +-------------------------------+"
	@echo -e "    | rules.mk version \033[36m$(_RULES_MK_FLAG_CURRENT_VERSION)\033[0m |"
	@echo "    +-------------------------------+"
	@awk 'BEGIN {FS = ":.*##"; printf "\nusage:\n  make \033[36m\033[0m\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: howto
howto: ## show how to use this Makefile in your Golang project
	@[ -t 1 ] && piped=0 || piped=1 ; echo "piped=$${piped}" > .piped
	@echo -e "In order to use this Make rules file, simply create a Makefile"
	@echo -e "in the root of your project, with the following $(red)mandatory$(reset) contents:"
	@echo
	@echo -e "_RULES_MK_VARS_NAME := KoolApp $(green)# replace with the name of your executable$(reset) "
	@echo -e "_RULES_MK_VARS_DESCRIPTION := KoolApp provides a cool way to do things. $(green)# replace with a description of your application$(reset) "
	@echo -e "_RULES_MK_VARS_COPYRIGHT := 2024 © Johanna Doe $(green)# replace with proper year @ your name$(reset) "
	@echo -e "_RULES_MK_VARS_LICENSE := MIT $(green)# replace with a license to your liking...$(reset) "
	@echo -e "_RULES_MK_VARS_LICENSE_URL := https://opensource.org/license/mit/ $(green)# ...and set the URL accordingly$(reset) "
	@echo -e "_RULES_MK_VARS_VERSION_MAJOR := 1 $(green)# replace with the major version$(reset)"
	@echo -e "_RULES_MK_VARS_VERSION_MINOR := 0 $(green)# replace with the minor version$(reset) "
	@echo -e "_RULES_MK_VARS_VERSION_PATCH := 2 $(green)# replace with the patch or revision$(reset)"
	@echo -e '_RULES_MK_VARS_VERSION := $$(_RULES_MK_VARS_VERSION_MAJOR).$$(_RULES_MK_VARS_VERSION_MINOR).$$(_RULES_MK_VARS_VERSION_PATCH) $(green)# leave it like this unless you need to override$(reset) '
	@echo -e "_RULES_MK_VARS_MAINTAINER := johanna.doe@example.com $(green)# replace with the email of the maintainer$(reset) "
	@echo -e "_RULES_MK_VARS_VENDOR := koolsoft@example.com $(green)# replace with the email of the vendor$(reset) "
	@echo -e "_RULES_MK_VARS_PRODUCER_URL := https://github.com/koolsoft/ $(green)# replace with the URL of the software producer$(reset)"
	@echo -e '_RULES_MK_VARS_DOWNLOAD_URL := $$(_RULES_MK_VARS_PRODUCER_URL)$$(_RULES_MK_VARS_NAME) $(green)# leave it like this unless you need to override$(reset)'
	@echo
	@echo -e "include rules.mk $(green)# this is where the Make rules are imported$(reset)"
	@echo
	@echo -e "$(green)$(bold)After$(reset) these lines you can add whatever targets you need."
	@echo -e "Please notice that $(magenta)rules.mk$(reset) will set the default target to $(magenta)linux/amd64$(reset)."
	@rm -f .piped

.PHONY: supported
supported: ## show supported build platforms
	@[ -t 1 ] && piped=0 || piped=1 ; echo "piped=$${piped}" > .piped
	@echo -e "Supported build platforms:"
	@OS=$$(uname -s); \
	OS=$${OS,,}; \
	ARCH=$$(uname -p); \
	if [ "$$ARCH" = "x86_64" ]; then \
		ARCH=amd64; \
	fi; \
	for platform in "$(platforms)"; do \
		if [ "$$OS/$$ARCH" = "$$platform" ]; then \
	 		echo -e " - $(green)$$platform$(reset) (current)"; \
		else \
	 		echo -e " - $$platform"; \
		fi; \
	done
	@rm -f .piped

.PHONY: setup-tools
setup-tools: ## install all necessary tools at the latest version
	@[ -t 1 ] && piped=0 || piped=1 ; echo "piped=$${piped}" > .piped
	@go install golang.org/x/tools/gopls@latest
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@go install github.com/securego/gosec/v2/cmd/gosec@latest
	@go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install github.com/mattn/goreman@latest
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.62.2
	@rm -rf .piped

# this in an internal task used to detect whether the main Make
# instance is running in an interactive shell or redirected/piped
# to file; only the main Make process will run it and create the
# temporary .piped file, which will then be included by child Make
# instances.
.piped:
	@[ -t 1 ] && piped=0 || piped=1 ; echo "piped=$${piped}" > .piped
