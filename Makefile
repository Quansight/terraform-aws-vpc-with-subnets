PROJECTNAME := $(shell basename "$(PWD)")
CUR_DIR := "$(shell pwd)"
UID := $(shell id -u)
GID := $(shell id -g)
ALL_MODS := $(wildcard modules/*)

.PHONY: help
help: Makefile
	@echo
	@echo " Available targets"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: tf-fmt/%
tf-fmt/%:
	cd $(*) && terraform fmt .

.PHONY: tf-fmt
## tf-fmt: Make sure all the terraform files are formatted properly
tf-fmt: $(patsubst %,tf-fmt/%, $(ALL_MODS))
	@#

.PHONY: test
## test: Run terratest tests
test:
	cd test && go test -v -timeout 30m

.PHONY: docs
## Generate documentation
docs: tf-fmt
	docker run --rm --volume "$$(pwd):/terraform-docs" -u $$(id -u) quay.io/terraform-docs/terraform-docs:latest markdown /terraform-docs --output-file README.md --output-mode replace
