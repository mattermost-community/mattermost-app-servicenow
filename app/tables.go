package app

import (
	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-plugin-apps/server/apps"
)

func GetTablesBindings() *apps.Binding {
	tt := filterTicketable(config.GetTables())
	out := &apps.Binding{
		Location: constants.BindingPathCreate,
		Label:    "Create ticket",
		Bindings: []*apps.Binding{},
	}
	switch len(tt) {
	case 0:
		return nil
	case 1:
		for _, t := range tt {
			out.Call = &apps.Call{
				URL: t.ID,
			}
			return out
		}
	}
	for _, t := range tt {
		out.Bindings = append(out.Bindings, &apps.Binding{
			Location: apps.Location(t.ID),
			Label:    t.DisplayName,
			Call: &apps.Call{
				URL: t.ID,
			},
		})
	}

	return out
}

func filterTicketable(tt config.TablesConfig) config.TablesConfig {
	out := config.TablesConfig{}
	for _, t := range tt {
		if t.Ticketable {
			out[t.ID] = t
		}
	}

	return out
}
