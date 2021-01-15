package mattermost

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
	"github.com/pkg/errors"
)

type callHandler func(http.ResponseWriter, *http.Request, *api.JWTClaims, *api.Call)

func Init(router *mux.Router) {
	router.HandleFunc(constants.ManifestPath, fManifest)
	router.HandleFunc(constants.InstallPath, extractCall(fInstall))
	router.HandleFunc(constants.BindingsPath, extractCall(fBindings))

	router.HandleFunc(constants.BindingPathCreate, extractCall(fCreateTicket))
	router.HandleFunc(constants.BindingPathConnect, extractCall(fConnect))
	router.HandleFunc(constants.BindingPathDisconnect, extractCall(fDisconnect))
	router.HandleFunc(constants.BindingPathConfigureOAuth, extractCall(fConfigureOAuth))
}

func extractCall(f callHandler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		claims, err := checkJWT(r)
		if err != nil {
			utils.WriteBadRequestError(rw, err)
			return
		}

		data, err := api.UnmarshalCallFromReader(r.Body)
		if err != nil {
			utils.WriteBadRequestError(rw, err)
			return
		}

		f(rw, r, claims, data)
	}
}

func checkJWT(req *http.Request) (*api.JWTClaims, error) {
	authValue := req.Header.Get(api.OutgoingAuthHeader)
	if !strings.HasPrefix(authValue, "Bearer ") {
		return nil, errors.Errorf("missing %s: Bearer header", api.OutgoingAuthHeader)
	}

	jwtoken := strings.TrimPrefix(authValue, "Bearer ")
	claims := api.JWTClaims{}
	_, err := jwt.ParseWithClaims(jwtoken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(constants.AppSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return &claims, nil
}
