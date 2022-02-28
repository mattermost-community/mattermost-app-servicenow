package function

import (
	"fmt"

	"github.com/mattermost/mattermost-plugin-apps/apps"

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
			message := "ServiceNow App"
			if BuildDate != "" {
				message += fmt.Sprintf(" built: %s from [%s](https://github.com/mattermost/mattermost-app-servicenow/commit/%s)",
					BuildDate, BuildHashShort, BuildHash)
			}
			if a.mode != "" {
				message += ", running as " + a.mode
			}
			message += "\n\n"
			message += fmt.Sprintf("☑ OAuth2 complete URL: `%s`\n", creq.Context.OAuth2.CompleteURL)

			connectLink := ""
			if creq.Context.OAuth2.ClientID != "" {
				message += "☑ OAuth2 App: configured.\n"
				connectLink = fmt.Sprintf(" Click [here](%s) to connect.\n", creq.Context.OAuth2.ConnectURL)
			} else {
				message += "☐ OAuth2 App: not configured. Please use `/servicenow configure`.\n"
			}

			if u := creq.OAuth2User(); u != nil {
				message += fmt.Sprintf("☑ Connected to ServiceNow as %s.\n", u.RemoteID)
			} else {
				message += fmt.Sprintf("☐ Not connected to ServiceNow.%s\n", connectLink)
			}

			return apps.NewTextResponse(message)
		},
	}
}
