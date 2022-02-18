package main

import (
	"net/http"

	"github.com/gorilla/mux"
	pluginapi "github.com/mattermost/mattermost-plugin-api"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/utils"
	"github.com/mattermost/mattermost-server/v6/plugin"

	root "github.com/mattermost/mattermost-app-servicenow"
	"github.com/mattermost/mattermost-app-servicenow/function"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	router *mux.Router
}

func (p *Plugin) OnActivate() error {
	p.router = mux.NewRouter()
	function.Mode = "plugin"

	function.InitHTTP(
		utils.NewPluginLogger(pluginapi.NewClient(p.API, p.Driver)),
		root.NewAppRoute(p.router, apps.PluginAppPath),
	)
	return nil
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	p.API.LogDebug("HTTP", "method", r.Method, "url", r.URL.Path)
	p.router.ServeHTTP(w, r)
}
