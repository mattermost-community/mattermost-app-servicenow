// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/mattermost/mattermost-plugin-apps/server/api"
	"github.com/mattermost/mattermost-plugin-apps/server/utils/md"
)

func NormalizeRemoteBaseURL(mattermostSiteURL, remoteURL string) (string, error) {
	u, err := url.Parse(remoteURL)
	if err != nil {
		return "", err
	}
	if u.Host == "" {
		ss := strings.Split(u.Path, "/")
		if len(ss) > 0 && ss[0] != "" {
			u.Host = ss[0]
			u.Path = path.Join(ss[1:]...)
		}
		u, err = url.Parse(u.String())
		if err != nil {
			return "", err
		}
	}
	if u.Host == "" {
		return "", fmt.Errorf("invalid URL, no hostname: %q", remoteURL)
	}
	if u.Scheme == "" {
		u.Scheme = "https"
	}

	remoteURL = strings.TrimSuffix(u.String(), "/")
	if remoteURL == strings.TrimSuffix(mattermostSiteURL, "/") {
		return "", fmt.Errorf("%s is the Mattermost site URL. Please use the remote application's URL", remoteURL)
	}

	return remoteURL, nil
}

func WriteCallResponse(w http.ResponseWriter, v api.CallResponse) {
	writeJSON(w, v)
}

func WriteBindings(w http.ResponseWriter, v []*api.Binding) {
	call := api.CallResponse{
		Type: api.CallResponseTypeOK,
		Data: v,
	}
	writeJSON(w, call)
}

func WriteManifest(w http.ResponseWriter, v api.Manifest) {
	writeJSON(w, v)
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func WriteInternalServerError(w http.ResponseWriter) {
	WriteCallResponse(w, newCallErrorResponse("An internal error has occurred. Check app server logs for details."))
}

func WriteBadRequestError(w http.ResponseWriter, err error) {
	WriteCallResponse(w, newCallErrorResponse(fmt.Sprintf("Invalid request. Error: %s", err.Error())))
}

func WriteNotFoundError(w http.ResponseWriter) {
	WriteCallResponse(w, newCallErrorResponse("Not found."))
}

func WriteUnauthorizedError(w http.ResponseWriter) {
	WriteCallErrorResponse(w, "Unauthorized")
}

func WriteCallErrorResponse(w http.ResponseWriter, message string) {
	WriteCallResponse(w, newCallErrorResponse(message))
}

func WriteCallStandardResponse(w http.ResponseWriter, message string) {
	WriteCallResponse(w, newCallStandardResponse(message))
}

func newCallStandardResponse(message string) api.CallResponse {
	return api.CallResponse{
		Type:     api.CallResponseTypeOK,
		Markdown: md.MD(message),
	}
}

func newCallErrorResponse(message string) api.CallResponse {
	return api.CallResponse{
		Type:      api.CallResponseTypeError,
		ErrorText: message,
	}
}
