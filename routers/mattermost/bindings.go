package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/app"
	"github.com/mattermost/mattermost-app-servicenow/clients/mattermostclient"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
)

func fBindings(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *api.Call) {
	commands := &api.Binding{
		Location: api.LocationCommand,
		Bindings: []*api.Binding{},
	}

	connectionCommand := &api.Binding{
		Location:    constants.LocationConnect,
		Label:       "connect",
		Icon:        "",
		Hint:        "",
		Description: "Connect your ServiceNow account",
		Call:        getConnectCall(),
	}

	if app.IsUserConnected(claims.ActingUserID) {
		connectionCommand = &api.Binding{
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
		commands.Bindings = append(commands.Bindings, &api.Binding{
			Location:    constants.LocationConfigure,
			Label:       "config",
			Icon:        "",
			Hint:        "",
			Description: "Configure the plugin",
			Bindings: []*api.Binding{
				{
					Location:    constants.LocationConfigureOAuth,
					Label:       "oauth",
					Icon:        "",
					Hint:        "",
					Description: "Configure OAuth options",
					Call:        getConfigureOAuthCall(configureOAuthActionOpen),
				},
			},
		})
	}

	out := []*api.Binding{
		commands,
	}

	if app.IsUserConnected(claims.ActingUserID) {
		tableBindings := app.GetTablesBindings()
		if tableBindings != nil {
			tableBindings = generateTableBindingsCalls(tableBindings)
			out = append(out, &api.Binding{
				Location: api.LocationPostMenu,
				Bindings: []*api.Binding{tableBindings},
			})
		}
	}

	utils.WriteBindings(w, out)
}

func generateTableBindingsCalls(b *api.Binding) *api.Binding {
	if len(b.Bindings) == 0 {
		b.Call = getCreateTicketCall(b.Call.URL)
	}

	for _, subBinding := range b.Bindings {
		subBinding.Call = getCreateTicketCall(subBinding.Call.URL)
	}

	return b
}
