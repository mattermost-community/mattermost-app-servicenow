# Service Now app for Mattermost

## Install

1. Run `go build` on the project folder.
2. Run the generated executable (i.e. `./mattermost-app-servicenow`). This will start the server.
  - You can add your base URL so links are sent based on that url (e.g. `./mattermost-app servicenow http://myurl.com`)
3. Run the following command in Mattermost `/apps install --url http://localhost:3000/manifest`.
  - If you set a custom base URL on step 2, you can run the install with that URL. (e.g. `/app install --url http://myurl.com/manifest`)
4. As secret key, use `1234`.
