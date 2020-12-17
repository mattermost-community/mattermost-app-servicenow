package app

import (
	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"golang.org/x/oauth2"
)

func GetOAuthConfig() *oauth2.Config {
	conf := config.OAuth()
	return &oauth2.Config{
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		RedirectURL:  config.BaseURL() + constants.OAuthPath + constants.OAuthCompletePath,
		Endpoint: oauth2.Endpoint{
			AuthURL:  config.ServiceNowInstance() + "/oauth_auth.do",
			TokenURL: config.ServiceNowInstance() + "/oauth_token.do",
		},
	}
}
