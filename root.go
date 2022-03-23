package root

import (
	"embed" // Need to embed manifest file
	"encoding/json"

	"github.com/mattermost/mattermost-server/v6/model"

	"github.com/mattermost/mattermost-plugin-apps/apps"
)

// PluginManifestData is preloaded with the plugin manifest.
//go:embed plugin.json
var PluginManifestData []byte

// AppManifestData is preloaded with the Mattermost App manifest.
//go:embed manifest.json
var AppManifestData []byte

// Static is preloaded with the contents of the ./static directory.
//go:embed static
var Static embed.FS

var Manifest model.Manifest
var AppManifest apps.Manifest

func init() {
	err := json.Unmarshal(PluginManifestData, &Manifest)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(AppManifestData, &AppManifest)
	if err != nil {
		panic(err)
	}
}
