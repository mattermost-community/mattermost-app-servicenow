package goapp

import (
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/utils"
	"github.com/mattermost/mattermost-plugin-apps/utils/httputils"
)

type App struct {
	Logger utils.Logger
	Router *mux.Router
	Icon   string
}

func NewApp(r *mux.Router, log utils.Logger) *App {
	// Ping.
	r.Path("/ping").HandlerFunc(httputils.DoHandleJSONData([]byte("{}")))

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("App request: not found: %q", r.URL.String())
		http.NotFound(w, r)
	})

	return &App{
		Router: r,
		Logger: log,
	}
}

func (a *App) WithManifest(m apps.Manifest) *App {
	a.Router.Path("/manifest.json").HandlerFunc(httputils.DoHandleJSON(m)).Methods("GET")
	return a
}

func (a *App) WithStatic(staticFS fs.FS) *App {
	a.Router.PathPrefix("/static/").Handler(http.FileServer(http.FS(staticFS)))
	return a
}

func (a App) WithIcon(iconPath string) *App {
	a.Icon = iconPath
	return &a
}

func AppendBinding(bb []apps.Binding, b *apps.Binding) []apps.Binding {
	var out []apps.Binding
	out = append(out, bb...)
	if b != nil {
		out = append(out, *b)
	}
	return out
}
