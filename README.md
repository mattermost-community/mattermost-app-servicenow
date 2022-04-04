# ServiceNow App

A ServiceNow app for Mattermost.

This repository is licensed under the [Apache 2.0 License](https://github.com/mattermost/mattermost-plugin-github/blob/master/LICENSE).

## Table of Contents

 - [For Users](#users-guide)
 - [For Admins](#admin-guide)
 - [For Developers](./docs/technical_documentation.md)

## Admin Guide

This guide is intended for Mattermost System Admins setting up the ServiceNow app. For more information about contributing to this plugin, visit the [Development section](#development). For more information about the ServiceNow app, read the [technical guide](docs/technical_documentation.md).

### Setting up

1. OAuth must be configured to use ServiceNow. In order to configure ServiceNow, refer to the [ServiceNow documentation](https://docs.servicenow.com/bundle/paris-platform-administration/page/administer/security/task/t_CreateEndpointforExternalClients.html).
2. For a redirect URL use `MATTERMOSTURL/plugins/com.mattermost.apps/apps/servicenow/oauth2/remote/complete`.
3. In Mattermost, run the command `/servicenow configure oauth` and introduce the required fields.
  - `Instance URL` is the URL for your ServiceNow instance.
  - `Client ID` is the client ID generated in step 1.
  - `Client Secret` is the client secret generated in step 1.

## User's Guide

This guide is intended for Mattermost users who want information about the app's functionality, and Mattermost users who want to connect their ServiceNow account to Mattermost.

- To **connect** your ServiceNow user account to Mattermost use `/servicenow
  connect` and follow the instructions provided.
- To **disconnect** your account, use `/servicenow disconnect`.
- To **create entries** in the `incident` table in ServiceNow
  - Use Mattermost Actions menu on a post in a channel to prepopulate the ticket
    with text and a link to the post;
  - Or use a channel header icon, or `servicenow` command to create an
    `incident` entry from scratch.


