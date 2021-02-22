package config

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

type TablesConfig map[string]TableConfig

type TableConfig struct {
	ID          string
	DisplayName string
	Fields      []*apps.Field
	Post        bool
	Command     bool
	Header      bool
	PostDefault string
}

func AddTable(conf TableConfig) {
	c := load()
	if c.Tables == nil {
		c.Tables = TablesConfig{}
	}

	c.Tables[conf.ID] = conf
	save(c)
}

func RemoveTable(id string) {
	c := load()
	delete(c.Tables, id)
	save(c)
}

func GetTables() TablesConfig {
	c := load()

	// Remove when Add and Remove table functionality is present. Adds default table.
	if c.Tables == nil {
		c.Tables = TablesConfig{}
	}

	c.Tables["incident"] = TableConfig{
		ID:          "incident",
		DisplayName: "Incidents",
		Fields: []*apps.Field{
			{
				Name:       "short_description",
				ModalLabel: "Short Description",
				Label:      "short_description",
				Type:       apps.FieldTypeText,
			},
			{
				Name:       "description",
				Label:      "description",
				ModalLabel: "Long Description",
				Type:       apps.FieldTypeText,
			},
		},
		Post:        true,
		Command:     true,
		Header:      true,
		PostDefault: "short_description",
	}

	return c.Tables
}
