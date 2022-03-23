package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap/zapcore"

	"github.com/mattermost/mattermost-plugin-apps/utils"

	"github.com/mattermost/mattermost-app-servicenow/function"
)

func main() {
	r := mux.NewRouter()
	log := utils.MustMakeCommandLogger(zapcore.DebugLevel)
	function.Init("http", r, log)
	http.Handle("/", r)

	listen := ":4445"
	log.Infof("servicenow app started, manifest at http://localhost%s/manifest.json", listen)
	panic(http.ListenAndServe(listen, nil))
}
