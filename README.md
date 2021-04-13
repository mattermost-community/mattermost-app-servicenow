# Service Now app for Mattermost

## Local Development Install

1. Running `make` will build the executable and start the server.
  - A base URL can be added so links are sent based on that url (e.g. `make BASE=http://myurl.com`). Defaults to `http://localhost:3000`.
  - An address can be added for the "ListenAndServe" function (e.g. `make ADDR=:3000`). Defaults to `:3000`.
2. Setup your instance to use the Apps Framework debug commands.
3. Add the manifest to your instance by entering the following command `/apps debug-add-manifest --url BASE/manifest`.
4. Run the following command in Mattermost `/apps install --app-id com.mattermost.servicenow`.
5. As secret key, use `1234`.

## Provision

To provision this PR to AWS run `make dist` to generate the App bundle and then follow the steps [here](https://github.com/mattermost/mattermost-plugin-apps#provisioning).

## Configuration

1. OAuth must be configured to use ServiceNow. In order to configure ServiceNow side, refer to [ServiceNow documentation](https://docs.servicenow.com/bundle/paris-platform-administration/page/administer/security/task/t_CreateEndpointforExternalClients.html).
2. For redirect URL please use `MATTERMOSTURL/plugins/com.mattermost.apps/apps/com.mattermost.servicenow/oauth2/remote/complete`.
3. In Mattermost, run the command `/servicenow configure oauth` and introduce the required fields.
  - `Instance URL` is the URL for your servicenow instance.
  - `Client ID` is the client ID generated on step 1.
  - `Client Secret` is the client secret generated on step 1.

## Connection
1. In Mattermost, run the command `/com.mattermost.servicenow connect` and follow the instructions.
2. To disconnect the account, run `/com.mattermost.servicenow disconnect`.

## Usage
1. In this version, only elements in the `incident` table on servicenow can be created.
2. Tickets can be created either by the post menu item, channel header icon, or slash command.
  - Tickets created from post menu will populate the short description with the post content.
  - Tickets created by commands will show a confirmation modal before creating the ticket.
