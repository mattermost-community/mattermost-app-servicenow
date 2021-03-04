package mattermost

import (
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/server/api"
	"github.com/pkg/errors"
)

var ErrUnexpectedSignMethod = errors.New("unexpected signing method")
var ErrMissingHeader = errors.Errorf("missing %s: Bearer header", api.OutgoingAuthHeader)

type callHandler func(http.ResponseWriter, *http.Request, *apps.Call)

func Init(router *mux.Router, m *apps.Manifest, staticAssets fs.FS, localMode bool) {
	router.HandleFunc(constants.ManifestPath, fManifest(m))
	router.HandleFunc(constants.InstallPath, extractCall(fInstall, localMode))
	router.HandleFunc(constants.BindingsPath, extractCall(fBindings, localMode))

	router.HandleFunc(constants.BindingPathCreate, extractCall(fCreateTicket, localMode))
	router.HandleFunc(constants.BindingPathConnect, extractCall(fConnect, localMode))
	router.HandleFunc(constants.BindingPathDisconnect, extractCall(fDisconnect, localMode))
	router.HandleFunc(constants.BindingPathConfigureOAuth, extractCall(fConfigureOAuth, localMode))

	router.PathPrefix(constants.StaticAssetPath).Handler(http.StripPrefix("/", http.FileServer(http.FS(staticAssets))))
}

func extractCall(f callHandler, localMode bool) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := apps.UnmarshalCallFromReader(r.Body)
		if err != nil {
			utils.WriteBadRequestError(rw, err)
			return
		}

		if localMode {
			claims, err := checkJWT(r)
			if err != nil {
				utils.WriteBadRequestError(rw, err)
				return
			}

			if data.Context.ActingUserID != "" && data.Context.ActingUserID != claims.ActingUserID {
				utils.WriteBadRequestError(rw, errors.New("JWT claim doesn't match actingUserID in context"))
				return
			}
		}

		f(rw, r, data)
	}
}

func checkJWT(req *http.Request) (*api.JWTClaims, error) {
	authValue := req.Header.Get(api.OutgoingAuthHeader)
	if !strings.HasPrefix(authValue, "Bearer ") {
		return nil, ErrMissingHeader
	}

	jwtoken := strings.TrimPrefix(authValue, "Bearer ")
	claims := api.JWTClaims{}
	_, err := jwt.ParseWithClaims(jwtoken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%w: %v", ErrUnexpectedSignMethod, token.Header["alg"])
		}
		return []byte(constants.AppSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return &claims, nil
}
