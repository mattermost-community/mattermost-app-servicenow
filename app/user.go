package app

import (
	"github.com/mattermost/mattermost-app-servicenow/store"
)

func IsUserConnected(botAccessToken, baseURL, userID string) bool {
	_, found := store.GetToken(botAccessToken, baseURL, userID)
	return found
}
