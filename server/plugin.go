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

	p.router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.log.Debugf("Plugin request: not found: %q", r.URL.String())
		http.NotFound(w, r)
	})

	functionRouter := mux.NewRouter()
	function.Init("plugin", functionRouter, p.log)
	p.router.PathPrefix(apps.PluginAppPath).Handler(http.StripPrefix(apps.PluginAppPath, functionRouter))
	return nil
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
}
