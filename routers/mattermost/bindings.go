package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/server/api"

	"github.com/mattermost/mattermost-app-servicenow/app"
	"github.com/mattermost/mattermost-app-servicenow/clients/mattermostclient"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
)

func fBindings(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *apps.Call) {
	mattermostSiteURL := c.Context.MattermostSiteURL
	baseCommand := &apps.Binding{
		Label:       "servicenow",
		Location:    "servicenow",
		Description: "Create incidents in your ServiceNow instance",
		Icon:        utils.GetIconURL(mattermostSiteURL, "now-mobile-icon.png"),
	}

	commands := &apps.Binding{
		Location: apps.LocationCommand,
		Bindings: []*apps.Binding{
			baseCommand,
		},
	}

	connectionCommand := getConnectBinding(mattermostSiteURL)

	if app.IsUserConnected(c.Context.BotAccessToken, c.Context.MattermostSiteURL, claims.ActingUserID) {
		connectionCommand = getDisconnectBinding()
	}

	baseCommand.Bindings = append(baseCommand.Bindings, connectionCommand)
	client := mattermostclient.NewMMClient(c.Context.BotUserID, c.Context.BotAccessToken, c.Context.MattermostSiteURL)

	user, err := client.GetUser(claims.ActingUserID)
	if err == nil && user.IsSystemAdmin() {
		baseCommand.Bindings = append(baseCommand.Bindings, getSysAdminCommandBindings())
	}

	out := []*apps.Binding{}

	if app.IsUserConnected(c.Context.BotAccessToken, c.Context.MattermostSiteURL, claims.ActingUserID) {
		postBindings, commandBindings, headerBindings := app.GetTablesBindings(c.Context.MattermostSiteURL)
		if postBindings != nil {
			out = append(out, &apps.Binding{
				Location: apps.LocationPostMenu,
				Bindings: []*apps.Binding{generateTableBindingsCalls(postBindings)},
			})
		}

		if commandBindings != nil {
			baseCommand.Bindings = append(baseCommand.Bindings, generateTableBindingsCalls(commandBindings))
		}

		if headerBindings != nil {
			out = append(out, &apps.Binding{
				Location: apps.LocationChannelHeader,
				Bindings: []*apps.Binding{generateTableBindingsCalls(headerBindings)},
			})
		}
	}

	out = append(out, commands)

	utils.WriteBindings(w, out)
}

func generateTableBindingsCalls(b *apps.Binding) *apps.Binding {
	if len(b.Bindings) == 0 {
		b.Call = getCreateTicketCall(b.Call.Path, formActionOpen)
	}

	for _, subBinding := range b.Bindings {
		subBinding.Call = getCreateTicketCall(subBinding.Call.Path, formActionOpen)
	}

	return b
}

func getSysAdminCommandBindings() *apps.Binding {
	return &apps.Binding{
		Location:    constants.LocationConfigure,
		Label:       "config",
		Icon:        "",
		Hint:        "",
		Description: "Configure the plugin",
		Bindings: []*apps.Binding{
			{
				Location:    constants.LocationConfigureOAuth,
				Label:       "oauth",
				Icon:        "",
				Hint:        "",
				Description: "Configure OAuth options",
				Call:        getConfigureOAuthCall(formActionOpen),
			},
		},
	}
}
func getConnectBinding(mattermostSiteURL string) *apps.Binding {
	return &apps.Binding{
		Location:    constants.LocationConnect,
		Label:       "connect",
		Icon:        utils.GetIconURL(mattermostSiteURL, "now-mobile-icon.png"),
		Hint:        "",
		Description: "Connect your ServiceNow account",
		Form:        &apps.Form{},
		Call:        getConnectCall(),
	}
}

func getDisconnectBinding() *apps.Binding {
	return &apps.Binding{
		Location:    constants.LocationDisconnect,
		Label:       "disconnect",
		Icon:        "",
		Hint:        "",
		Description: "Disconnect from ServiceNow",
		Form:        &apps.Form{},
		Call:        getDisconnectCall(),
	}
}
