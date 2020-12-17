package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

func CreateOAuthState(userID, channelID string) string {
	return fmt.Sprintf("%v_%v_%v", model.NewId()[0:15], userID, channelID)
}

func ParseOAuthState(state string) (string, string, error) {
	splitted := strings.Split(state, "_")
	if len(splitted) != 3 {
		return "", "", errors.New("Bad state")
	}

	return splitted[1], splitted[2], nil
}
