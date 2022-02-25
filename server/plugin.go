package main

import (
	"net/http"

	"github.com/gorilla/mux"
	pluginapi "github.com/mattermost/mattermost-plugin-api"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/utils"
	"github.com/mattermost/mattermost-server/v6/plugin"

	"github.com/mattermost/mattermost-app-servicenow/function"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	router *mux.Router
	client *pluginapi.Client
	log    utils.Logger
}

func (p *Plugin) OnActivate() error {
	p.router = mux.NewRouter()
	p.client = pluginapi.NewClient(p.API, p.Driver)
	p.log = utils.NewPluginLogger(p.client)

	functionRouter := p.router.Path(apps.PluginAppPath).Subrouter()
	function.Init("plugin", functionRouter, p.log)
	return nil
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	p.API.LogDebug("HTTP", "method", r.Method, "url", r.URL.Path)
	p.router.ServeHTTP(w, r)
}
