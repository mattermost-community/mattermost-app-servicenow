package config

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-servicenow/store"
)

type config struct {
	ServiceNowInstance string
	Tables             TablesConfig
}

type LocalConfig struct {
	BaseURL        string
	MattermostURL  string
	BotAccessToken string
	BotID          string
}

func ServiceNowInstance(cc apps.Context) string {
	c := load(cc)
	return c.ServiceNowInstance
}

func SetServiceNowInstance(s string, cc apps.Context) {
	c := load(cc)
	c.ServiceNowInstance = s
	save(c, cc)
}

func save(c config, cc apps.Context) {
	dat, err := json.Marshal(c)
	if err != nil {
		log.Printf("Could not marshal config: %v", err)
		return
	}

	err = store.SaveConfig(dat, cc)
	if err != nil {
		log.Printf("Could not store config: %v", err)
	}
}

func load(cc apps.Context) config {
	defaultConfig := config{}

	data, err := store.LoadConfig(cc)
	if err != nil {
		log.Printf("Could not load config: %v", err)
		return defaultConfig
	}

	// No config stored yet
	if len(data) == 0 {
		return defaultConfig
	}

	var c config

	err = json.Unmarshal(data, &c)
	if err != nil {
		log.WithError(err).Warn("Could not unmarshal config")
		return defaultConfig
	}

	return c
}
