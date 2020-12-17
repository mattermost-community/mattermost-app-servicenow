package mattermost

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/store"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/server/apps"
)

func fDisconnect(w http.ResponseWriter, r *http.Request, claims *apps.JWTClaims, c *apps.Call) {
	store.DeleteToken(claims.ActingUserID)
	utils.WriteCallStandardResponse(w, "You are disconnected from Service Now.")
}

func getDisconnectCall() *apps.Call {
	return &apps.Call{URL: fmt.Sprintf("%s%s", config.BaseURL(), constants.BindingPathDisconnect)}
}
