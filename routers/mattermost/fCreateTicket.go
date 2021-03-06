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

func fCreateTicket(w http.ResponseWriter, r *http.Request, c *apps.Call) {
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
	if c.Type == apps.CallTypeForm {
		utils.WriteCallResponse(w, apps.CallResponse{
			Type: apps.CallResponseTypeForm,
			Form: getCreateTicketForm(t.Fields, table, formActionOpen),
		})

		return
	}

	// Modal submits the information
	if action == string(formActionSubmit) {
		id, err := submitTicket(c.Context.ActingUserID, table, c)
		if err != nil {
			utils.WriteCallErrorResponse(w, fmt.Sprintf("Could not create the ticket. Error: %s", err.Error()))
			return
		}

		navToURI := fmt.Sprintf("/%s?sys_id=%s", table, url.QueryEscape(id))
		ticketLink := fmt.Sprintf("%s/nav_to.do?uri=%s", config.ServiceNowInstance(), url.QueryEscape(navToURI))
		utils.WriteCallStandardResponse(w, fmt.Sprintf("Ticket created with [sys_id %s](%s).", id, ticketLink))

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

func getCreateTicketForm(fields []*apps.Field, table string, action formAction) *apps.Form {
	return &apps.Form{
		Title:  "Create ticket",
		Fields: fields,
		Call:   getCreateTicketCall(table, action),
	}
}

func submitTicket(userID, table string, call *apps.Call) (string, error) {
	c := servicenowclient.NewClient(call.Context.BotAccessToken, call.Context.MattermostSiteURL, userID)
	if c == nil {
		return "", ErrCannotCreateClient
	}

	return c.CreateIncident(table, call.Values)
}

func getCreateTicketCall(table string, action formAction) *apps.Call {
	return &apps.Call{
		Path: fmt.Sprintf("%s?%s=%s&%s=%s",
			constants.BindingPathCreate,
			constants.TableIDGetField,
			table,
			formActionQueryField,
			action),
		Expand: &apps.Expand{Post: apps.ExpandAll},
	}
}
