package function

import (
	"github.com/pkg/errors"
	"golang.org/x/oauth2"

	"github.com/mattermost/mattermost-app-servicenow/goapp"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/appclient"
)

func oauth2Config(creq goapp.CallRequest) *oauth2.Config {
	cc := creq.Context
	if cc.OAuth2.ClientID == "" || cc.OAuth2.Data == nil {
		return nil
	}

	appconf := appConfig(creq)
	return &oauth2.Config{
		ClientID:     cc.OAuth2.ClientID,
		ClientSecret: cc.OAuth2.ClientSecret,
		RedirectURL:  cc.OAuth2.CompleteURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  appconf.RemoteURL + "/oauth_auth.do",
			TokenURL: appconf.RemoteURL + "/oauth_token.do",
		},
	}
}

func oauth2Connect(creq goapp.CallRequest) apps.CallResponse {
	state := creq.GetValue(fState, "")
	url := oauth2Config(creq).AuthCodeURL(state, oauth2.AccessTypeOffline)
	return apps.NewDataResponse(url)
}

func oauth2Complete(creq goapp.CallRequest) apps.CallResponse {
	code := creq.GetValue(fCode, "")
	oauth2Config := oauth2Config(creq)

	token, err := oauth2Config.Exchange(creq.GoContext, code)
	if err != nil {
		return apps.NewErrorResponse(errors.Wrap(err, "failed token exchange"))
	}

	asActingUser := appclient.AsActingUser(creq.Context)
	user := goapp.User{
		Token:        token,
		MattermostID: creq.Context.ActingUserID,
	}
	if err = asActingUser.StoreOAuth2User(creq.Context.AppID, user); err != nil {
		return apps.NewErrorResponse(errors.Wrap(err, "failed to store OAuth user info to Mattermost"))
	}

	return apps.NewTextResponse("completed connecting to ServiceNow with OAuth2.")
}
