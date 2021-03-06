package app

import (
	"golang.org/x/oauth2"

	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/constants"
)

func GetOAuthConfig() *oauth2.Config {
	conf := config.OAuth()

	return &oauth2.Config{
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		RedirectURL:  config.Local().BaseURL + constants.OAuthPath + constants.OAuthCompletePath,
		Endpoint: oauth2.Endpoint{
			AuthURL:  config.ServiceNowInstance() + "/oauth_auth.do",
			TokenURL: config.ServiceNowInstance() + "/oauth_token.do",
		},
	}
}
