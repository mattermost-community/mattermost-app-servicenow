package function

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-servicenow/goapp"
)

func bindings(creq goapp.CallRequest) apps.CallResponse {
	bindings := goapp.AppendBinding(nil, &apps.Binding{
		Location: apps.LocationCommand,
		Bindings: []apps.Binding{
			{
				Label:       servicenow,
				Description: "Create incidents in your ServiceNow instance",
				Icon:        icon,

				Bindings: goapp.AppendBindings(
					commandBindings(creq),
					debugCommandBindings(creq),
				),
			},
		},
	})
	bindings = goapp.AppendBinding(bindings, postMenuBinding(creq))
	bindings = goapp.AppendBinding(bindings, channelHeaderBinding(creq))

	return apps.NewDataResponse(bindings)
}

func commandBindings(creq goapp.CallRequest) []apps.Binding {
	var bindings []apps.Binding

	// admin commands
	if creq.Context.ActingUser != nil && creq.Context.ActingUser.IsSystemAdmin() {
		bindings = append(bindings) // configure.Binding(creq),
		// info.Binding(creq),

	}

	// Do not show any more commands unless the app is configured
	if creq.Context.OAuth2.ClientID == "" {
		return bindings
	}

	// user commands
	if creq.Context.OAuth2.User == nil {
		// Not connected
		bindings = append(bindings, goapp.ConnectCommand(ServiceNow).Binding(creq))
	} else {
		// Connected
		bindings = append(bindings, goapp.DisconnectCommand(ServiceNow).Binding(creq))
	}

	return bindings
}

func postMenuBinding(creq goapp.CallRequest) *apps.Binding {
	if creq.Context.OAuth2.User == nil {
		return nil
	}
	return &apps.Binding{
		Location: apps.LocationPostMenu,
		Bindings: []apps.Binding{
			createTicketBinding(creq, apps.LocationPostMenu),
		},
	}
}

func channelHeaderBinding(creq goapp.CallRequest) *apps.Binding {
	if creq.Context.OAuth2.User == nil {
		return nil
	}
	return &apps.Binding{
		Location: apps.LocationPostMenu,
		Bindings: []apps.Binding{
			createTicketBinding(creq, apps.LocationChannelHeader),
		},
	}
}

func debugCommandBindings(creq goapp.CallRequest) []apps.Binding {
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
