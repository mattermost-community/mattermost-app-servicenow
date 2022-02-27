package function

import (
	"github.com/gorilla/mux"

	"github.com/mattermost/mattermost-plugin-apps/utils"

	root "github.com/mattermost/mattermost-app-servicenow"
	"github.com/mattermost/mattermost-app-servicenow/goapp"
)

var BuildHash string
var BuildHashShort string
var BuildDate string

// Field names.
const (
	fURL          = "url"
	fClientID     = "client_id"
	fClientSecret = "client_secret"
	fState        = "state"
	fCode         = "code"
	fTable        = "table"
)

type App struct {
	goapp.App
	mode string
}

func Init(mode string, r *mux.Router, log utils.Logger) {
	app := App{
		mode: mode,
		App: *goapp.NewApp(r, log).
			WithManifest(root.AppManifest).
			WithStatic(root.Static).
			WithIcon(root.AppManifest.Icon),
	}

	app.HandleCall("/install", install)

	// Bindings.
	app.HandleCall("/bindings", app.getBindings)

	// OAuth2 callbacks.
	app.HandleCall("/oauth2/connect", oauth2Connect)
	app.HandleCall("/oauth2/complete", oauth2Complete)

	// 	router.HandleFunc(constants.BindingPathCreate.Submit(), extractCall(fCreateTicketSubmit, localMode))
	// 	router.HandleFunc(constants.BindingPathCreate.Form(), extractCall(fCreateTicketForm, localMode))

	// Command submit handlers.
	app.HandleCommand(app.infoCommand())
	app.HandleCommand(configureCommand)
	app.HandleCommand(goapp.ConnectCommand("ServiceNow"))
	app.HandleCommand(goapp.DisconnectCommand("ServiceNow"))

	// Create ticket handlers.
	app.HandleCall("/create-ticket", goapp.RequireConnectedUser(goapp.CallHandler(createTicketHandler)))
	app.HandleCall("/form/create-ticket", goapp.FormHandler(app.createTicketFormHandler))
}
