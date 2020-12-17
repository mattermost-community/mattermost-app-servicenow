package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/server/apps"
)

const (
	displayName = "Service Now"
	description = "Service Now integration"
)

func fManifest(w http.ResponseWriter, r *http.Request) {
	manifest := apps.Manifest{
		AppID:       "com.mattermost.servicenow",
		DisplayName: displayName,
		Description: description,
		RootURL:     config.BaseURL(),
		RequestedLocations: apps.Locations{
			apps.LocationPostMenu,
			apps.LocationCommand,
		},
	}

	utils.WriteManifest(w, manifest)
}
