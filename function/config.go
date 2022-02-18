package function

import (
	"github.com/mattermost/mattermost-plugin-apps/utils"

	"github.com/mattermost/mattermost-app-servicenow/goapp"
)

type Config struct {
	RemoteURL string
	Tables    Tables
}

func appConfig(creq goapp.CallRequest) Config {
	conf := Config{}
	utils.Remarshal(&conf, creq.Context.OAuth2.Data)
	return conf
}
