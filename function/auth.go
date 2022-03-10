package function

import (
	"github.com/pkg/errors"
	"golang.org/x/oauth2"

	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-servicenow/goapp"
)

func oauth2Config(creq goapp.CallRequest) (*oauth2.Config, error) {
	cc := creq.Context
	if cc.OAuth2.ClientID == "" || cc.OAuth2.RemoteURL == "" {
		return nil, errors.New("oauth2 is not configured. Please have a system administrator use `/servicenow configure` command")
	}
	return &oauth2.Config{
		ClientID:     cc.OAuth2.ClientID,
		ClientSecret: cc.OAuth2.ClientSecret,
		RedirectURL:  cc.OAuth2.CompleteURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  cc.OAuth2.RemoteURL + "/oauth_auth.do",
			TokenURL: cc.OAuth2.RemoteURL + "/oauth_token.do",
		},
	}, nil
}

func (a *App) oauth2Connect(creq goapp.CallRequest) apps.CallResponse {
	state := creq.GetValue(fState, "")
	c, err := oauth2Config(creq)
	if err != nil {
		return apps.NewErrorResponse(err)
	}

	return apps.NewDataResponse(c.AuthCodeURL(state, oauth2.AccessTypeOffline))
}

func (a *App) oauth2Complete(creq goapp.CallRequest) apps.CallResponse {
	code := creq.GetValue(fCode, "")
	oauth2Config, err := oauth2Config(creq)
	if err != nil {
		return apps.NewErrorResponse(err)
	}

	token, err := oauth2Config.Exchange(creq.GoContext, code)
	if err != nil {
		return apps.NewErrorResponse(errors.Wrap(err, "failed token exchange"))
	}

	user := goapp.User{
		Token:        token,
		MattermostID: creq.Context.ActingUserID,
	}
	if err = a.StoreConnectedUser(creq, &user); err != nil {
		return apps.NewErrorResponse(errors.Wrap(err, "failed to store OAuth user info to Mattermost"))
	}

	return apps.NewTextResponse("completed connecting to ServiceNow with OAuth2.")
}
