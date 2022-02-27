package function

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	root "github.com/mattermost/mattermost-app-servicenow"
	"github.com/mattermost/mattermost-app-servicenow/goapp"
)

func (a *App) getBindings(creq goapp.CallRequest) apps.CallResponse {
	bindings := goapp.AppendBinding(nil, &apps.Binding{
		Location: apps.LocationCommand,
		Bindings: []apps.Binding{
			{
				Label:       "servicenow",
				Description: "Create incidents in your ServiceNow instance",
				Icon:        root.AppManifest.Icon,

				Bindings: goapp.AppendBindings(
					a.commandBindings(creq),
					a.debugCommandBindings(creq),
				),
			},
		},
	})
	bindings = goapp.AppendBinding(bindings, a.postMenuBinding(creq))
	// bindings = goapp.AppendBinding(bindings, a.channelHeaderBinding(creq))

	return apps.NewDataResponse(bindings)
}

func (a *App) commandBindings(creq goapp.CallRequest) []apps.Binding {
	var bindings []apps.Binding

	// admin commands
	if creq.Context.ActingUser != nil && creq.Context.ActingUser.IsSystemAdmin() {
		bindings = append(bindings,
			configureCommand.Binding(creq),
			a.infoCommand().Binding(creq),
		)
	}

	// Do not show any more commands unless the app is configured
	if creq.Context.OAuth2.ClientID == "" {
		return bindings
	}

	// user commands
	if creq.Context.OAuth2.User == nil {
		// Not connected
		bindings = append(bindings, goapp.ConnectCommand("ServiceNow").Binding(creq))
		return bindings
	}

	// Connected
	bindings = append(bindings, a.createTicketCommandBinding(creq))
	bindings = append(bindings, goapp.DisconnectCommand("ServiceNow").Binding(creq))
	return bindings
}

func (a *App) postMenuBinding(creq goapp.CallRequest) *apps.Binding {
	if creq.Context.OAuth2.User == nil {
		return nil
	}
	return &apps.Binding{
		Location: apps.LocationPostMenu,
		Bindings: []apps.Binding{
			a.createTicketPostMenuBinding(creq),
		},
	}
}

func (a *App) channelHeaderBinding(creq goapp.CallRequest) *apps.Binding {
	if creq.Context.OAuth2.User == nil {
		return nil
	}
	return &apps.Binding{
		Location: apps.LocationChannelHeader,
		Bindings: []apps.Binding{
			a.createTicketChannelHeaderBinding(creq),
		},
	}
}

func (a *App) debugCommandBindings(creq goapp.CallRequest) []apps.Binding {
	if !creq.Context.DeveloperMode &&
		(creq.Context.ActingUser == nil || !creq.Context.ActingUser.IsSystemAdmin()) {
		return nil
	}

	return nil
	// return []apps.Binding{
	// 	{
	// 		Label:    "debug",
	// 		Location: "debug",
	// 		Bindings: []apps.Binding{
	// 			debugUserInfo.Binding(creq),
	// 		},
	// 		Icon: icon,
	// 	},
	// }
}
