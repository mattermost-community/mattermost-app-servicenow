package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

const (
	displayName = "Service Now"
	description = "Service Now integration"
)

func fManifest(w http.ResponseWriter, r *http.Request) {
	baseURL := config.Local().BaseURL
	manifest := apps.Manifest{
		AppID:       "com.mattermost.servicenow",
		DisplayName: displayName,
		Description: description,
		HTTPRootURL: baseURL,
		HomepageURL: baseURL,
		Type:        apps.AppTypeHTTP,
		RequestedLocations: apps.Locations{
			apps.LocationPostMenu,
			apps.LocationCommand,
			apps.LocationChannelHeader,
		},
	}

	utils.WriteManifest(w, manifest)
}
