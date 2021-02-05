package mattermost

import (
	"fmt"
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

	if c.Type == api.CallTypeForm {
		utils.WriteCallResponse(w, api.CallResponse{
			Type: api.CallResponseTypeForm,
			Form: getConfigureOAuthForm(nil, formActionOpen),
		})
		return
	}

	action := r.URL.Query().Get(string(formActionQueryField))

	if action == string(formActionSubmit) {
		config.SetServiceNowInstance(c.GetValue(configureOAuthServiceNowInstanceValue, ""))
		config.SetOAuthConfig(config.OAuthConfig{
			ClientID:     c.GetValue(configureOAuthClientIDValue, ""),
			ClientSecret: c.GetValue(configureOAuthClientSecretValue, ""),
		})

		utils.WriteCallStandardResponse(w, "Configuration updated")
		return
	}

	utils.WriteCallResponse(w, api.CallResponse{
		Type: api.CallResponseTypeForm,
		Form: getConfigureOAuthForm(c.Values, formActionSubmit),
	})
}

func getConfigureOAuthForm(v map[string]interface{}, action formAction) *api.Form {
	conf := config.OAuth()

	return &api.Form{
		Title: "Configure OAuth",
		Fields: []*api.Field{
			{
				Name:       configureOAuthServiceNowInstanceValue,
				ModalLabel: "Service Now Instance",
				Type:       api.FieldTypeText,
				Value:      utils.GetStringFromMapInterface(v, configureOAuthServiceNowInstanceValue, config.ServiceNowInstance()),
			},
			{
				Name:       configureOAuthClientIDValue,
				ModalLabel: "Client ID",
				Type:       api.FieldTypeText,
				Value:      utils.GetStringFromMapInterface(v, configureOAuthClientIDValue, conf.ClientID),
			},
			{
				Name:        configureOAuthClientSecretValue,
				ModalLabel:  "Client Secret",
				Type:        api.FieldTypeText,
				TextSubtype: "password",
				Value:       utils.GetStringFromMapInterface(v, configureOAuthClientSecretValue, conf.ClientSecret),
			},
		},
		Call: getConfigureOAuthCall(action),
	}
}

func getConfigureOAuthCall(action formAction) *api.Call {
	return &api.Call{
		URL:    fmt.Sprintf("%s?%s=%s", constants.BindingPathConfigureOAuth, formActionQueryField, action),
		Expand: &api.Expand{ActingUser: api.ExpandAll},
	}
}
