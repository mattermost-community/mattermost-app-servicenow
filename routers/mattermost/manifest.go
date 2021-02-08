package mattermost

import (
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
)

const (
	displayName = "Service Now"
	description = "Service Now integration"
)

func fManifest(w http.ResponseWriter, r *http.Request) {
	baseURL := config.Local().BaseURL
	manifest := api.Manifest{
		AppID:       "com.mattermost.servicenow",
		DisplayName: displayName,
		Description: description,
		HTTPRootURL: baseURL,
		HomepageURL: baseURL,
		Type:        api.AppTypeHTTP,
		RequestedLocations: api.Locations{
			api.LocationPostMenu,
			api.LocationCommand,
			api.LocationChannelHeader,
		},
	}

	utils.WriteManifest(w, manifest)
}
