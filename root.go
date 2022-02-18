package root

import (
	"embed" // Need to embed manifest file
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-server/v6/model"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/utils/httputils"
)

//go:embed plugin.json
var pluginManifestData []byte

//go:embed manifest.json
var appManifestData []byte

//go:embed static
var staticFS embed.FS

var Manifest model.Manifest
var AppManifest apps.Manifest

func init() {
	_ = json.Unmarshal(pluginManifestData, &Manifest)
	_ = json.Unmarshal(appManifestData, &AppManifest)
}

func NewAppRoute(r *mux.Router, prefix string) *mux.Router {
	r = r.Path(prefix).Subrouter()
	// router.Path(prefix + "/manifest.json").HandlerFunc(httputils.DoHandleJSONData(appManifestData)).Methods("GET")
	r.Path("/manifest.json").HandlerFunc(httputils.DoHandleJSONData(appManifestData)).Methods("GET")
	r.Path("/static/").Handler(http.StripPrefix(prefix+"/", http.FileServer(http.FS(staticFS))))
	return r
}
