package store

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/mmclient"
)

func SaveConfig(conf []byte, cc *apps.Context) error {
	c := mmclient.AsBot(cc)
	_, err := c.KVSet("config", "", map[string]interface{}{"config": conf})

	if err != nil {
		return err
	}

	return nil
}

func LoadConfig(cc *apps.Context) ([]byte, error) {
	c := mmclient.AsBot(cc)
	stored := map[string][]byte{}

	err := c.KVGet("config", "", &stored)
	if err != nil {
		return nil, err
	}

	if conf, ok := stored["config"]; ok {
		return conf, nil
	}

	return nil, nil
}
