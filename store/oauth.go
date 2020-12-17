package store

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/mattermost/mattermost-app-servicenow/constants"
	"golang.org/x/oauth2"
)

var states = map[string]time.Time{}
var tokens = map[string]*oauth2.Token{}

func init() {
	ticker := time.NewTicker(constants.OAuthStateGCTicker)
	go func() {
		for {
			select {
			case <-ticker.C:
				gcStates()
			}
		}
	}()
}

func gcStates() {
	for key, value := range states {
		if value.Add(constants.OAuthStateTTL).Before(time.Now()) {
			delete(states, key)
		}
	}
}

func StoreState(state string) {
	states[state] = time.Now()
}

func VerifyState(state string) bool {
	_, found := states[state]
	return found
}

func StoreToken(token *oauth2.Token, userID string) {
	tokens[userID] = token
	storeTokens()
}

func GetToken(userID string) (*oauth2.Token, bool) {
	token, found := tokens[userID]
	return token, found
}

func DeleteToken(userID string) {
	delete(tokens, userID)
	storeTokens()
}

func storeTokens() {
	dat, err := json.Marshal(tokens)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(constants.TokenFile, dat, 0644)
	if err != nil {
		return
	}
}

func LoadTokens() {
	dat, err := ioutil.ReadFile(constants.TokenFile)
	if err != nil {
		return
	}

	err = json.Unmarshal(dat, &tokens)
	if err != nil {
		return
	}
}
