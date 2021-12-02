package mattermost

import (
	"fmt"
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

	callState := &configureOAuthCallState{}
	callState.FromState(c.State)

	action := callState.Action

	if action == formActionSubmit {
		config.SetServiceNowInstance(c.GetValue(configureOAuthServiceNowInstanceValue, ""), c.Context)

		err := config.SetOAuthConfig(
			c.GetValue(configureOAuthClientIDValue, ""),
			c.GetValue(configureOAuthClientSecretValue, ""),
			c.Context,
		)
		if err != nil {
			utils.WriteCallErrorResponse(w, fmt.Sprintf("Could not setup the configuration. Error: %v", err.Error()))
			return
		}

		utils.WriteCallStandardResponse(w, "Configuration updated")

		return
	}

	utils.WriteCallResponse(w, apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: getConfigureOAuthForm(c.Values, formActionSubmit, c.Context),
	})
}

func fConfigureOAuthForm(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	if !c.Context.ExpandedContext.ActingUser.IsSystemAdmin() {
		utils.WriteCallErrorResponse(w, "You must be a system admin to configure oauth.")
		return
	}

	utils.WriteCallResponse(w, apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: getConfigureOAuthForm(nil, formActionOpen, c.Context),
	})
}

func getConfigureOAuthForm(v map[string]interface{}, action formAction, cc apps.Context) *apps.Form {
	return &apps.Form{
		Title: "Configure OAuth",
		Fields: []apps.Field{{
			Name:       configureOAuthServiceNowInstanceValue,
			Label:      configureOAuthServiceNowInstanceValue,
			ModalLabel: "Service Now Instance",
			Type:       apps.FieldTypeText,
			Value:      utils.GetStringFromMapInterface(v, configureOAuthServiceNowInstanceValue, config.ServiceNowInstance(cc)),
		}, {
			Name:       configureOAuthClientIDValue,
			Label:      configureOAuthClientIDValue,
			ModalLabel: "Client ID",
			Type:       apps.FieldTypeText,
			Value:      utils.GetStringFromMapInterface(v, configureOAuthClientIDValue, cc.OAuth2.ClientID),
		}, {
			Name:        configureOAuthClientSecretValue,
			Label:       configureOAuthClientSecretValue,
			ModalLabel:  "Client Secret",
			Type:        apps.FieldTypeText,
			TextSubtype: "password",
			Value:       utils.GetStringFromMapInterface(v, configureOAuthClientSecretValue, cc.OAuth2.ClientSecret),
		}},
		Call: getConfigureOAuthCall(action),
	}
}

func getConfigureOAuthCall(action formAction) *apps.Call {
	return &apps.Call{
		Path: string(constants.BindingPathConfigureOAuth),
		Expand: &apps.Expand{
			ActingUser:            apps.ExpandAll,
			OAuth2App:             apps.ExpandAll,
			ActingUserAccessToken: apps.ExpandAll,
		},
		State: configureOAuthCallState{
			Action: action,
		},
	}
}
