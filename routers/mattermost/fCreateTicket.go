package mattermost

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-app-servicenow/app"
	"github.com/mattermost/mattermost-app-servicenow/clients/servicenowclient"
	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
)

var ErrCannotCreateClient = errors.New("cannot create client")

func fCreateTicketSubmit(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	if !app.IsUserConnected(c.Context) {
		utils.WriteCallErrorResponse(w, "User is not connected. Please connect before creating a ticket.")
		return
	}

	callState := &createTicketCallState{}
	callState.FromState(c.State)

	table := callState.Table
	action := callState.Action

	t, found := config.GetTables(c.Context)[table]
	if !found {
		utils.WriteCallErrorResponse(w, fmt.Sprintf("Table definition '%s' not found", table))
		return
	}

	// Modal submits the information
	if action == formActionSubmit {
		id, err := submitTicket(table, c)
		if err != nil {
			utils.WriteCallErrorResponse(w, fmt.Sprintf("Could not create the ticket. Error: %s", err.Error()))
			return
		}

		navToURI := fmt.Sprintf("/%s?sys_id=%s", table, url.QueryEscape(id))
		ticketLink := fmt.Sprintf("%s/nav_to.do?uri=%s", config.ServiceNowInstance(c.Context), url.QueryEscape(navToURI))
		utils.WriteCallStandardResponse(w, fmt.Sprintf("Ticket created [here](%s).", ticketLink))

		return
	}

	// Open the modal with the information provided by the command or the post action
	var postField string
	if c.Context.ExpandedContext.Post != nil {
		postField = c.Context.ExpandedContext.Post.Message
	}

	fields := []*apps.Field{}

	for _, v := range t.Fields {
		field := *v
		field.Value = c.GetValue(v.Name, "")

		if t.PostDefault == v.Name && len(postField) != 0 {
			field.Value = postField
		}

		fields = append(fields, &field)
	}

	utils.WriteCallResponse(w, apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: getCreateTicketForm(fields, table, formActionSubmit),
	})
}

func fCreateTicketForm(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	if !app.IsUserConnected(c.Context) {
		utils.WriteCallErrorResponse(w, "User is not connected. Please connect before creating a ticket.")
		return
	}

	callState := &createTicketCallState{}
	callState.FromState(c.State)

	table := callState.Table
	action := callState.Action

	t, found := config.GetTables(c.Context)[table]
	if !found {
		utils.WriteCallErrorResponse(w, fmt.Sprintf("Table definition '%s' not found", table))
		return
	}

	if action == formActionOpen {
		for _, field := range t.Fields {
			if field.Name == config.FieldShortDescription {
				field.IsRequired = false
			}
		}
	}

	utils.WriteCallResponse(w, apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: getCreateTicketForm(t.Fields, table, formActionOpen),
	})
}

func getCreateTicketForm(fields []*apps.Field, table string, action formAction) *apps.Form {
	return &apps.Form{
		Title:  "Create ticket",
		Fields: fields,
		Call:   getCreateTicketCall(table, action),
	}
}

func submitTicket(table string, call *apps.CallRequest) (string, error) {
	c := servicenowclient.NewClient(call.Context)
	if c == nil {
		return "", ErrCannotCreateClient
	}

	return c.CreateIncident(table, call.Values, call.Context)
}

func getCreateTicketCall(table string, action formAction) *apps.Call {
	return &apps.Call{
		Path: string(constants.BindingPathCreate),
		Expand: &apps.Expand{
			Post:       apps.ExpandAll,
			OAuth2App:  apps.ExpandAll,
			OAuth2User: apps.ExpandAll,
		},
		State: createTicketCallState{
			Action: action,
			Table:  table,
		},
	}
}
