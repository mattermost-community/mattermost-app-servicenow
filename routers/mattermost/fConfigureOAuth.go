package mattermost

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/server/apps"
)

const (
	configureOAuthClientIDValue           = "clientID"
	configureOAuthClientSecretValue       = "clientSecret"
	configureOAuthServiceNowInstanceValue = "instance"
)

func fConfigureOAuth(w http.ResponseWriter, r *http.Request, claims *apps.JWTClaims, c *apps.Call) {
	if !c.Context.ExpandedContext.ActingUser.IsSystemAdmin() {
		utils.WriteCallErrorResponse(w, "You must be a system admin to configure oauth.")
		return
	}

	if len(c.Values) > 0 {
		config.SetServiceNowInstance(c.GetValue(configureOAuthServiceNowInstanceValue, ""))
		config.SetOAuthConfig(config.OAuthConfig{
			ClientID:     c.GetValue(configureOAuthClientIDValue, ""),
			ClientSecret: c.GetValue(configureOAuthClientSecretValue, ""),
		})

		utils.WriteCallStandardResponse(w, "Configuration updated")
		return
	}

	conf := config.OAuth()

	utils.WriteCallResponse(w, apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title: "Configure OAuth",
			Fields: []*apps.Field{
				{
					Name:       configureOAuthServiceNowInstanceValue,
					ModalLabel: "Service Now Instance",
					Type:       apps.FieldTypeText,
					Value:      config.ServiceNowInstance(),
				},
				{
					Name:       configureOAuthClientIDValue,
					ModalLabel: "Client ID",
					Type:       apps.FieldTypeText,
					Value:      conf.ClientID,
				},
				{
					Name:        configureOAuthClientSecretValue,
					ModalLabel:  "Client Secret",
					Type:        apps.FieldTypeText,
					TextSubtype: "password",
					Value:       conf.ClientSecret,
				},
			},
		},
		Call: getConfigureOAuthCall(),
	})
}

func getConfigureOAuthCall() *apps.Call {
	return &apps.Call{
		URL:    fmt.Sprintf("%s%s", config.BaseURL(), constants.BindingPathConfigureOAuth),
		Expand: &apps.Expand{ActingUser: apps.ExpandAll},
	}
}
