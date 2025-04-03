#
# This value is updated each time a new feature is added
# to the nfpm.mk targets and build rules file.
#
_NFPM_MK_CURRENT_VERSION := 202502241730
ifeq ($(_NFPM_MK_MINIMUM_VERSION),)
	_NFPM_MK_MINIMUM_VERSION := 0
endif

#
# test if minimum nfpm.mk version requirement is met
#
ifneq ($(shell test $(_NFPM_MK_CURRENT_VERSION) -ge $(_NFPM_MK_MINIMUM_VERSION); echo $$?),0)
	@echo "minimum nfpm.mk version requirement not met (expected at least $(_NFPM_MK_MINIMUM_VERSION), got $(_NFPM_MK_CURRENT_VERSION))" && exit 1
endif

#
# Extract application variable values from Makefile global context 
# into nfpm.mk specific variables if available.
#
ifdef _APPLICATION_NAME
	_NFPM_MK_VARS_NAME ?= $(_APPLICATION_NAME)
endif
ifdef _APPLICATION_VERSION
	_NFPM_MK_VARS_VERSION ?= $(_APPLICATION_VERSION)
endif

#
# default application metadata
#
_NFPM_MK_VARS_NAME ?= my-app
_NFPM_MK_VARS_VERSION ?= 0.0.1

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

.PHONY: nfpm-show-vars
nfpm-show-vars: ## show actual packaging variables values
	@echo -e "Packaging Variables:"
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
	@echo -e "Creating $(green)DEB$(reset) package for $(green)$(_NFPM_MK_VARS_NAME)$(reset) version $(green)$(_NFPM_MK_VARS_VERSION)$(reset) (for platform $(green)$(PLATFORM)$(reset))..."
	@NAME=$(_NFPM_MK_VARS_NAME) VERSION=$(_NFPM_MK_VARS_VERSION) GOOS=$(GOOS) GOARCH=$(GOARCH) PLATFORM=$(PLATFORM) nfpm package --packager deb --target dist/$(PLATFORM)/
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
	@echo -e "Creating $(green)RPM$(reset) package for $(green)$(_NFPM_MK_VARS_NAME)$(reset) version $(green)$(_NFPM_MK_VARS_VERSION)$(reset) (for platform $(green)$(PLATFORM)$(reset))..."
	@NAME=$(_NFPM_MK_VARS_NAME) VERSION=$(_NFPM_MK_VARS_VERSION) GOOS=$(GOOS) GOARCH=$(GOARCH) PLATFORM=$(PLATFORM) nfpm package --packager rpm --target dist/$(PLATFORM)/
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
	@echo -e "Creating $(green)APK$(reset) package for $(green)$(_NFPM_MK_VARS_NAME)$(reset) version $(green)$(_NFPM_MK_VARS_VERSION)$(reset) (for platform $(green)$(PLATFORM)$(reset))..."
	@NAME=$(_NFPM_MK_VARS_NAME) VERSION=$(_NFPM_MK_VARS_VERSION) GOOS=$(GOOS) GOARCH=$(GOARCH) PLATFORM=$(PLATFORM) nfpm package --packager apk --target dist/$(PLATFORM)/
	@rm -f .piped

