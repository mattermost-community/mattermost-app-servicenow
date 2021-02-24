package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

const (
	displayName = "Service Now"
	description = "Service Now integration"
)

func fManifest(m *apps.Manifest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.WriteManifest(w, *m)
	}
}
