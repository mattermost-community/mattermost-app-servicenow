package oauth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-app-servicenow/constants"
)

func Init(router *mux.Router) {
	router.HandleFunc(constants.OAuthPath+constants.OAuthCompletePath, oauth2Complete).Methods(http.MethodGet)
}
