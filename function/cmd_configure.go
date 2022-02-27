package function

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/appclient"
	"github.com/mattermost/mattermost-plugin-apps/utils"
	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-app-servicenow/goapp"
)

var configureCommand = goapp.Command{
	Name:        "configure",
	Hint:        "[ --url --client-id --client-secret ]",
	Description: "Configure the ServiceNow URL, OAuth2 app credentials",

	BaseForm: &apps.Form{
		Submit: &apps.Call{
			Expand: &apps.Expand{
				ActingUserAccessToken: apps.ExpandAll,
				ActingUser:            apps.ExpandSummary,
			},
		},
		Title: "Configure ServiceNow App",
		Fields: []apps.Field{
			{
				Name:        fURL,
				ModalLabel:  "ServiceNow URL",
				Type:        apps.FieldTypeText,
				Description: "The root URL of your ServiceNow instance.",
				IsRequired:  true,
			},
			{
				Name:        fClientID,
				ModalLabel:  "Client ID",
				Type:        apps.FieldTypeText,
				Description: "ServiceNow Client ID for the Mattermost App.",
				IsRequired:  true,
			},
			{
				Type:        apps.FieldTypeText,
				TextSubtype: apps.TextFieldSubtypePassword,
				Name:        fClientSecret,
				ModalLabel:  "Client Secret",
				Description: "ServiceNow Client Secret for the Mattermost App.",
				IsRequired:  true,
			},
		},
	},

	Handler: goapp.RequireAdmin(func(creq goapp.CallRequest) apps.CallResponse {
		clientID := creq.GetValue(fClientID, "")
		clientSecret := creq.GetValue(fClientSecret, "")
		serviceNowURL := creq.GetValue(fURL, "")

		asActingUser := appclient.AsActingUser(creq.Context)
		err := asActingUser.StoreOAuth2App(creq.Context.AppID, apps.OAuth2App{
			RemoteURL:    serviceNowURL,
			ClientID:     clientID,
			ClientSecret: clientSecret,
		})
		if err != nil {
			return apps.NewErrorResponse(errors.Wrap(err, "failed to store Oauth2 configuration to Mattermost"))
		}

		return apps.NewTextResponse(
			"updated OAuth client credentials:\n"+
				"  - ServiceNow URL: `%s`\n"+
				"  - Client ID, ending in `%s`\n"+
				"  - Client Secret, ending in `%s`\n",
			serviceNowURL,
			utils.LastN(clientID, 8),
			utils.LastN(clientSecret, 4),
		)
	}),
}
