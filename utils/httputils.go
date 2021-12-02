// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func WriteCallResponse(w http.ResponseWriter, v apps.CallResponse) {
	writeJSON(w, v)
}

func WriteBindings(w http.ResponseWriter, v []apps.Binding) {
	call := apps.CallResponse{
		Type: apps.CallResponseTypeOK,
		Data: v,
	}
	writeJSON(w, call)
}

func WriteManifest(w http.ResponseWriter, v apps.Manifest) {
	writeJSON(w, v)
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func WriteBadRequestError(w http.ResponseWriter, err error) {
	WriteCallResponse(w, newCallErrorResponse(fmt.Sprintf("Invalid request. Error: %s", err.Error())))
}

func WriteCallErrorResponse(w http.ResponseWriter, message string) {
	WriteCallResponse(w, newCallErrorResponse(message))
}

func WriteCallStandardResponse(w http.ResponseWriter, message string) {
	WriteCallResponse(w, newCallStandardResponse(message))
}

func newCallStandardResponse(message string) apps.CallResponse {
	return apps.CallResponse{
		Type:     apps.CallResponseTypeOK,
		Markdown: message,
	}
}

func newCallErrorResponse(message string) apps.CallResponse {
	return apps.CallResponse{
		Type:      apps.CallResponseTypeError,
		ErrorText: message,
	}
}
