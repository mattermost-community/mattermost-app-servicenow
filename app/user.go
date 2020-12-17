package app

import "github.com/mattermost/mattermost-app-servicenow/store"

func IsUserConnected(userID string) bool {
	_, found := store.GetToken(userID)
	return found
}
