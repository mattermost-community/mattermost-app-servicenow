package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-servicenow/app"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
)

func fBindings(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	baseCommand := apps.Binding{
		Label:       constants.CommandTrigger,
		Location:    constants.CommandTrigger,
		Description: "Create incidents in your ServiceNow instance",
		Icon:        "now-mobile-icon.png",
	}

	connectionCommand := getConnectBinding(c.Context)

	if app.IsUserConnected(c.Context) {
		connectionCommand = getDisconnectBinding(c.Context)
	}

	baseCommand.Bindings = append(baseCommand.Bindings, connectionCommand)

	user := c.Context.ActingUser
	if user != nil && user.IsSystemAdmin() {
		baseCommand.Bindings = append(baseCommand.Bindings, getSysAdminCommandBindings(c.Context))
	}

	out := []apps.Binding{}

	if app.IsUserConnected(c.Context) {
		postBindings, commandBindings, headerBindings := app.GetTablesBindings(c.Context)
		if postBindings != nil {
			out = append(out, apps.Binding{
				Location: apps.LocationPostMenu,
				Bindings: []apps.Binding{generateTableBindingsCalls(*postBindings)},
			})
		}

		if commandBindings != nil {
			baseCommand.Bindings = append(baseCommand.Bindings, generateTableBindingsCalls(*commandBindings))
		}

		if headerBindings != nil {
			out = append(out, apps.Binding{
				Location: apps.LocationChannelHeader,
				Bindings: []apps.Binding{generateTableBindingsCalls(*headerBindings)},
			})
		}
	}

	commands := apps.Binding{
		Location: apps.LocationCommand,
		Bindings: []apps.Binding{
			baseCommand,
		},
	}

	out = append(out, commands)

	utils.WriteBindings(w, out)
}

func generateTableBindingsCalls(b apps.Binding) apps.Binding {
	if len(b.Bindings) == 0 {
		b.Call = getCreateTicketCall(b.Call.Path, formActionOpen)
	}

	for _, subBinding := range b.Bindings {
		subBinding.Call = getCreateTicketCall(subBinding.Call.Path, formActionOpen)
	}

	return b
}

func getSysAdminCommandBindings(_ apps.Context) apps.Binding {
	return apps.Binding{
		Location:    constants.LocationConfigure,
		Label:       "config",
		Icon:        "now-mobile-icon.png",
		Hint:        "",
		Description: "Configure the app",
		Bindings: []apps.Binding{{
			Location:    constants.LocationConfigureOAuth,
			Label:       "oauth",
			Icon:        "now-mobile-icon.png",
			Hint:        "",
			Description: "Configure OAuth options",
			Call:        getConfigureOAuthCall(formActionOpen),
		}},
	}
}
func getConnectBinding(_ apps.Context) apps.Binding {
	return apps.Binding{
		Location:    constants.LocationConnect,
		Label:       "connect",
		Icon:        "now-mobile-icon.png",
		Hint:        "",
		Description: "Connect your ServiceNow account",
		Form:        &apps.Form{},
		Call:        getConnectCall(),
	}
}

func getDisconnectBinding(_ apps.Context) apps.Binding {
	return apps.Binding{
		Location:    constants.LocationDisconnect,
		Label:       "disconnect",
		Icon:        "now-mobile-icon.png",
		Hint:        "",
		Description: "Disconnect from ServiceNow",
		Form:        &apps.Form{},
		Call:        getDisconnectCall(),
	}
}
