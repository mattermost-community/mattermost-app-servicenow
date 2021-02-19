package app

import (
	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
)

func GetTablesBindings() (post, command, header *api.Binding) {
	pt, ct, ht := filterTables(config.GetTables())
	pb := baseBinding("Create Ticket")
	cb := baseBinding("create-ticket")
	hb := baseBinding("Create Ticket")
	post = subBindings(pt, pb, false)
	command = subBindings(ct, cb, true)
	header = subBindings(ht, hb, false)

	return
}

func baseBinding(label string) *api.Binding {
	return &api.Binding{
		Location: constants.BindingPathCreate,
		Label:    label,
		Icon:     "https://docs.servicenow.com/bundle/mobile-rn/page/release-notes/mobile-apps/now-mobile/image/now-mobile-icon.png",
		Bindings: []*api.Binding{},
	}
}

func subBindings(tt config.TablesConfig, base *api.Binding, useLocationLabel bool) *api.Binding {
	switch len(tt) {
	case 0:
		return nil
	case 1:
		for _, t := range tt {
			base.Call = &api.Call{
				URL: t.ID,
			}

			return base
		}
	}

	for _, t := range tt {
		label := t.DisplayName
		if useLocationLabel {
			label = t.ID
		}

		base.Bindings = append(base.Bindings, &api.Binding{
			Location: api.Location(t.ID),
			Label:    label,
			Icon:     "https://docs.servicenow.com/bundle/mobile-rn/page/release-notes/mobile-apps/now-mobile/image/now-mobile-icon.png",
			Call: &api.Call{
				URL: t.ID,
			},
		})
	}

	return base
}

func filterTables(tt config.TablesConfig) (post, command, header config.TablesConfig) {
	post = config.TablesConfig{}
	command = config.TablesConfig{}
	header = config.TablesConfig{}

	for _, t := range tt {
		if t.Post {
			post[t.ID] = t
		}

		if t.Command {
			command[t.ID] = t
		}

		if t.Header {
			header[t.ID] = t
		}
	}

	return
}
