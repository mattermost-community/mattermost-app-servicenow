# Service Now App

A ServiceNow app for Mattermost.

This repository is licensed under the [Apache 2.0 License](https://github.com/mattermost/mattermost-plugin-github/blob/master/LICENSE).

## Table of Contents

 - [Admin Guide](#admin-guide)
    - [Setting up](#setting-up)    
 - [User's Guide](#users-guide)
    - [Usage](#usage)    
 - [Development](#development)

## Admin Guide

This guide is intended for Mattermost System Admins setting up the ServiceNow app. For more information about contributing to this plugin, visit the [Development section](#development). For more information about the ServiceNow app, read the [technical guide](../docs/technical_documentation).

### Setting up

1. OAuth must be configured to use ServiceNow. In order to configure ServiceNow, refer to the [ServiceNow documentation](https://docs.servicenow.com/bundle/paris-platform-administration/page/administer/security/task/t_CreateEndpointforExternalClients.html).
2. For a redirect URL use `MATTERMOSTURL/plugins/com.mattermost.apps/apps/com.mattermost.servicenow/oauth2/remote/complete`.
3. In Mattermost, run the command `/servicenow configure oauth` and introduce the required fields.
  - `Instance URL` is the URL for your servicenow instance.
  - `Client ID` is the client ID generated on step 1.
  - `Client Secret` is the client secret generated on step 1.

## User's Guide

This guide is intended for Mattermost users who want information about the app's functionality, and Mattermost users who want to connect their ServiceNow account to Mattermost. Connect your ServiceNow account to Mattermost using: `/com.mattermost.servicenow connect` and follow the instructions provided.

To disconnect the account, run `/com.mattermost.servicenow disconnect`.

### Usage

1. In this version, only elements in the `incident` table on ServiceNow can be created.
2. Tickets can be created either by the post menu item, channel header icon, or slash command.
  - Tickets created from the post menu will populate the short description with the post content.
  - Tickets created by commands will show a confirmation modal before creating the ticket.

## Development

### Local development install

1. Running `make` will build the executable and start the server.
  - A base URL can be added so links are sent based on that url (e.g. `make BASE=http://myurl.com`). Defaults to `http://localhost:3000`.
  - An address can be added for the "ListenAndServe" function (e.g. `make ADDR=:3000`). Defaults to `:3000`.
2. Set up your instance to use the Apps Framework debug commands:
  - Go to **System Console > Environment > Developer**.
  - Set **Enable Testing Commands** to **true**.
  - Set **Enable Developer Mode** to **true**.
  - Select **Save**.
3. Add the manifest to your instance using the following command: `/apps debug-add-manifest --url BASE/manifest`.
4. Run the following slash command in Mattermost: `/apps install --app-id com.mattermost.servicenow`.
5. Use `1234` as the secret key.

### Provision

To provision this PR to AWS run `make dist` to generate the App bundle and then follow the steps [here](https://github.com/mattermost/mattermost-plugin-apps#provisioning).
