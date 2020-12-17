package config

import (
	"encoding/json"

	"github.com/mattermost/mattermost-app-servicenow/store"
)

type config struct {
	BaseURL            string
	ServiceNowInstance string
	OAuth              OAuthConfig
	Mattermost         MattermostConfig
	Tables             TablesConfig
}

type OAuthConfig struct {
	ClientID     string
	ClientSecret string
}

type MattermostConfig struct {
	MattermostURL  string
	BotID          string
	BotAccessToken string
}

var c = config{
	BaseURL: "http://localhost:3000",
}

func BaseURL() string {
	return c.BaseURL
}

func ServiceNowInstance() string {
	return c.ServiceNowInstance
}

func OAuth() OAuthConfig {
	return c.OAuth
}

func Mattermost() MattermostConfig {
	return c.Mattermost
}

func SetBaseURL(s string) {
	c.BaseURL = s
	save()
}

func SetServiceNowInstance(s string) {
	c.ServiceNowInstance = s
	save()
}

func SetOAuthConfig(v OAuthConfig) {
	c.OAuth = v
	save()
}

func SetMattermostConfig(v MattermostConfig) {
	c.Mattermost = v
	save()
}

func Load() {
	dat, err := store.LoadConfig()
	err = json.Unmarshal(dat, &c)
	if err != nil {
		return
	}
}

func save() {
	dat, err := json.Marshal(c)
	if err != nil {
		return
	}
	err = store.StoreConfig(dat)
}
