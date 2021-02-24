BASE?=http://localhost:3000
ADDR?=:3000

.PHONY: all
## all: builds and runs the service
all: run

.PHONY: build
## build: build the executable
build:
	go1.16 build -o dist/mattermost-app-servicenow

.PHONY: run
## run: runs the service
run: build
	LOCAL=true ./dist/mattermost-app-servicenow ${BASE} ${ADDR}

.PHONY: dist
## dist: creates the bundle file
dist: build
	cp manifest.json dist/; cd dist/; zip -qr go-function *; zip -r bundle.zip go-function.zip manifest.json

.PHONY: clean
## clean: deletes all
clean:
	rm -rf dist/

.PHONY: help
## help: prints this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
