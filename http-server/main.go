package main

import (
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap/zapcore"

	"github.com/mattermost/mattermost-plugin-apps/utils"

	root "github.com/mattermost/mattermost-app-servicenow"
	"github.com/mattermost/mattermost-app-servicenow/function"
)

func main() {
	rootURL := os.Getenv("ROOT_URL")
	if rootURL != "" {
		root.AppManifest.Deploy.HTTP.RootURL = rootURL
	}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		u, err := url.Parse(root.AppManifest.Deploy.HTTP.RootURL)
		if err != nil {
			panic(err)
		}
		portStr = u.Port()
		if portStr == "" {
			portStr = "8080"
		}
	}

	r := mux.NewRouter()
	log := utils.MustMakeCommandLogger(zapcore.DebugLevel)
	function.Init("http", r, log)

	listen := ":" + portStr
	server := &http.Server{
		Addr:              listen,
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      5 * time.Second,
		Handler:           r,
	}

	log.Infof("servicenow app started, listening on port %s, manifest at %s/manifest.json", portStr, root.AppManifest.Deploy.HTTP.RootURL)
	panic(server.ListenAndServe())
}
