# ServiceNow App Technical Documentation

### Functionality

- (admin) Configure OAuth2 app on a ServiceNow instance.
- Connect/disconnect user accounts with OAuth2.
- Create entries in the `incident` table; blank, or from a post in Mattermost
  Channels.

### Overview

- `./aws` contains the `main` used to bundle for AWS (Lambda+S3). The AWS bundle
  can be made with `make dist-aws`.
- `./build` is part of a Mattermost plugin bolierplate, with a `custom.mk` that
  builds other app bundles. 
- `./docs` contains documentation (this!)
- `./function` contains the business logic of the app, that does not depend on the
  packaging. It is essentially the app; the name `function` is used for OpenFaas
  template compatibility.
- `./goapp` contains an experimental utility framework for building simple
  Mattermost apps in Go. It shall be moved to the main apps reporitory soon.
- `./http-server` contains the `main` to run as an HTTP server. It can be executed
  as `[PORT=port_number] [ROOT_URL=root_http(s)_url] make run`.
- `./server` contains the `main` used to bundle as a Mattermost plugin. The plugin
  bundle can be made with `make dist`.
- `./static` contains the app's icon.
- `root.go` contains preloaded root-level data: manifest(s) and ./static files.

### Development Environment


The recommended development environment is a local Mattermost server, and the
2. Set up your instance to use the apps framework debug commands:
  - Go to **System Console > Environment > Developer**.
  - Set **Enable Testing Commands** to **true**.
  - Set **Enable Developer Mode** to **true**.
  - Select **Save**.



#### HTTP mode (recommended)

The recommended development environment is a local Mattermost server, and the
App running in the HTTP mode, also locally.

To run the app, use `[PORT=port_number] [ROOT_URL=root_http(s)_url] make run`
command.

Upon startup, the app will display the URL of the manifest, to use with the
`/apps install http` command in Mattermost.

#### Mattermost plugin mode

You can also develop with the app running as a plugin. Use `make dist-plugin` to
build, and `appsctl plugin deploy --install` to (re-)install the App to the
Mattermost server.

For quick iterations when re-installing the app is not required, use `make
deploy`.