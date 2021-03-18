package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
)

const (
	configureOAuthClientIDValue           = "client_id"
	configureOAuthClientSecretValue       = "client_secret"
	configureOAuthServiceNowInstanceValue = "instance"
)

func fConfigureOAuthSubmit(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	if !c.Context.ExpandedContext.ActingUser.IsSystemAdmin() {
		utils.WriteCallErrorResponse(w, "You must be a system admin to configure oauth.")
		return
	}

	callState := ConfigureOAuthCallState{}
	callState.FromState(c.State)

	action := callState.Action

	if action == formActionSubmit {
		config.SetServiceNowInstance(c.GetValue(configureOAuthServiceNowInstanceValue, ""))
		config.SetOAuthConfig(config.OAuthConfig{
			ClientID:     c.GetValue(configureOAuthClientIDValue, ""),
			ClientSecret: c.GetValue(configureOAuthClientSecretValue, ""),
		})

		utils.WriteCallStandardResponse(w, "Configuration updated")

		return
	}

	utils.WriteCallResponse(w, apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: getConfigureOAuthForm(c.Values, formActionSubmit),
	})
}

func fConfigureOAuthForm(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	if !c.Context.ExpandedContext.ActingUser.IsSystemAdmin() {
		utils.WriteCallErrorResponse(w, "You must be a system admin to configure oauth.")
		return
	}

	utils.WriteCallResponse(w, apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: getConfigureOAuthForm(nil, formActionOpen),
	})
}

func getConfigureOAuthForm(v map[string]interface{}, action formAction) *apps.Form {
	conf := config.OAuth()

	return &apps.Form{
		Title: "Configure OAuth",
		Fields: []*apps.Field{
			{
				Name:       configureOAuthServiceNowInstanceValue,
				Label:      configureOAuthServiceNowInstanceValue,
				ModalLabel: "Service Now Instance",
				Type:       apps.FieldTypeText,
				Value:      utils.GetStringFromMapInterface(v, configureOAuthServiceNowInstanceValue, config.ServiceNowInstance()),
			},
			{
				Name:       configureOAuthClientIDValue,
				Label:      configureOAuthClientIDValue,
				ModalLabel: "Client ID",
				Type:       apps.FieldTypeText,
				Value:      utils.GetStringFromMapInterface(v, configureOAuthClientIDValue, conf.ClientID),
			},
			{
				Name:        configureOAuthClientSecretValue,
				Label:       configureOAuthClientSecretValue,
				ModalLabel:  "Client Secret",
				Type:        apps.FieldTypeText,
				TextSubtype: "password",
				Value:       utils.GetStringFromMapInterface(v, configureOAuthClientSecretValue, conf.ClientSecret),
			},
		},
		Call: getConfigureOAuthCall(action),
	}
}

func getConfigureOAuthCall(action formAction) *apps.Call {
	return &apps.Call{
		Path:   string(constants.BindingPathConfigureOAuth),
		Expand: &apps.Expand{ActingUser: apps.ExpandAll},
		State: ConfigureOAuthCallState{
			Action: action,
		},
	}
}
