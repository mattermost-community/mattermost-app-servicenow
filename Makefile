BASE?=http://localhost:3000
ADDR?=:3000

GO ?= go
GO_TEST_FLAGS ?= -race

.PHONY: all
## all: builds and runs the service
all: run

.PHONY: build-linux
## build-linux: build the executable for linux
build-linux:
	GOOS=linux GOARCH=amd64 $(GO) build -o dist/mattermost-app-servicenow

.PHONY: build
## build: build the executable
build:
	$(GO) build -o dist/mattermost-app-servicenow

.PHONY: run
## run: runs the service
run: build
	LOCAL=true ./dist/mattermost-app-servicenow ${BASE} ${ADDR} -v

.PHONY: test
## test: tests all packages
test:
	$(GO) test $(GO_TEST_FLAGS) ./...

.PHONY: lint
## lint: Run golangci-lint on codebase
lint: 
	@if ! [ -x "$$(command -v golangci-lint)" ]; then \
		echo "golangci-lint is not installed. Please see https://github.com/golangci/golangci-lint#install for installation instructions."; \
		exit 1; \
	fi; \

	@echo Running golangci-lint
	golangci-lint run ./...

.PHONY: dist
## dist: creates the bundle file
dist: build-linux
	cp -r static dist; cp manifest.json dist/; cd dist/; zip -qr go-function mattermost-app-servicenow; zip -r bundle.zip go-function.zip manifest.json static/

.PHONY: clean
## clean: deletes all
clean:
	rm -rf dist/

.PHONY: help
## help: prints this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
