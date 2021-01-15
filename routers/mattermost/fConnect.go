package mattermost

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/app"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/store"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
)

func fConnect(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *api.Call) {
	state := utils.CreateOAuthState(c.Context.ActingUserID, c.Context.ChannelID)
	store.StoreState(state)
	conf := app.GetOAuthConfig()
	utils.WriteCallStandardResponse(w, fmt.Sprintf("Follow this link to connect: [link](%s)", conf.AuthCodeURL(state)))
}

func getConnectCall() *api.Call {
	return &api.Call{URL: constants.BindingPathConnect}
}
