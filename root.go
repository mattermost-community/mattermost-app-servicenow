package root

import (
	"embed" // Need to embed manifest file
	"encoding/json"

	"github.com/mattermost/mattermost-server/v6/model"

	"github.com/mattermost/mattermost-plugin-apps/apps"
)

// pluginManifestData is preloaded with the plugin manifest.
//
//go:embed plugin.json
var pluginManifestData []byte

// appManifestData is preloaded with the Mattermost App manifest.
//
//go:embed manifest.json
var appManifestData []byte

// Static is preloaded with the contents of the ./static directory.
//
//go:embed static
var Static embed.FS

var Manifest model.Manifest
var AppManifest apps.Manifest

func init() {
	err := json.Unmarshal(pluginManifestData, &Manifest)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(appManifestData, &AppManifest)
	if err != nil {
		panic(err)
	}
}
