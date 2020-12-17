package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/server/apps"
	"github.com/mattermost/mattermost-plugin-apps/server/utils/md"
)

func fInstall(w http.ResponseWriter, r *http.Request, claims *apps.JWTClaims, c *apps.Call) {
	response := apps.CallResponse{}
	if c.Type != apps.CallTypeSubmit {
		response.Type = apps.CallResponseTypeError
		response.Error = "Not supported call type"
		utils.WriteCallResponse(w, response)
	}

	config.SetMattermostConfig(config.MattermostConfig{
		MattermostURL:  c.Context.Config.SiteURL,
		BotID:          c.Context.App.BotUserID,
		BotAccessToken: c.GetValue(apps.PropBotAccessToken, ""),
	})

	response.Type = apps.CallResponseTypeOK
	response.Markdown = md.Markdownf("Service now installed! Please run `/%s configure oauth` to configure the link between Mattermost and your Service Now instance.", constants.CommandTrigger)
	utils.WriteCallResponse(w, response)
}
