package goapp

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/appclient"
)

func ConnectCommand(name string) Command {
	return Command{
		Name: "connect",
		BaseSubmit: &apps.Call{
			Expand: &apps.Expand{
				OAuth2App: apps.ExpandAll,
			},
		},

		Handler: func(creq CallRequest) apps.CallResponse {
			message := fmt.Sprintf("[Connect](%s) your account", creq.Context.OAuth2.ConnectURL)
			if name != "" {
				message += "to " + name
			}
			message += "."
			return apps.NewTextResponse(message)
		},
	}
}

func DisconnectCommand(name string) Command {
	return Command{
		Name: "disconnect",
		BaseSubmit: &apps.Call{
			Expand: &apps.Expand{
				ActingUserAccessToken: apps.ExpandAll,
			},
		},

		Handler: func(creq CallRequest) apps.CallResponse {
			asActingUser := appclient.AsActingUser(creq.Context)
			err := asActingUser.StoreOAuth2User(nil)
			if err != nil {
				return apps.NewErrorResponse(errors.Wrap(err, "failed to reset the stored user"))
			}
			message := "Disconnected your account"
			if name != "" {
				message += "from " + name
			}
			message += "."
			return apps.NewTextResponse(message)
		},
	}
}
