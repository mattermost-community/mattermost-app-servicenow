package config

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/appclient"
)

func SetOAuthConfig(clientID, clientSecret string, cc apps.Context) error {
	c := appclient.AsActingUser(cc)

	err := c.StoreOAuth2App(apps.OAuth2App{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	})
	if err != nil {
		return err
	}

	return nil
}
