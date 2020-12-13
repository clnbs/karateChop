# note: call scripts from /scripts
.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

all: recursive ## Compile all implementation of binary chop

testing: ## Start all static test for this project and create a coverage file in HTML
	bash scripts/test.sh

recursive: ## Compile Recursive implementation of binary chop
	bash scripts/build.sh recursive