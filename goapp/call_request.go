package goapp

import (
	"context"
	"path"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/appclient"
	"github.com/mattermost/mattermost-plugin-apps/utils"
)

type CallRequest struct {
	apps.CallRequest

	App          App
	GoContext    context.Context
	asBot        *appclient.Client
	asActingUser *appclient.Client
	user         *User
}

type HandlerFunc func(CallRequest) apps.CallResponse

func (creq CallRequest) AsBot() *appclient.Client {
	if creq.asBot != nil {
		return creq.asBot
	}
	creq.asBot = appclient.AsBot(creq.Context)
	return creq.asBot
}

func (creq CallRequest) AsActingUser() *appclient.Client {
	if creq.asActingUser != nil {
		return creq.asActingUser
	}
	creq.asActingUser = appclient.AsActingUser(creq.Context)
	return creq.asActingUser
}

func FormHandler(h func(CallRequest) (apps.Form, error)) HandlerFunc {
	return func(creq CallRequest) apps.CallResponse {
		f, err := h(creq)
		if err != nil {
			return apps.NewErrorResponse(err)
		}
		return apps.NewFormResponse(f)
	}
}

func LookupHandler(h func(CallRequest) []apps.SelectOption) HandlerFunc {
	return func(creq CallRequest) apps.CallResponse {
		opts := h(creq)
		return apps.NewLookupResponse(opts)
	}
}

func CallHandler(h func(CallRequest) (string, error)) HandlerFunc {
	return func(creq CallRequest) apps.CallResponse {
		text, err := h(creq)
		if err != nil {
			return apps.NewErrorResponse(err)
		}
		return apps.NewTextResponse(text)
	}
}

func RequireAdmin(h HandlerFunc) HandlerFunc {
	return func(creq CallRequest) apps.CallResponse {
		if creq.Context.ActingUser != nil && !creq.Context.ActingUser.IsSystemAdmin() {
			return apps.NewErrorResponse(
				utils.NewUnauthorizedError("system administrator role is required to invoke " + creq.Path))
		}
		return h(creq)
	}
}

func RequireConnectedUser(h HandlerFunc) HandlerFunc {
	return func(creq CallRequest) apps.CallResponse {
		if creq.Context.OAuth2.User == nil {
			return apps.NewErrorResponse(
				utils.NewUnauthorizedError("missing user record, required for " + creq.Path + ". Please use `/apps connect` to connect your Google account."))
		}
		user := User{}
		utils.Remarshal(&user, creq.Context.OAuth2.User)
		creq.user = &user
		return h(creq)
	}
}

func (creq CallRequest) appProxyURL(paths ...string) string {
	p := path.Join(append([]string{creq.Context.AppPath}, paths...)...)
	return creq.Context.MattermostSiteURL + p
}
