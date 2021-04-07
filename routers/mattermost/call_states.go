package mattermost

import (
	"encoding/json"
)

type createTicketCallState struct {
	Action formAction `json:"action"`
	Table  string     `json:"table"`
}

func (s *createTicketCallState) FromState(i interface{}) {
	if i == nil {
		return
	}

	b, err := json.Marshal(i)
	if err != nil {
		return
	}

	_ = json.Unmarshal(b, &s)
}

type configureOAuthCallState struct {
	Action formAction `json:"action"`
}

func (s *configureOAuthCallState) FromState(i interface{}) {
	if i == nil {
		return
	}

	b, err := json.Marshal(i)
	if err != nil {
		return
	}

	_ = json.Unmarshal(b, &s)
}
