# ServiceNow app Technical Documentation

## Main functionality

- OAuth connection with one Servicenow instance.
- Creation of tickets on a predefined table (incidents).

## Code overview

- App package:
  - Contains the business logic of the App. Handles OAuth information, transformation between tables and bindings, and checks on user connection.
- Clients package:
  - Contains the external clients used by the app. Right now only the ServiceNow client.
  - It also defines the models used by the client (e.g. Ticket model on ServiceNow side).
- Config package:
  - Handles the configuration and its storage. Mainly covers ServiceNow instance and tables (stored in the KV store) and the OAuth config (stored in the OAuth system).
- Constants package:
  - Stores all the constant values used by the app.
- Routers package:
  - Contain all the routers that accept external input. Right now only covers Mattermost (AWS apps should only receive information from Mattermost).
  - It creates the router and handles all calls, including bindings.
- Utils package:
  - Several util functions, mainly covering HTTP responses writing, path handling for bindings, and others.

## Behavior

When the app is initially installed, all users receive the command binding `connect`. The System Admin receives an additional command binding: `onfig OAuth`. Through `config OAuth` we will be able to add the needed information to perform the OAuth connection with ServiceNow: the link to the instance, the client ID, and the client secret.

The OAuth configuration is stored in two different places:
- The client ID and Secret is stored through the Mattermost OAuth API. 
- The link to the instance is stored in the KV store along with the configuration.

The connect command generates an ephemeral message with the link to start the OAuth process through Mattermost.

The connect call will:
- Create the oauth configuration using the information in the context.
- Fetch the state from the call values and fail in case it doesn't exist.
- Return the AuthCodeURL to Mattermost.

The complete call will:
- Get the code from the call values.
- Create the oauth configuration using the information in the context.
- Generate the token out of the code and the OAuth configuration.
- Store the token directly as the OAuth user information service provided by Mattermost.

Any check on whether the user is connected or not will be made checking whether the OAth user information exists in the call context. Connected users will no longer see the connect command. They will have a `disconnect` command, and bindings on post, channel header, and command for creating tickets.

When a user disconnects, the call sets the OAuth user information to `nil`.

All `create ticket` bindings will open a modal dialog. The channel header binding will open an empty modal, post menu binding will open a modal with “short description” set to the post content, and command binding will pre-populate all fields with the values we set on the command.

When you submit the modal, the ticket will be created on the `incident` table on ServiceNow. If any error occurs, it will display in the modal.

All forms for commands are created on binding request to keep the communication to the app at minimum.

Create ticket bindings are prepared to allow several tables. This has not been tested since we need a proper way to add more tables to the config and submenus on channel header and post menu.
