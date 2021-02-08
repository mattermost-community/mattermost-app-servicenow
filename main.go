package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/routers/mattermost"
	"github.com/mattermost/mattermost-app-servicenow/routers/oauth"
)

const (
	baseURLPosition = 1
	addressPosition = 2
)

func main() {
	// Init routers
	r := mux.NewRouter()
	mattermost.Init(r)
	oauth.Init(r)

	if len(os.Args) > baseURLPosition {
		config.SetBaseURL(os.Args[baseURLPosition])
	}

	addr := ":3000"
	if len(os.Args) > addressPosition {
		addr = os.Args[addressPosition]
	}

	http.Handle("/", r)
	_ = http.ListenAndServe(addr, nil)
}
