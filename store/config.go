package store

import (
	"io/ioutil"

	"github.com/mattermost/mattermost-app-servicenow/clients/mattermostclient"
	"github.com/mattermost/mattermost-app-servicenow/constants"
)

func SaveConfig(conf []byte, botAccessToken, baseURL, botID string) error {
	c := mattermostclient.NewKVClient(botAccessToken, baseURL, botID)
	err := c.KVSet("config", map[string]interface{}{"config": conf})

	if err != nil {
		return err
	}

	return nil
}

func LoadConfig(botAccessToken, baseURL, botID string) ([]byte, error) {
	c := mattermostclient.NewKVClient(botAccessToken, baseURL, botID)
	stored := map[string][]byte{}

	err := c.KVGet("config", &stored)
	if err != nil {
		return nil, err
	}

	if conf, ok := stored["config"]; ok {
		return conf, nil
	}

	return nil, nil
}

func SaveLocalConfig(conf []byte) error {
	err := ioutil.WriteFile(constants.ConfigFile, conf, 0600)
	if err != nil {
		return err
	}

	return nil
}

func LoadLocalConfig() ([]byte, error) {
	dat, err := ioutil.ReadFile(constants.ConfigFile)
	if err != nil {
		return nil, err
	}

	return dat, nil
}
