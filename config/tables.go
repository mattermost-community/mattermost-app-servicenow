package config

import "github.com/mattermost/mattermost-plugin-apps/server/api"

type TablesConfig map[string]TableConfig

type TableConfig struct {
	ID          string
	DisplayName string
	Fields      []*api.Field
	Ticketable  bool
	PostDefault string
}

func AddTable(conf TableConfig) {
	if c.Tables == nil {
		c.Tables = TablesConfig{}
	}
	c.Tables[conf.ID] = conf
	save()
}

func RemoveTable(id string) {
	delete(c.Tables, id)
	save()
}

func GetTables() TablesConfig {
	return c.Tables
}
