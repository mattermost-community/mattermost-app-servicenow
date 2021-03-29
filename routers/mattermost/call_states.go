package mattermost

import (
	"encoding/json"
)

type CreateTicketCallState struct {
	Action formAction `json:"action"`
	Table  string     `json:"table"`
}

func (s *CreateTicketCallState) FromState(i interface{}) {
	if i == nil {
		return
	}

	b, err := json.Marshal(i)
	if err != nil {
		return
	}

	_ = json.Unmarshal(b, &s)
}

type ConfigureOAuthCallState struct {
	Action formAction `json:"action"`
}

func (s *ConfigureOAuthCallState) FromState(i interface{}) {
	if i == nil {
		return
	}

	b, err := json.Marshal(i)
	if err != nil {
		return
	}

	_ = json.Unmarshal(b, &s)
}
