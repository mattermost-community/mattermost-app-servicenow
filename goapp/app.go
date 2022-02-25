package goapp

import (
	"encoding/json"
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

func (a *App) HandleCall(p string, h HandlerFunc) {
	a.Router.Path(p).HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		creq := CallRequest{
			GoContext: req.Context(),
		}
		err := json.NewDecoder(req.Body).Decode(&creq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		creq.App = *a
		creq.App.Logger = a.Logger.With("path", creq.Path)

		cresp := h(creq)
		if cresp.Type == apps.CallResponseTypeError {
			creq.App.Logger.WithError(cresp).Debugw("Call failed.")
		}
		_ = httputils.WriteJSON(w, cresp)
	})
}

func AppendBindings(bb1, bb2 []apps.Binding) []apps.Binding {
	var out []apps.Binding
	if len(bb1) != 0 {
		out = append(out, bb1...)
	}
	if len(bb2) != 0 {
		out = append(out, bb2...)
	}
	return out
}

func AppendBinding(bb []apps.Binding, b *apps.Binding) []apps.Binding {
	var out []apps.Binding
	out = append(out, bb...)
	if b != nil {
		out = append(out, *b)
	}
	return out
}
