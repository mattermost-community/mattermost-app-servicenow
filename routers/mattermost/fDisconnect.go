package mattermost

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/appclient"

	"github.com/mattermost/mattermost-app-servicenow/app"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
)

func fDisconnect(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	if !app.IsUserConnected(c.Context) {
		utils.WriteCallErrorResponse(w, "You are not connected yet.")
		return
	}

	client := appclient.AsActingUser(c.Context)

	err := client.StoreOAuth2User(nil)
	if err != nil {
		utils.WriteCallErrorResponse(w, fmt.Sprintf("Cannot disconnect. Error: %v", err))
		return
	}

	utils.WriteCallStandardResponse(w, "You are disconnected from Service Now.")
}

func getDisconnectCall() *apps.Call {
	return &apps.Call{
		Path: string(constants.BindingPathDisconnect),
		Expand: &apps.Expand{
			ActingUserAccessToken: apps.ExpandAll,
			OAuth2User:            apps.ExpandAll,
		},
	}
}
