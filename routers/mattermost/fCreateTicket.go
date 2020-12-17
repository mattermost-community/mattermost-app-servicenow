package mattermost

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/clients/servicenowclient"
	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/server/apps"
)

func fCreateTicket(w http.ResponseWriter, r *http.Request, claims *apps.JWTClaims, c *apps.Call) {
	table := r.URL.Query().Get(constants.TableIDGetField)
	if len(c.Values) > 0 {
		err := submitTicket(claims.ActingUserID, table, c)
		if err != nil {
			utils.WriteCallErrorResponse(w, fmt.Sprintf("Could not create the ticket. Are you connected to Service Now? Error: %s", err.Error()))
			return
		}

		utils.WriteCallStandardResponse(w, "Ticket created")
		return
	}

	t, found := config.GetTables()[table]
	if !found {
		utils.WriteCallErrorResponse(w, "Table definition not found.")
	}

	var postField string
	if c.Context.ExpandedContext.Post != nil {
		postField = c.Context.ExpandedContext.Post.Message
	}

	fields := []*apps.Field{}
	for _, v := range t.Fields {
		field := *v
		if t.PostDefault == v.Name {
			field.Value = postField
		}
		fields = append(fields, &field)
	}

	utils.WriteCallResponse(w, apps.CallResponse{
		Type: apps.CallResponseTypeForm,
		Form: &apps.Form{
			Title:  "Create ticket",
			Fields: fields,
		},
		Call: getCreateTicketCall(table),
	})
}

func submitTicket(userID, table string, call *apps.Call) error {
	c := servicenowclient.NewClient(userID)
	if c == nil {
		return fmt.Errorf("cannot create client")
	}
	return c.CreateIncident(table, call.Values)
}

func getCreateTicketCall(table string) *apps.Call {
	return &apps.Call{
		URL:    fmt.Sprintf("%s%s?%s=%s", config.BaseURL(), constants.BindingPathCreate, constants.TableIDGetField, table),
		Expand: &apps.Expand{Post: apps.ExpandAll},
	}
}
