package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/app"
	"github.com/mattermost/mattermost-app-servicenow/clients/mattermostclient"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/server/apps"
)

func fBindings(w http.ResponseWriter, req *http.Request, claims *apps.JWTClaims, cc *apps.Context) {
	commands := &apps.Binding{
		Location: apps.LocationCommand,
		Bindings: []*apps.Binding{},
	}

	connectionCommand := &apps.Binding{
		Location:    constants.LocationConnect,
		Label:       "connect",
		Icon:        "",
		Hint:        "",
		Description: "Connect your ServiceNow account",
		Call:        getConnectCall(),
	}

	if app.IsUserConnected(claims.ActingUserID) {
		connectionCommand = &apps.Binding{
			Location:    constants.LocationDisconnect,
			Label:       "disconnect",
			Icon:        "",
			Hint:        "",
			Description: "Disconnect from ServiceNow",
			Call:        getDisconnectCall(),
		}
	}

	commands.Bindings = append(commands.Bindings, connectionCommand)

	user, err := mattermostclient.GetUser(claims.ActingUserID)
	if err == nil && user.IsSystemAdmin() {
		commands.Bindings = append(commands.Bindings, &apps.Binding{
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
					Call:        getConfigureOAuthCall(),
				},
			},
		})
	}

	out := []*apps.Binding{
		commands,
	}

	if app.IsUserConnected(claims.ActingUserID) {
		tableBindings := app.GetTablesBindings()
		if tableBindings != nil {
			tableBindings = generateTableBindingsCalls(tableBindings)
			out = append(out, &apps.Binding{
				Location: apps.LocationPostMenu,
				Bindings: []*apps.Binding{tableBindings},
			})
		}
	}

	utils.WriteBindings(w, out)
}

func generateTableBindingsCalls(b *apps.Binding) *apps.Binding {
	if len(b.Bindings) == 0 {
		b.Call = getCreateTicketCall(b.Call.URL)
	}

	for _, subBinding := range b.Bindings {
		subBinding.Call = getCreateTicketCall(subBinding.Call.URL)
	}

	return b
}
