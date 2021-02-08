package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/store"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
)

func fDisconnect(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *api.Call) {
	store.DeleteToken(c.Context.BotAccessToken, c.Context.MattermostSiteURL, c.Context.ActingUserID)
	utils.WriteCallStandardResponse(w, "You are disconnected from Service Now.")
}

func getDisconnectCall() *api.Call {
	return &api.Call{URL: constants.BindingPathDisconnect}
}
