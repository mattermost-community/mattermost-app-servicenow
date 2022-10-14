package function

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-servicenow/goapp"
)

const (
	FieldShortDescription = "short_description"
	FieldLongDescription  = "description"
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
	out := Tables{}
	for k, v := range getAppConfig(creq).Tables {
		out[k] = v
	}

	out["incident"] = Table{
		ID:          "incident",
		DisplayName: "Incidents",
		Fields: []*apps.Field{
			{
				Name:        FieldShortDescription,
				Label:       FieldShortDescription,
				ModalLabel:  "Short Description",
				Type:        apps.FieldTypeText,
				TextSubtype: apps.TextFieldSubtypeTextarea,
				IsRequired:  true,
			},
			{
				Name:        FieldLongDescription,
				Label:       FieldLongDescription,
				ModalLabel:  "Long Description",
				Type:        apps.FieldTypeText,
				TextSubtype: apps.TextFieldSubtypeTextarea,
			},
		},
		BindTo:        []apps.Location{apps.LocationCommand, apps.LocationChannelHeader, apps.LocationPostMenu},
		PostFieldName: "short_description",
	}

	return out
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
