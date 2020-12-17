package store

import (
	"io/ioutil"

	"github.com/mattermost/mattermost-app-servicenow/constants"
)

func StoreConfig(conf []byte) error {
	err := ioutil.WriteFile(constants.ConfigFile, conf, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfig() ([]byte, error) {
	dat, err := ioutil.ReadFile(constants.ConfigFile)
	if err != nil {
		return nil, err
	}
	return dat, nil
}
