package servicenowclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-app-servicenow/app"
	"github.com/mattermost/mattermost-app-servicenow/config"

	"github.com/mattermost/mattermost-plugin-apps/apps"
)

type Client struct {
	client *http.Client
}

var ErrUnexpectedStatus = errors.New("returned with unexpected status")

func NewClient(cc *apps.Context) *Client {
	ctx := context.Background()
	oAuthConf := app.GetOAuthConfig(cc)

	token := app.GetTokenFromContext(cc)
	if token == nil {
		return nil
	}

	return &Client{
		client: oAuthConf.Client(ctx, token),
	}
}

func (c *Client) CreateIncident(table string, v interface{}, cc *apps.Context) (string, error) {
	url := fmt.Sprintf("%s%s/%s", config.ServiceNowInstance(cc), "/api/now/table", table)

	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	resp, err := c.client.Post(url, "application/json", bytes.NewReader(b))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("%w: %v", ErrUnexpectedStatus, resp.Status)
	}

	var ticket CreateTicketResponse

	err = json.NewDecoder(resp.Body).Decode(&ticket)
	if err != nil {
		return "", errors.Wrap(err, "could not decode create ticket response")
	}

	return ticket.Result.ID, nil
}
