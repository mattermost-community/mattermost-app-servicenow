package oauth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/app"
	"github.com/mattermost/mattermost-app-servicenow/store"
	"github.com/mattermost/mattermost-app-servicenow/utils"
)

func oauth2Complete(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	state := r.URL.Query().Get("state")
	if !store.VerifyState(state) {
		http.Error(w, "State not found", http.StatusBadRequest)
		return
	}

	userID, _, err := utils.ParseOAuthState(state)
	if err != nil {
		http.Error(w, "State badly formed", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	token, err := app.GetOAuthConfig().Exchange(ctx, code)
	if err != nil {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	store.StoreToken(token, userID)

	connectedString := "You have successfuly connected the ServiceNow Mattermost App to ServiceNow. Please close this window."
	html := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
			<head>
				<script>
					window.close();
				</script>
			</head>
			<body>
				<p>%s</p>
			</body>
		</html>
		`, connectedString)

	w.Header().Set("Content-Type", "text/html")
	_, _ = w.Write([]byte(html))
}
