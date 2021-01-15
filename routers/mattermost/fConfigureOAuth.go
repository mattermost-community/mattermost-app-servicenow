package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
)

const (
	configureOAuthClientIDValue           = "clientID"
	configureOAuthClientSecretValue       = "clientSecret"
	configureOAuthServiceNowInstanceValue = "instance"
)

func fConfigureOAuth(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *api.Call) {
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

	utils.WriteCallResponse(w, api.CallResponse{
		Type: api.CallResponseTypeForm,
		Form: &api.Form{
			Title: "Configure OAuth",
			Fields: []*api.Field{
				{
					Name:       configureOAuthServiceNowInstanceValue,
					ModalLabel: "Service Now Instance",
					Type:       api.FieldTypeText,
					Value:      config.ServiceNowInstance(),
				},
				{
					Name:       configureOAuthClientIDValue,
					ModalLabel: "Client ID",
					Type:       api.FieldTypeText,
					Value:      conf.ClientID,
				},
				{
					Name:        configureOAuthClientSecretValue,
					ModalLabel:  "Client Secret",
					Type:        api.FieldTypeText,
					TextSubtype: "password",
					Value:       conf.ClientSecret,
				},
			},
		},
		Call: getConfigureOAuthCall(),
	})
}

func getConfigureOAuthCall() *api.Call {
	return &api.Call{
		Type:   api.CallTypeSubmit,
		URL:    constants.BindingPathConfigureOAuth,
		Expand: &api.Expand{ActingUser: api.ExpandAll},
	}
}
