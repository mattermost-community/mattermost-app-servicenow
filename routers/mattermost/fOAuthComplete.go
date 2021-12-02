package mattermost

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/apps/appclient"

	"github.com/mattermost/mattermost-app-servicenow/app"
	"github.com/mattermost/mattermost-app-servicenow/utils"
)

func fOAuthComplete(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	code, ok := c.Values["code"].(string)
	if !ok {
		utils.WriteCallErrorResponse(w, "No code provided")
		return
	}

	ctx := context.Background()

	token, err := app.GetOAuthConfig(c.Context).Exchange(ctx, code)
	if err != nil {
		utils.WriteCallErrorResponse(w, "Cannot exchage code for token.")
		return
	}

	client := appclient.AsActingUser(c.Context)

	err = client.StoreOAuth2User(token)
	if err != nil {
		errMsg := fmt.Sprintf("Error storing the user: %v", err)
		utils.WriteCallErrorResponse(w, errMsg)
		log.Printf(errMsg, err)

		return
	}

	utils.WriteCallStandardResponse(w, "You are now connected to ServiceNow")
}
