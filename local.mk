#
# This value is updated each time a new feature is added
# to the local.mk targets and build rules file.
#
_LOCAL_MK_CURRENT_VERSION := 202502241730
ifeq ($(_LOCAL_MK_MINIMUM_VERSION),)
	_LOCAL_MK_MINIMUM_VERSION := 0
endif

#
# test if minimum local.mk version requirement is met
#
ifneq ($(shell test $(_LOCAL_MK_CURRENT_VERSION) -ge $(_LOCAL_MK_MINIMUM_VERSION); echo $$?),0)
	@echo "minimum local.mk version requirement not met (expected at least $(_LOCAL_MK_MINIMUM_VERSION), got $(_LOCAL_MK_CURRENT_VERSION))" && exit 1
endif

#
# Extract application variable values from Makefile global context 
# into local.mk specific variables if available.
#
ifdef _APPLICATION_NAME
	_LOCAL_MK_VARS_NAME ?= $(_APPLICATION_NAME)
endif

#
# default application metadata
#
_LOCAL_MK_VARS_NAME ?= my-app

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
# local-install installs the application to a PREFIX (default: /usr/local/bin).
#
.PHONY: local-install
local-install: ## [deprecated] install to a PREFIX (default: /usr/local/bin)
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
	@echo -e "Installing $(green)$(PLATFORM)/$(_LOCAL_MK_VARS_NAME)$(reset) to $(PREFIX)/$(_LOCAL_MK_VARS_NAME)..."
	@cp dist/$(PLATFORM)/$(_LOCAL_MK_VARS_NAME) $(PREFIX)
	@chmod 755 $(PREFIX)/$(_LOCAL_MK_VARS_NAME)
endif
	@rm -f .piped

#
# local-uninstall removes the application from a PREFIX (default: /usr/local/bin).

.PHONY: local-uninstall
local-uninstall: ## [deprecated] remove from a PREFIX (default: /usr/local/bin)
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
	@echo "Uninstalling $(PREFIX)/$(_LOCAL_MK_VARS_NAME)..."
	@rm -rf $(PREFIX)/$(_LOCAL_MK_VARS_NAME)
endif
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
