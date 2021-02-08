package config

import "github.com/mattermost/mattermost-plugin-apps/server/api"

type TablesConfig map[string]TableConfig

type TableConfig struct {
	ID          string
	DisplayName string
	Fields      []*api.Field
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

	c.Tables["incidents"] = TableConfig{
		ID:          "incidents",
		DisplayName: "Incidents",
		Fields: []*api.Field{
			{
				Name:       "short_description",
				ModalLabel: "Short Description",
				Type:       api.FieldTypeText,
			},
			{
				Name:       "description",
				ModalLabel: "Long Description",
				Type:       api.FieldTypeText,
			},
		},
		Post:        true,
		Command:     true,
		Header:      true,
		PostDefault: "short_description",
	}

	return c.Tables
}
