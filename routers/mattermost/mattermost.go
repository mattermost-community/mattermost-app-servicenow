package mattermost

import (
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
)

var ErrUnexpectedSignMethod = errors.New("unexpected signing method")
var ErrMissingHeader = errors.Errorf("missing %s: Bearer header", apps.OutgoingAuthHeader)
var ErrActingUserMismatch = errors.New("JWT claim doesn't match actingUser.Id in context")

type callHandler func(http.ResponseWriter, *http.Request, *apps.CallRequest)

func Init(router *mux.Router, m *apps.Manifest, staticAssets fs.FS, localMode bool) {
	router.HandleFunc(constants.ManifestPath, fManifest(m))
	router.HandleFunc(constants.InstallPath, extractCall(fInstall, localMode))
	router.HandleFunc(constants.BindingsPath, extractCall(fBindings, localMode))
	router.HandleFunc(constants.OAuthPath+constants.OAuthConnectPath, extractCall(fOAuthConnect, localMode))
	router.HandleFunc(constants.OAuthPath+constants.OAuthCompletePath, extractCall(fOAuthComplete, localMode))

	router.HandleFunc(constants.BindingPathCreate.Submit(), extractCall(fCreateTicketSubmit, localMode))
	router.HandleFunc(constants.BindingPathCreate.Form(), extractCall(fCreateTicketForm, localMode))

	router.HandleFunc(constants.BindingPathConfigureOAuth.Submit(), extractCall(fConfigureOAuthSubmit, localMode))
	router.HandleFunc(constants.BindingPathConfigureOAuth.Form(), extractCall(fConfigureOAuthForm, localMode))

	router.HandleFunc(constants.BindingPathConnect.Submit(), extractCall(fConnect, localMode))
	router.HandleFunc(constants.BindingPathDisconnect.Submit(), extractCall(fDisconnect, localMode))

	router.PathPrefix(constants.StaticAssetPath).Handler(http.StripPrefix("/", http.FileServer(http.FS(staticAssets))))
}

func extractCall(f callHandler, localMode bool) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := apps.CallRequestFromJSONReader(r.Body)
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

			if data.Context.ActingUser.Id != "" && data.Context.ActingUser.Id != claims.ActingUserID {
				utils.WriteBadRequestError(rw, ErrActingUserMismatch)
				return
			}
		}

		f(rw, r, data)
	}
}

func checkJWT(req *http.Request) (*apps.JWTClaims, error) {
	authValue := req.Header.Get(apps.OutgoingAuthHeader)
	if !strings.HasPrefix(authValue, "Bearer ") {
		return nil, ErrMissingHeader
	}

	jwtoken := strings.TrimPrefix(authValue, "Bearer ")
	claims := apps.JWTClaims{}
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
