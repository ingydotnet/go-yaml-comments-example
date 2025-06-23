# Using the "Makes" Makefile setup - https://github.com/makeplus/makes
M := $(or $(MAKES_REPO_DIR),.cache/makes)
$(shell [ -d $M ] || git clone -q https://github.com/makeplus/makes $M)
include $M/init.mk
include $M/clean.mk
include $M/go.mk
include $M/help.mk
include $M/shell.mk


default:: help

define HELP

Usage:
  make run        # Run the program
  make test       # Run the tests
  make test v=1   # Run the tests with verbose output
  make tidy       # Tidy the Go modules
  make lint       # Run the linter
  make vet        # Run the vet
  make help       # Show this help message

endef
