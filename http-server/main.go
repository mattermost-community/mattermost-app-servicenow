package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.uber.org/zap/zapcore"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/utils"

	root "github.com/mattermost/mattermost-app-servicenow"
	"github.com/mattermost/mattermost-app-servicenow/function"
)

func main() {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "4445"
	}

	rootURL := os.Getenv("ROOT_URL")
	if rootURL == "" {
		rootURL = "http://localhost:" + portStr
	}
	root.AppManifest.Deploy.HTTP = &apps.HTTP{
		RootURL: rootURL,
	}

	r := mux.NewRouter()
	log := utils.MustMakeCommandLogger(zapcore.DebugLevel)
	function.Init("http", r, log)
	http.Handle("/", r)

	listen := ":" + portStr
	log.Infof("servicenow app started, listening on port %s, manifest at %s/manifest.json", portStr, rootURL)
	panic(http.ListenAndServe(listen, nil))
}
