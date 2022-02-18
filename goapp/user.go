package goapp

import (
	"golang.org/x/oauth2"
)

type User struct {
	Token *oauth2.Token
	ID    string
}
