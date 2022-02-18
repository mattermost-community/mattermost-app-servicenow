package function

import (
	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-app-servicenow/goapp"
	"github.com/mattermost/mattermost-plugin-apps/utils"
	"github.com/mattermost/mattermost-plugin-apps/utils/httputils"
)

const servicenow = "servicenow"
const ServiceNow = "ServiceNow"
const icon = "icon.png"

var Mode = ""

var BuildHash string
var BuildHashShort string
var BuildDate string

// KV store
const (
// KVTODOPrefix = TODO
)

// Field names
const (
	fClientID     = "client_id"
	fClientSecret = "client_secret"
	fID           = "id"
	fState        = "state"
	fCode         = "code"
	fTable        = "table"
)

type CallRequest struct {
	goapp.CallRequest
}

func InitHTTP(log utils.Logger, r *mux.Route) {
	// Ping.
	r.Path("/ping").HandlerFunc(httputils.DoHandleJSONData([]byte("{}")))

	a := goapp.App{
		Logger: log,
		Route:  r,
		Icon:   icon,
	}

	// Bindings.
	a.HandleCall("/bindings", bindings)

	// OAuth2 (Google Calendar) connect commands and callbacks.
	a.HandleCall("/oauth2/connect", oauth2Connect)
	a.HandleCall("/oauth2/complete", oauth2Complete)

	// func Init(router *mux.Router, m *apps.Manifest, staticAssets fs.FS, localMode bool) {
	// 	router.HandleFunc(constants.ManifestPath, fManifest(m))
	a.HandleCall("/install", install)
	// 	router.HandleFunc(constants.BindingsPath, extractCall(fBindings, localMode))
	// 	router.HandleFunc(constants.OAuthPath+constants.OAuthConnectPath, extractCall(fOAuthConnect, localMode))
	// 	router.HandleFunc(constants.OAuthPath+constants.OAuthCompletePath, extractCall(fOAuthComplete, localMode))

	// 	router.HandleFunc(constants.BindingPathCreate.Submit(), extractCall(fCreateTicketSubmit, localMode))
	// 	router.HandleFunc(constants.BindingPathCreate.Form(), extractCall(fCreateTicketForm, localMode))

	// 	router.HandleFunc(constants.BindingPathConfigureOAuth.Submit(), extractCall(fConfigureOAuthSubmit, localMode))
	// 	router.HandleFunc(constants.BindingPathConfigureOAuth.Form(), extractCall(fConfigureOAuthForm, localMode))

	// 	router.HandleFunc(constants.BindingPathConnect.Submit(), extractCall(fConnect, localMode))
	// 	router.HandleFunc(constants.BindingPathDisconnect.Submit(), extractCall(fDisconnect, localMode))

	// 	router.PathPrefix(constants.StaticAssetPath).Handler(http.StripPrefix("/", http.FileServer(http.FS(staticAssets))))
	// }

	// Command submit handlers.
	// HandleCommand(configure)
	// HandleCommand(connect)
	// HandleCommand(debugGetEvent)
	// HandleCommand(debugListCalendars)
	// HandleCommand(debugListEvents)
	// HandleCommand(debugStopWatch)
	// HandleCommand(debugUserInfo)
	// HandleCommand(disconnect)
	// HandleCommand(info)
	// HandleCommand(watchList)
	// HandleCommand(watchStart)
	// HandleCommand(watchStop)

	// // Configure modal (submit+source).
	// HandleCall("/configure-modal", RequireAdmin(
	// 	configureModal))
	// HandleCall("/f/configure-modal", RequireAdmin(
	// 	FormHandler(configureModalForm)))

	// // Lookups TODO rework when the paths are decoupled from forms.
	// HandleCall("/q/cal", RequireGoogleAuth(
	// 	LookupHandler(calendarIDLookup)))
	// HandleCall("/q/event", RequireGoogleAuth(
	// 	LookupHandler(eventLookup)))
	// HandleCall("/q/sub", RequireGoogleAuth(
	// 	LookupHandler(subscriptionIDLookup)))

	// 	func Init(router *mux.Router, m *apps.Manifest, staticAssets fs.FS, localMode bool) {
	// 		router.HandleFunc(constants.ManifestPath, fManifest(m))
	// 		router.HandleFunc(constants.InstallPath, extractCall(fInstall, localMode))
	// 		router.HandleFunc(constants.BindingsPath, extractCall(fBindings, localMode))
	// 		router.HandleFunc(constants.OAuthPath+constants.OAuthConnectPath, extractCall(fOAuthConnect, localMode))
	// 		router.HandleFunc(constants.OAuthPath+constants.OAuthCompletePath, extractCall(fOAuthComplete, localMode))

	// 		router.HandleFunc(constants.BindingPathCreate.Submit(), extractCall(fCreateTicketSubmit, localMode))
	// 		router.HandleFunc(constants.BindingPathCreate.Form(), extractCall(fCreateTicketForm, localMode))

	// 		router.HandleFunc(constants.BindingPathConfigureOAuth.Submit(), extractCall(fConfigureOAuthSubmit, localMode))
	// 		router.HandleFunc(constants.BindingPathConfigureOAuth.Form(), extractCall(fConfigureOAuthForm, localMode))

	// 		router.HandleFunc(constants.BindingPathConnect.Submit(), extractCall(fConnect, localMode))
	// 		router.HandleFunc(constants.BindingPathDisconnect.Submit(), extractCall(fDisconnect, localMode))

	// 		router.PathPrefix(constants.StaticAssetPath).Handler(http.StripPrefix("/", http.FileServer(http.FS(staticAssets))))
	// 	}

	// // Log NOT FOUND.
	// http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	Log.Warnw("not found", "path", req.URL.Path, "method", req.Method)
	// 	http.Error(w, fmt.Sprintf("Not found: %s %q", req.Method, req.URL.Path), http.StatusNotFound)
	// })
}
