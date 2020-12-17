package servicenowclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/app"
	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-app-servicenow/store"
)

type Client struct {
	client *http.Client
}

func NewClient(userID string) *Client {
	ctx := context.Background()
	conf := app.GetOAuthConfig()
	token, found := store.GetToken(userID)
	if !found {
		return nil
	}
	return &Client{
		client: conf.Client(ctx, token),
	}
}

func (c *Client) CreateIncident(table string, v interface{}) error {
	url := fmt.Sprintf("%s/%s/%s", config.ServiceNowInstance(), "/api/now/table/", table)
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	resp, err := c.client.Post(url, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return fmt.Errorf("call returned with status %v", resp.Status)
	}

	return nil
}
