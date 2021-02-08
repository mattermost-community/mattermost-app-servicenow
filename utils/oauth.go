package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

const stateLength = 3

var ErrBadState = errors.New("bad state")

func CreateOAuthState(userID, channelID string) string {
	return fmt.Sprintf("%v_%v_%v", model.NewId()[0:15], userID, channelID)
}

func ParseOAuthState(state string) (string, string, error) {
	splitted := strings.Split(state, "_")
	if len(splitted) != stateLength {
		return "", "", ErrBadState
	}

	return splitted[1], splitted[2], nil
}
