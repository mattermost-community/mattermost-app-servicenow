package mattermost

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-plugin-apps/apps"

	"github.com/mattermost/mattermost-app-servicenow/constants"
	"github.com/mattermost/mattermost-app-servicenow/utils"
)

func fInstall(w http.ResponseWriter, r *http.Request, c *apps.CallRequest) {
	utils.WriteCallStandardResponse(w, fmt.Sprintf("Service now installed! "+
		"Please run `/%s configure oauth` to configure the link between Mattermost and your Service Now instance.",
		constants.CommandTrigger))
}
