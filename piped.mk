#
# This in an internal task used to detect whether the main Make
# instance is running in an interactive shell or redirected/piped
# to file; only the main Make process will run it and create the
# temporary .piped file, which will then be included by child Make
# instances.
#
.PHONY: .piped
.piped:
	@[ -t 1 ] && piped=0 || piped=1 ; echo "piped=$${piped}" > .piped
