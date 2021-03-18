package store

import (
	"log"
	"time"

	"golang.org/x/oauth2"

	"github.com/mattermost/mattermost-app-servicenow/clients/mattermostclient"
)

func SaveState(botAccessToken, baseURL, state, botID string) {
	states := loadStates(botAccessToken, baseURL, botID)
	states[state] = time.Now()
	saveStates(botAccessToken, baseURL, botID, states)
}

func VerifyState(botAccessToken, baseURL, state, botID string) bool {
	states := loadStates(botAccessToken, baseURL, botID)
	_, found := states[state]

	return found
}

func saveStates(botAccessToken, baseURL, botID string, states map[string]time.Time) {
	c := mattermostclient.NewKVClient(botAccessToken, baseURL, botID)

	err := c.KVSet("states", states)
	if err != nil {
		log.Printf("Could not store states: %v", err)
		return
	}
}

func loadStates(botAccessToken, baseURL, botID string) map[string]time.Time {
	defaultStates := map[string]time.Time{}
	states := map[string]time.Time{}
	c := mattermostclient.NewKVClient(botAccessToken, baseURL, botID)

	err := c.KVGet("states", &states)
	if err != nil {
		log.Printf("Could not get states: %v", err)
		return defaultStates
	}

	return states
}

func SaveToken(botAccessToken, baseURL, botID string, token *oauth2.Token, userID string) {
	tokens := loadTokens(botAccessToken, baseURL, botID)
	tokens[userID] = token
	saveTokens(botAccessToken, baseURL, botID, tokens)
}

func GetToken(botAccessToken, baseURL, botID string, userID string) (*oauth2.Token, bool) {
	tokens := loadTokens(botAccessToken, baseURL, botID)
	token, found := tokens[userID]

	return token, found
}

func DeleteToken(botAccessToken, baseURL, botID string, userID string) {
	tokens := loadTokens(botAccessToken, baseURL, botID)
	delete(tokens, userID)
	saveTokens(botAccessToken, baseURL, botID, tokens)
}

func saveTokens(botAccessToken, baseURL, botID string, tokens map[string]*oauth2.Token) {
	c := mattermostclient.NewKVClient(botAccessToken, baseURL, botID)

	err := c.KVSet("tokens", tokens)
	if err != nil {
		log.Printf("Could not store tokens: %v", err)
		return
	}
}

func loadTokens(botAccessToken, baseURL, botID string) map[string]*oauth2.Token {
	defaultTokens := map[string]*oauth2.Token{}
	tokens := map[string]*oauth2.Token{}
	c := mattermostclient.NewKVClient(botAccessToken, baseURL, botID)

	err := c.KVGet("tokens", &tokens)
	if err != nil {
		log.Printf("Could not get tokens: %v", err)
		return defaultTokens
	}

	return tokens
}
