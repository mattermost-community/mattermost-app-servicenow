package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/routers/mattermost"
	"github.com/mattermost/mattermost-app-servicenow/routers/oauth"
	"github.com/mattermost/mattermost-app-servicenow/store"
	"github.com/mattermost/mattermost-plugin-apps/server/apps"
)

func main() {
	r := mux.NewRouter()
	mattermost.Init(r)
	oauth.Init(r)

	store.LoadTokens()
	config.Load()

	if len(os.Args) > 1 {
		config.SetBaseURL(os.Args[1])
	}

	config.AddTable(config.TableConfig{
		ID:          "incidents",
		DisplayName: "Incidents",
		Fields: []*apps.Field{
			{
				Name:       "short_description",
				ModalLabel: "Short Description",
				Type:       apps.FieldTypeText,
			},
			{
				Name:       "description",
				ModalLabel: "Long Description",
				Type:       apps.FieldTypeText,
			},
		},
		Ticketable:  true,
		PostDefault: "short_description",
	})

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
