package mattermost

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/clients/servicenowclient"
	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
)

func fCreateTicket(w http.ResponseWriter, r *http.Request, claims *api.JWTClaims, c *api.Call) {
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

	fields := []*api.Field{}
	for _, v := range t.Fields {
		field := *v
		if t.PostDefault == v.Name {
			field.Value = postField
		}
		fields = append(fields, &field)
	}

	utils.WriteCallResponse(w, api.CallResponse{
		Type: api.CallResponseTypeForm,
		Form: &api.Form{
			Title:  "Create ticket",
			Fields: fields,
		},
		Call: getCreateTicketCall(table),
	})
}

func submitTicket(userID, table string, call *api.Call) error {
	c := servicenowclient.NewClient(userID)
	if c == nil {
		return fmt.Errorf("cannot create client")
	}
	return c.CreateIncident(table, call.Values)
}

func getCreateTicketCall(table string) *api.Call {
	return &api.Call{
		URL:    fmt.Sprintf("%s?%s=%s", constants.BindingPathCreate, constants.TableIDGetField, table),
		Expand: &api.Expand{Post: api.ExpandAll},
	}
}
