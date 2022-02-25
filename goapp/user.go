package goapp

import (
	"golang.org/x/oauth2"
)

type User struct {
	MattermostID string
	RemoteID     string
	Token        *oauth2.Token
}
