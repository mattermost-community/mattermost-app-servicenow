package store

import (
	"log"
	"time"

	"github.com/mattermost/mattermost-app-servicenow/clients/mattermostclient"
	"golang.org/x/oauth2"
)

func SaveState(botAccessToken, baseURL, state string) {
	states := loadStates(botAccessToken, baseURL)
	states[state] = time.Now()
	saveStates(botAccessToken, baseURL, states)
}

func VerifyState(botAccessToken, baseURL, state string) bool {
	states := loadStates(botAccessToken, baseURL)
	_, found := states[state]

	return found
}

func saveStates(botAccessToken, baseURL string, states map[string]time.Time) {
	c := mattermostclient.NewKVClient(botAccessToken, baseURL)

	err := c.KVSet("states", states)
	if err != nil {
		log.Printf("Could not store states: %v", err)
		return
	}
}

func loadStates(botAccessToken, baseURL string) map[string]time.Time {
	defaultStates := map[string]time.Time{}
	states := map[string]time.Time{}
	c := mattermostclient.NewKVClient(botAccessToken, baseURL)

	err := c.KVGet("states", &states)
	if err != nil {
		log.Printf("Could not get states: %v", err)
		return defaultStates
	}

	return states
}

func SaveToken(botAccessToken, baseURL string, token *oauth2.Token, userID string) {
	tokens := loadTokens(botAccessToken, baseURL)
	tokens[userID] = token
	saveTokens(botAccessToken, baseURL, tokens)
}

func GetToken(botAccessToken, baseURL string, userID string) (*oauth2.Token, bool) {
	tokens := loadTokens(botAccessToken, baseURL)
	token, found := tokens[userID]

	return token, found
}

func DeleteToken(botAccessToken, baseURL string, userID string) {
	tokens := loadTokens(botAccessToken, baseURL)
	delete(tokens, userID)
	saveTokens(botAccessToken, baseURL, tokens)
}

func saveTokens(botAccessToken, baseURL string, tokens map[string]*oauth2.Token) {
	c := mattermostclient.NewKVClient(botAccessToken, baseURL)

	err := c.KVSet("tokens", tokens)
	if err != nil {
		log.Printf("Could not store tokens: %v", err)
		return
	}
}

func loadTokens(botAccessToken, baseURL string) map[string]*oauth2.Token {
	defaultTokens := map[string]*oauth2.Token{}
	tokens := map[string]*oauth2.Token{}
	c := mattermostclient.NewKVClient(botAccessToken, baseURL)

	err := c.KVGet("tokens", &tokens)
	if err != nil {
		log.Printf("Could not get tokens: %v", err)
		return defaultTokens
	}

	return tokens
}
