module github.com/mattermost/mattermost-app-servicenow

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/mattermost/mattermost-plugin-api v0.0.21
	github.com/mattermost/mattermost-plugin-apps v0.7.1-0.20220214174025-5e0b38769475
	github.com/mattermost/mattermost-server/v6 v6.0.0-20210901153517-42e75fad4dae
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	go.uber.org/zap v1.17.0
	golang.org/x/oauth2 v0.0.0-20210402161424-2e8d93401602
	google.golang.org/api v0.44.0
)

replace github.com/mattermost/mattermost-plugin-apps => ../mattermost-plugin-apps
