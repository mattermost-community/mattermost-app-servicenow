# Servicenow app technical document

## Main functionality:
- OAuth connection with one Servicenow Instance
- Creation of tickets on a predefined table (incidents)

## Code overview:
- App package:
  - Contains the business logic of the App. Handles OAuth information, transformation between tables and bindings and checks on user connection.
- Clients package:
  - Contains the external clients used by the app. Right now only the servicenow client.
  - It also defines the models used by the client (e.g. Ticket model on Servicenow side).
- Config package:
  - Handles the configuration and its storage. Mainly covers Servicenow instance and tables (stored in the KV store) and the OAuth config (stored in the OAuth system).
- Constants package:
  - Stores all the constant values used by the app.
- Routers package:
  - Contain all the routers that accept external input. Right now only covers mattermost (AWS apps should only receive information from Mattermost)
  - It creates the router and handles all calls, including bindings.
- Utils package:
  - Several util functions, mainly covering HTTP responses writing, path handling for bindings and others.

## Behaviour:

When the app is initially installed, all users will receive the command binding “connect”. The System Administrator will also receive the command binding “config oauth”. Through config OAuth we will be able to add the needed information to perform the OAuth connection with Servicenow: the link to the instance, the client ID and the client secret.

The OAuth configuration will be stored in two different places. The client ID and Secret will be stored through the Mattermost OAuth API. The link to the instance will be stored in the KV store along with the configuration.

The connect command will just generate a ephemeral message with the link to start the OAuth process through Mattermost.

The connect call will:
- Create the oauth configuration using the information in the context
- Fetch the state from the call values and fail in case it does not exist.
- Return the AuthCodeURL to Mattermost

The complete call will:
- Get the code from the call values
- Create the oauth configuration using the information in the context
- Generate the token out of the code and the oauth configuration
- Store the token directly as the oauth user information service provided by mattermost

Any check on whether the user is connected or not will be made checking whether the oauth user information exists in the call context.

For connected users, they will no longer see the connect command. They will have a disconnect command, and bindings on post, channel header and command for creating tickets.

When a user disconnects, the call just set the oauth user information to `nil`.

All “create ticket” bindings will open a modal dialog. Channel header binding will open an empty modal, post menu binding will open a modal with “short description” set to the post content, and command binding will pre-populate all fields with the values we set on the command.

When you submit the modal, the ticket will be created on the “incident” table on ServiceNow. If any error happens, it will show the error in the modal.

All forms for commands are created on binding request to keep the communication to the app at minimum.

Create ticket bindings are prepared to allow several tables. This has not been tested since we need a proper way to add more tables to the config and submenus on channel header and post menu.
