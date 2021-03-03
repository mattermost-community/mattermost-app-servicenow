package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
	"github.com/mattermost/mattermost-plugin-apps/server/utils/md"

	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
)

func fInstall(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *apps.Call) {
	response := apps.CallResponse{}
	if c.Type != apps.CallTypeSubmit {
		response.Type = apps.CallResponseTypeError
		response.ErrorText = "Not supported call type"
		utils.WriteCallResponse(w, response)
	}

	config.SetLocalConfig(config.LocalConfig{
		BaseURL:        config.Local().BaseURL,
		MattermostURL:  c.Context.MattermostSiteURL,
		BotAccessToken: c.Context.BotAccessToken,
	})

	response.Type = apps.CallResponseTypeOK
	response.Markdown = md.Markdownf("Service now installed! "+
		"Please run `/%s configure oauth` to configure the link between Mattermost and your Service Now instance.",
		constants.CommandTrigger)
	utils.WriteCallResponse(w, response)
}
