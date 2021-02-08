package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
	"github.com/mattermost/mattermost-plugin-apps/server/utils/md"
)

func fInstall(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *api.Call) {
	response := api.CallResponse{}
	if c.Type != api.CallTypeSubmit {
		response.Type = api.CallResponseTypeError
		response.ErrorText = "Not supported call type"
		utils.WriteCallResponse(w, response)
	}

	config.SetLocalConfig(config.LocalConfig{
		MattermostURL:  c.Context.MattermostSiteURL,
		BotID:          c.Context.App.BotUserID,
		BotAccessToken: c.Context.BotAccessToken,
	})

	response.Type = api.CallResponseTypeOK
	response.Markdown = md.Markdownf("Service now installed! "+
		"Please run `/%s configure oauth` to configure the link between Mattermost and your Service Now instance.",
		constants.CommandTrigger)
	utils.WriteCallResponse(w, response)
}
