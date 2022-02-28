package goapp

import (
	"time"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/appclient"
	"github.com/mattermost/mattermost-plugin-apps/utils"
)

type User struct {
	MattermostID string
	RemoteID     string
	Token        *oauth2.Token
}

func (a *App) RemoveConnectedUser(creq CallRequest) error {
	asActingUser := appclient.AsActingUser(creq.Context)
	err := asActingUser.StoreOAuth2User(creq.Context.AppID, nil)
	if err != nil {
		return apps.NewErrorResponse(errors.Wrap(err, "failed to removed the user record"))
	}

	creq.App.Logger.Debugw("Removed user record", "id", creq.Context.ActingUserID)
	return nil
}

func (a *App) StoreConnectedUser(creq CallRequest, user *User) error {
	if user == nil {
		return a.RemoveConnectedUser(creq)
	}

	asActingUser := appclient.AsActingUser(creq.Context)
	user.MattermostID = creq.Context.ActingUserID
	err := asActingUser.StoreOAuth2User(creq.Context.AppID, user)
	if err != nil {
		return apps.NewErrorResponse(errors.Wrap(err, "failed to store the user record"))
	}

	accessTokenLog := ""
	expires := ""
	refreshTokenLog := ""
	if user.Token != nil {
		accessTokenLog = utils.LastN(user.Token.AccessToken, 4)
		expires = user.Token.Expiry.Format(time.RFC822)
		refreshTokenLog = utils.LastN(user.Token.RefreshToken, 4)
	}
	creq.App.Logger.Debugw("Updated user record", "id", user.MattermostID, "access_token", accessTokenLog, "expires", expires, "refresh_token", refreshTokenLog)
	return nil
}
