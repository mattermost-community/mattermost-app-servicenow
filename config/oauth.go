package config

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/mmclient"
)

func SetOAuthConfig(clientID, clientSecret string, cc *apps.Context) error {
	c := mmclient.AsActingUser(cc)

	err := c.StoreOAuth2App(cc.AppID, clientID, clientSecret)
	if err != nil {
		return err
	}

	return nil
}
