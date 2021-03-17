package app

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
)

func GetTablesBindings(siteURL string) (post, command, header *apps.Binding) {
	pt, ct, ht := filterTables(config.GetTables())
	pb := baseBinding(siteURL, "Create Ticket")
	cb := baseBinding(siteURL, "create-ticket")
	hb := baseBinding(siteURL, "Create Ticket")
	post = subBindings(siteURL, pt, pb, false)
	command = subBindings(siteURL, ct, cb, true)
	header = subBindings(siteURL, ht, hb, false)

	return
}

func baseBinding(siteURL, label string) *apps.Binding {
	return &apps.Binding{
		Location: constants.BindingPathCreate,
		Label:    label,
		Icon:     utils.GetIconURL(siteURL, "now-mobile-icon.png"),
		Bindings: []*apps.Binding{},
	}
}

func subBindings(siteURL string, tt config.TablesConfig, base *apps.Binding, useLocationLabel bool) *apps.Binding {
	switch len(tt) {
	case 0:
		return nil
	case 1:
		for _, t := range tt {
			base.Call = &apps.Call{
				Path: t.ID,
			}

			return base
		}
	}

	for _, t := range tt {
		label := t.DisplayName
		if useLocationLabel {
			label = t.ID
		}

		base.Bindings = append(base.Bindings, &apps.Binding{
			Location: apps.Location(t.ID),
			Label:    label,
			Icon:     utils.GetIconURL(siteURL, "now-mobile-icon.png"),
			Call: &apps.Call{
				Path: t.ID,
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
