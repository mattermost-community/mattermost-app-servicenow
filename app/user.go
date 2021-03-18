package app

import (
	"github.com/mattermost/mattermost-app-servicenow/store"
)

func IsUserConnected(botAccessToken, baseURL, userID, botID string) bool {
	_, found := store.GetToken(botAccessToken, baseURL, botID, userID)
	return found
}
