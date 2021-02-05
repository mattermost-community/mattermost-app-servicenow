package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/routers/mattermost"
	"github.com/mattermost/mattermost-app-servicenow/routers/oauth"
	"github.com/mattermost/mattermost-app-servicenow/store"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
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

	addr := ":3000"
	if len(os.Args) > 2 {
		addr = os.Args[2]
	}

	config.AddTable(config.TableConfig{
		ID:          "incidents",
		DisplayName: "Incidents",
		Fields: []*api.Field{
			{
				Name:       "short_description",
				ModalLabel: "Short Description",
				Type:       api.FieldTypeText,
			},
			{
				Name:       "description",
				ModalLabel: "Long Description",
				Type:       api.FieldTypeText,
			},
		},
		Post:        true,
		Command:     true,
		Header:      true,
		PostDefault: "short_description",
	})

	http.Handle("/", r)
	_ = http.ListenAndServe(addr, nil)
}
