package mattermost

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-servicenow/app"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
)

func fConnect(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	if app.IsUserConnected(c.Context) {
		utils.WriteCallStandardResponse(w, "You are already connected.")
		return
	}

	utils.WriteCallStandardResponse(w, fmt.Sprintf("Follow this link to connect: [link](%s)", c.Context.OAuth2.ConnectURL))
}

func getConnectCall() *apps.Call {
	return &apps.Call{
		Path: string(constants.BindingPathConnect),
		Expand: &apps.Expand{
			App:        apps.ExpandAll,
			OAuth2App:  apps.ExpandAll,
			OAuth2User: apps.ExpandAll,
		},
	}
}
