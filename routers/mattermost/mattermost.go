package mattermost

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/server/apps"
	"github.com/pkg/errors"
)

type callHandler func(http.ResponseWriter, *http.Request, *apps.JWTClaims, *apps.Call)
type contextHandler func(http.ResponseWriter, *http.Request, *apps.JWTClaims, *apps.Context)

func Init(router *mux.Router) {
	router.HandleFunc(constants.ManifestPath, fManifest)
	router.HandleFunc(constants.InstallPath, extractCall(fInstall))
	router.HandleFunc(constants.BindingsPath, handleGetWithContext(fBindings))

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

		data, err := apps.UnmarshalCallFromReader(r.Body)
		if err != nil {
			utils.WriteBadRequestError(rw, err)
			return
		}

		f(rw, r, claims, data)
	}
}

func handleGetWithContext(f contextHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		claims, err := checkJWT(req)
		if err != nil {
			utils.WriteBadRequestError(w, err)
			return
		}

		f(w, req, claims, &apps.Context{
			TeamID:       req.Form.Get(apps.PropTeamID),
			ChannelID:    req.Form.Get(apps.PropChannelID),
			ActingUserID: req.Form.Get(apps.PropActingUserID),
			PostID:       req.Form.Get(apps.PropPostID),
		})
	}
}

func checkJWT(req *http.Request) (*apps.JWTClaims, error) {
	authValue := req.Header.Get(apps.OutgoingAuthHeader)
	if !strings.HasPrefix(authValue, "Bearer ") {
		return nil, errors.Errorf("missing %s: Bearer header", apps.OutgoingAuthHeader)
	}

	jwtoken := strings.TrimPrefix(authValue, "Bearer ")
	claims := apps.JWTClaims{}
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
