package mattermost

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/server/api"

	"github.com/mattermost/mattermost-app-servicenow/app"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/store"
	"github.com/mattermost/mattermost-app-servicenow/utils"
)

func fConnect(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *apps.Call) {
	state := utils.CreateOAuthState(c.Context.ActingUserID, c.Context.ChannelID)
	conf := app.GetOAuthConfig()

	store.SaveState(c.Context.BotAccessToken, c.Context.MattermostSiteURL, state)
	utils.WriteCallStandardResponse(w, fmt.Sprintf("Follow this link to connect: [link](%s)", conf.AuthCodeURL(state)))
}

func getConnectCall() *apps.Call {
	return &apps.Call{Path: constants.BindingPathConnect}
}
