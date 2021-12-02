package app

import (
	"encoding/json"

	"golang.org/x/oauth2"

	"github.com/mattermost/mattermost-app-servicenow/config"

	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func GetOAuthConfig(cc apps.Context) *oauth2.Config {
	instance := config.ServiceNowInstance(cc)

	return &oauth2.Config{
		ClientID:     cc.OAuth2.ClientID,
		ClientSecret: cc.OAuth2.ClientSecret,
		RedirectURL:  cc.OAuth2.CompleteURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  instance + "/oauth_auth.do",
			TokenURL: instance + "/oauth_token.do",
		},
	}
}

func GetTokenFromContext(cc apps.Context) *oauth2.Token {
	token := &oauth2.Token{}

	b, err := json.Marshal(cc.OAuth2.User)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(b, &token)
	if err != nil {
		return nil
	}

	return token
}
