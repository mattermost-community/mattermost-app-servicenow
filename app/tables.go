package app

import (
	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
)

func GetTablesBindings() *api.Binding {
	tt := filterTicketable(config.GetTables())
	out := &api.Binding{
		Location: constants.BindingPathCreate,
		Label:    "Create ticket",
		Bindings: []*api.Binding{},
	}
	switch len(tt) {
	case 0:
		return nil
	case 1:
		for _, t := range tt {
			out.Call = &api.Call{
				URL: t.ID,
			}
			return out
		}
	}
	for _, t := range tt {
		out.Bindings = append(out.Bindings, &api.Binding{
			Location: api.Location(t.ID),
			Label:    t.DisplayName,
			Call: &api.Call{
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
