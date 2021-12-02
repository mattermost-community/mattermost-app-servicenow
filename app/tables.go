package app

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
)

func GetTablesBindings(cc apps.Context) (post, command, header *apps.Binding) {
	pt, ct, ht := filterTables(config.GetTables(cc))
	pb := baseBinding("Create Ticket", cc)
	cb := baseBinding("create-ticket", cc)
	hb := baseBinding("Create Ticket", cc)
	post = subBindings(pt, pb, false, cc)
	command = subBindings(ct, cb, true, cc)
	header = subBindings(ht, hb, false, cc)

	return
}

func baseBinding(label string, _ apps.Context) apps.Binding {
	return apps.Binding{
		Location: constants.LocationCreate,
		Label:    label,
		Icon:     "now-mobile-icon.png",
		Bindings: []apps.Binding{},
	}
}

func subBindings(tt config.TablesConfig, base apps.Binding, useLocationLabel bool, _ apps.Context) *apps.Binding {
	switch len(tt) {
	case 0:
		return nil
	case 1:
		for _, t := range tt {
			base.Call = &apps.Call{
				Path: t.ID,
			}

			return &base
		}
	}

	for _, t := range tt {
		label := t.DisplayName
		if useLocationLabel {
			label = t.ID
		}

		base.Bindings = append(base.Bindings, apps.Binding{
			Location: apps.Location(t.ID),
			Label:    label,
			Icon:     "now-mobile-icon.png",
			Call: &apps.Call{
				Path: t.ID,
			},
		})
	}

	return &base
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
