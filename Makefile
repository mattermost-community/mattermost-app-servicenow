SITE?=http://localhost:3000
ADDR?=:3000

.PHONY: build
## build: build the executable
build:
	go build -o bin/mattermost-apps-servicenow

.PHONY: run
## run: runs the service
run:
	./bin/mattermost-apps-servicenow ${SITE} ${ADDR}

.PHONY: all
## all: builds and runs the service
all: build run

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
