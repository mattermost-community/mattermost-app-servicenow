package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-servicenow/app"
	"github.com/mattermost/mattermost-app-servicenow/utils"
)

func fOAuthConnect(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	conf := app.GetOAuthConfig(c.Context)

	state, ok := c.Values["state"].(string)
	if !ok {
		utils.WriteCallErrorResponse(w, "State not provided.")
		return
	}

	utils.WriteCallResponse(w, apps.CallResponse{
		Type: apps.CallResponseTypeOK,
		Data: conf.AuthCodeURL(state),
	})
}
