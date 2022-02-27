package function

import (
	"fmt"
	"net/url"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/utils"
	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-app-servicenow/goapp"
)

func createTicketHandler(creq goapp.CallRequest) (string, error) {
	tableID := creq.GetValue(fTable, "")
	tables := GetTables(creq)
	if _, ok := tables[tableID]; !ok {
		return "", utils.NewNotFoundError("%q: unknown table", tableID)
	}

	c, err := makeClient(creq)
	if err != nil {
		return "", err
	}

	id, err := c.CreateIncident(creq, tableID, creq.Values)
	if err != nil {
		return "", errors.Wrap(err, "failed to create ticket")
	}

	navToURI := fmt.Sprintf("/%s?sys_id=%s", url.PathEscape(tableID), url.QueryEscape(id))
	ticketLink := fmt.Sprintf("%s/nav_to.do?uri=%s", creq.Context.OAuth2.RemoteURL, navToURI)
	return fmt.Sprintf("Ticket created [here](%s).", ticketLink), nil
}

func (a *App) createTicketFormHandler(creq goapp.CallRequest) (apps.Form, error) {
	tableID := creq.GetValue(fTable, "")
	tables := GetTables(creq)
	table, ok := tables[tableID]
	if !ok {
		return apps.Form{}, utils.NewNotFoundError("%q: unknown table", tableID)
	}
	return *a.createTicketForm(creq, table), nil
}

func (a *App) createTicketForm(creq goapp.CallRequest, table Table) *apps.Form {
	var message string
	if creq.Context.Post != nil {
		message = creq.Context.Post.Message
	}

	fields := []apps.Field{}
	for _, v := range table.Fields {
		field := *v
		field.Value = creq.GetValue(v.Name, "")
		if table.PostFieldName == v.Name && message != "" {
			field.Value = message
		}
		fields = append(fields, field)
	}

	return &apps.Form{
		Title:  "Create a ServiceNow ticket",
		Header: fmt.Sprintf("Will create a new entry in table %q.", table.DisplayName),
		Icon:   a.Icon,
		Fields: fields,
		Submit: apps.NewCall("/create-ticket").
			WithState(map[string]string{
				fTable: table.ID,
			}).
			WithExpand(apps.Expand{
				ActingUserAccessToken: apps.ExpandAll,
				ActingUser:            apps.ExpandSummary,
				OAuth2App:             apps.ExpandAll,
				OAuth2User:            apps.ExpandAll,
			}),
	}
}

func (a *App) createTicketCommandBinding(creq goapp.CallRequest) apps.Binding {
	b := apps.Binding{
		Label:    "create-ticket",
		Location: apps.Location("create-ticket"),
		Icon:     a.App.Icon,
	}

	tables := GetTables(creq).forLocation(apps.LocationCommand)
	for _, table := range tables {
		if len(tables) == 1 {
			// If there is only 1 table, no sub-bindings needed, bind the form for the only table.
			b.Form = a.createTicketForm(creq, table)
			break
		}

		// add the sub-binding for the table.
		b.Bindings = append(b.Bindings, apps.Binding{
			Label:    table.ID,
			Location: apps.Location(table.ID),
			Icon:     a.App.Icon,
			Form:     a.createTicketForm(creq, table),
		})
	}
	return b
}

func (a *App) createTicketChannelHeaderBinding(creq goapp.CallRequest) apps.Binding {
	b := apps.Binding{
		Label:    "Create ticket",
		Location: apps.Location("create-ticket"),
		Icon:     a.App.Icon,
	}

	tables := GetTables(creq).forLocation(apps.LocationChannelHeader)
	for _, table := range tables {
		if len(tables) == 1 {
			// If there is only 1 table, no sub-bindings needed, bind the form for the only table.
			b.Form = a.createTicketForm(creq, table)
			break
		}

		// add the sub-binding for the table.
		b.Bindings = append(b.Bindings, apps.Binding{
			Label:    table.DisplayName,
			Location: apps.Location(table.ID),
			Icon:     a.App.Icon,
			Form:     a.createTicketForm(creq, table),
		})
	}
	return b
}

func (a *App) createTicketPostMenuBinding(creq goapp.CallRequest) apps.Binding {
	b := apps.Binding{
		Label:    "Create ticket",
		Location: apps.Location("create-ticket"),
		Icon:     a.App.Icon,
	}

	tables := GetTables(creq).forLocation(apps.LocationPostMenu)
	for _, table := range tables {
		submitToForm := apps.NewCall("/form/create-ticket").
			WithState(map[string]string{
				fTable: table.ID,
			}).
			WithExpand(apps.Expand{
				OAuth2App: apps.ExpandAll,
				Post:      apps.ExpandSummary,
			})

		if len(tables) == 1 {
			// If there is only 1 table, no sub-bindings needed, bind the form for the only table.
			b.Submit = submitToForm
			break
		}

		b.Bindings = append(b.Bindings, apps.Binding{
			Label:    table.DisplayName,
			Location: apps.Location(table.ID),
			Icon:     a.App.Icon,
			Submit:   submitToForm,
		})
	}
	return b
}
