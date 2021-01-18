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

	configureOAuthActionQueryField = "action"
	configureOAuthActionSubmit     = "submit"
	configureOAuthActionOpen       = "open"
)

type configureOAuthAction string

func fConfigureOAuth(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *api.Call) {
	if c.Type == api.CallTypeForm {
		utils.WriteCallResponse(w, api.CallResponse{
			Type: api.CallResponseTypeForm,
			Form: getConfigureOAuthForm(nil),
			Call: getConfigureOAuthCall(configureOAuthActionOpen),
		})
	}

	if !c.Context.ExpandedContext.ActingUser.IsSystemAdmin() {
		utils.WriteCallErrorResponse(w, "You must be a system admin to configure oauth.")
		return
	}

	action := r.URL.Query().Get(configureOAuthActionQueryField)

	if len(c.Values) > 0 && action == configureOAuthActionSubmit {
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
		Form: getConfigureOAuthForm(c.Values),
		Call: getConfigureOAuthCall(configureOAuthActionSubmit),
	})
}

func getConfigureOAuthForm(v map[string]interface{}) *api.Form {
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
	}
}

func getConfigureOAuthCall(action configureOAuthAction) *api.Call {
	return &api.Call{
		Type:   api.CallTypeSubmit,
		URL:    fmt.Sprintf("%s?%s=%s", constants.BindingPathConfigureOAuth, configureOAuthActionQueryField, action),
		Expand: &api.Expand{ActingUser: api.ExpandAll},
	}
}
