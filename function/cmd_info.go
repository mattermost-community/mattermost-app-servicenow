package function

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/utils"

	"github.com/mattermost/mattermost-app-servicenow/goapp"
)

func (a *App) infoCommand() goapp.Command {
	return goapp.Command{
		Name:        "info",
		Description: "ServiceNow App information",

		BaseSubmit: &apps.Call{
			Expand: &apps.Expand{
				ActingUser: apps.ExpandSummary,
				OAuth2App:  apps.ExpandAll,
				OAuth2User: apps.ExpandAll,
			},
		},

		Handler: func(creq goapp.CallRequest) apps.CallResponse {
			title := "ServiceNow App"
			if BuildDate != "" {
				title += fmt.Sprintf(" built: %s from [%s](https://github.com/mattermost/mattermost-app-servicenow/commit/%s)",
					BuildDate, BuildHashShort, BuildHash)
			}
			if a.mode != "" {
				title += ", running as " + a.mode
			}
			title += "\n"

			oauth2CallbackMessage := fmt.Sprintf("☑ OAuth2 callback URL: `%s`\n", creq.Context.OAuth2.CompleteURL)

			connectLink := ""
			oauth2AppMessage := "☐ OAuth2 App: not configured. Please use `/servicenow configure`.\n"
			if creq.Context.OAuth2.ClientID != "" {
				oauth2AppMessage = fmt.Sprintf("☑ OAuth2 App: `%s`, Client ID `%s`, Secret `%s`\n",
					creq.Context.OAuth2.RemoteRootURL,
					utils.LastN(creq.Context.OAuth2.ClientID, 8),
					utils.LastN(creq.Context.OAuth2.ClientSecret, 4))
				connectLink = fmt.Sprintf(" Click [here](%s) to connect.\n", creq.Context.OAuth2.ConnectURL)
			}

			connectMessage := fmt.Sprintf("☐ Not connected to ServiceNow.%s\n", connectLink)
			if u := creq.OAuth2User(); u != nil {
				remote := u.RemoteID
				if remote == "" {
					remote = "_unknown_"
				}
				connectMessage = fmt.Sprintf("☑ Connected to ServiceNow as %s.\n", remote)
			}

			return apps.NewTextResponse(strings.Join([]string{
				title, "\n",
				oauth2CallbackMessage,
				oauth2AppMessage,
				connectMessage,
			}, "\n"))
		},
	}
}
