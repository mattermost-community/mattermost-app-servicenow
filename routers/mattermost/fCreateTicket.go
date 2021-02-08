package mattermost

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/app"
	"github.com/mattermost/mattermost-app-servicenow/clients/servicenowclient"
	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
	"github.com/pkg/errors"
)

var ErrCannotCreateClient = errors.New("cannot create client")

func fCreateTicket(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *api.Call) {
	if !app.IsUserConnected(c.Context.BotAccessToken, c.Context.MattermostSiteURL, c.Context.ActingUserID) {
		utils.WriteCallErrorResponse(w, "User is not connected. Please connect before creating a ticket.")
		return
	}

	table := r.URL.Query().Get(constants.TableIDGetField)
	action := r.URL.Query().Get(string(formActionQueryField))

	t, found := config.GetTables()[table]
	if !found {
		utils.WriteCallErrorResponse(w, "Table definition not found.")
	}

	// Command is asking for the form definition
	if c.Type == api.CallTypeForm {
		utils.WriteCallResponse(w, api.CallResponse{
			Type: api.CallResponseTypeForm,
			Form: getCreateTicketForm(t.Fields, table, formActionOpen),
		})

		return
	}

	// Modal submits the information
	if action == string(formActionSubmit) {
		id, err := submitTicket(claims.ActingUserID, table, c)
		if err != nil {
			utils.WriteCallErrorResponse(w, fmt.Sprintf("Could not create the ticket. Error: %s", err.Error()))
			return
		}

		utils.WriteCallStandardResponse(w, fmt.Sprintf("Ticket created with sys_id %s.", id))

		return
	}

	// Open the modal with the information provided by the command or the post action
	var postField string
	if c.Context.ExpandedContext.Post != nil {
		postField = c.Context.ExpandedContext.Post.Message
	}

	fields := []*api.Field{}

	for _, v := range t.Fields {
		field := *v
		field.Value = c.GetValue(v.Name, "")

		if t.PostDefault == v.Name && len(postField) != 0 {
			field.Value = postField
		}

		fields = append(fields, &field)
	}

	utils.WriteCallResponse(w, api.CallResponse{
		Type: api.CallResponseTypeForm,
		Form: getCreateTicketForm(fields, table, formActionSubmit),
	})
}

func getCreateTicketForm(fields []*api.Field, table string, action formAction) *api.Form {
	return &api.Form{
		Title:  "Create ticket",
		Fields: fields,
		Call:   getCreateTicketCall(table, action),
	}
}

func submitTicket(userID, table string, call *api.Call) (string, error) {
	c := servicenowclient.NewClient(call.Context.BotAccessToken, call.Context.MattermostSiteURL, userID)
	if c == nil {
		return "", ErrCannotCreateClient
	}

	return c.CreateIncident(table, call.Values)
}

func getCreateTicketCall(table string, action formAction) *api.Call {
	return &api.Call{
		URL: fmt.Sprintf("%s?%s=%s&%s=%s",
			constants.BindingPathCreate,
			constants.TableIDGetField,
			table,
			formActionQueryField,
			action),
		Expand: &api.Expand{Post: api.ExpandAll},
	}
}
