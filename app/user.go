package app

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func IsUserConnected(cc *apps.Context) bool {
	token := GetTokenFromContext(cc)
	return token != nil
}
