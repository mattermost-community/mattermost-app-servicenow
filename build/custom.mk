# Include custom targets and environment variables here
default: all

BUILD_DATE = $(shell date -u)
BUILD_HASH = $(shell git rev-parse HEAD)
BUILD_HASH_SHORT = $(shell git rev-parse --short HEAD)
LDFLAGS += -X "github.com/mattermost/mattermost-app-servicenow/function.BuildDate=$(BUILD_DATE)"
LDFLAGS += -X "github.com/mattermost/mattermost-app-servicenow/function.BuildHash=$(BUILD_HASH)"
LDFLAGS += -X "github.com/mattermost/mattermost-app-servicenow/function.BuildHashShort=$(BUILD_HASH_SHORT)"
GO_BUILD_FLAGS += -ldflags '$(LDFLAGS)'
GO_TEST_FLAGS += -ldflags '$(LDFLAGS)'

AWS_BUNDLE_NAME ?= $(PLUGIN_ID)-$(PLUGIN_VERSION)-aws.zip
CLOUD_BUNDLE_NAME ?= bundle.zip

## run: runs the app locally
.PHONY: run
run:
	cd http-server ; $(GO) run $(GO_BUILD_FLAGS) .

## dist-aws: creates the bundle file for AWS Lambda deployments
.PHONY: dist-aws
dist-aws:
	rm -rf dist/aws && mkdir -p dist/aws
	cd aws ; \
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GO_BUILD_FLAGS) -o ../dist/aws/servicenow .
	cp manifest.json dist/aws
	cp -r static dist/aws
	cd dist/aws ; \
		zip -m servicenow.zip servicenow ; \
		zip -rm ../$(AWS_BUNDLE_NAME) manifest.json static servicenow.zip
	rm -r dist/aws

## deploy-aws: deploys the app to AWS, and (re-)installs it to Mattermost using appsctl.
.PHONY: deploy-aws
deploy-aws: dist-aws
	appsctl aws deploy -v dist/$(AWS_BUNDLE_NAME) --install --update

## dist-cloud: creates the bundle file for Mattertmost Cloud deployments, using the apps 0.7.0 manifest format
.PHONY: dist-cloud
dist-cloud:
	rm -rf dist/cloud && mkdir -p dist/cloud
	cd aws ; \
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GO_BUILD_FLAGS) -o ../dist/cloud/servicenow .
	cp manifest-v0.7.0.json dist/cloud/manifest.json
	cp -r static dist/cloud
	cd dist/cloud ; \
		zip -m servicenow.zip servicenow ; \
		zip -rm ../$(CLOUD_BUNDLE_NAME) manifest.json static servicenow.zip
	rm -r dist/cloud

## deploy-plugin: deploys and (re-)installs the app as a plugin, locally using appsctl.
.PHONY: deploy-plugin
deploy-plugin: dist
	appsctl plugin deploy -v dist/$(BUNDLE_NAME) --install
