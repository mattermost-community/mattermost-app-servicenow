package config

import (
	"encoding/json"
	"log"

	"github.com/mattermost/mattermost-app-servicenow/store"
)

type config struct {
	ServiceNowInstance string
	OAuth              OAuthConfig
	Tables             TablesConfig
}

type OAuthConfig struct {
	ClientID     string
	ClientSecret string
}

type LocalConfig struct {
	BaseURL        string
	MattermostURL  string
	BotAccessToken string
}

func ServiceNowInstance() string {
	c := load()
	return c.ServiceNowInstance
}

func OAuth() OAuthConfig {
	c := load()
	return c.OAuth
}

func Local() LocalConfig {
	return loadLocal()
}

func SetBaseURL(s string) {
	c := loadLocal()
	c.BaseURL = s
	saveLocal(c)
}

func SetServiceNowInstance(s string) {
	c := load()
	c.ServiceNowInstance = s
	save(c)
}

func SetOAuthConfig(v OAuthConfig) {
	c := load()
	c.OAuth = v
	save(c)
}

func SetLocalConfig(v LocalConfig) {
	saveLocal(v)
}

func save(c config) {
	lc := loadLocal()

	dat, err := json.Marshal(c)
	if err != nil {
		log.Printf("Could not marshal config: %v", err)
		return
	}

	err = store.SaveConfig(dat, lc.BotAccessToken, lc.MattermostURL)
	if err != nil {
		log.Printf("Could not store config: %v", err)
	}
}

func load() config {
	defaultConfig := config{}
	lc := loadLocal()

	dat, err := store.LoadConfig(lc.BotAccessToken, lc.MattermostURL)
	if err != nil {
		log.Printf("Could not load config: %v", err)
		return defaultConfig
	}

	c := config{}

	err = json.Unmarshal(dat, &c)
	if err != nil {
		log.Printf("Could not unmarshal config: %v", err)
		return defaultConfig
	}

	return c
}

func loadLocal() LocalConfig {
	defaultLocalConfig := LocalConfig{
		BaseURL: "http://localhost:3000",
	}

	dat, err := store.LoadLocalConfig()
	if err != nil {
		log.Printf("Could not load local config: %v", err)
		return defaultLocalConfig
	}

	c := LocalConfig{}

	err = json.Unmarshal(dat, &c)
	if err != nil {
		log.Printf("Could not unmarshall local config: %v", err)
		return defaultLocalConfig
	}

	return c
}

func saveLocal(c LocalConfig) {
	dat, err := json.Marshal(c)
	if err != nil {
		log.Printf("Could not marshal local config: %v", err)
		return
	}

	err = store.SaveLocalConfig(dat)
	if err != nil {
		log.Printf("Could not store local config: %v", err)
	}
}
