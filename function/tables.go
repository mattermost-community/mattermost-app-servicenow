package function

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-servicenow/goapp"
)

type Tables map[string]Table

type Table struct {
	// Table ID and name
	ID          string
	DisplayName string

	// Fields to include in the form to create tickets
	Fields []*apps.Field

	// BindTo controls what top-level bindings the table is exposed in.
	// PostFieldName indicates what text field is to be defaulted to the post's
	// message if invoked from the post menu.
	BindTo        apps.Locations
	PostFieldName string
}

func UpdateTable(creq goapp.CallRequest, t Table) error {
	return nil
}

func RemoveTable(creq goapp.CallRequest, id string) error {
	return nil
}

func GetTables(creq goapp.CallRequest) Tables {
	return Tables{
		"incident": Table{
			ID:          "incident",
			DisplayName: "Incidents",
			Fields: []*apps.Field{
				{
					Name:        "short_description",
					ModalLabel:  "Short Description",
					Label:       "short_description",
					Type:        apps.FieldTypeText,
					TextSubtype: apps.TextFieldSubtypeTextarea,
				},
				{
					Name:        "description",
					Label:       "description",
					ModalLabel:  "Long Description",
					Type:        apps.FieldTypeText,
					TextSubtype: apps.TextFieldSubtypeTextarea,
				},
			},
			BindTo:        []apps.Location{apps.LocationCommand, apps.LocationChannelHeader, apps.LocationPostMenu},
			PostFieldName: "short_description",
		},
	}
}

func (t Tables) forLocation(root apps.Location) Tables {
	out := Tables{}
	for _, table := range t {
		for _, loc := range table.BindTo {
			if loc == root {
				out[table.ID] = table
			}
		}
	}
	return out
}
