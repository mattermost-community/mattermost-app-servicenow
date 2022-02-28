package function

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"

	"github.com/mattermost/mattermost-app-servicenow/goapp"
)

type Client struct {
	client      *http.Client
	tokenSource oauth2.TokenSource
	original    oauth2.Token
}

type BaseTicket struct {
	ID         string       `json:"sys_id"`
	ClassName  string       `json:"sys_class_name"`
	Tags       string       `json:"sys_tags"`
	Domain     TicketDomain `json:"sys_domain"`
	DomainPath string       `json:"sys_domain_path"`
	ModCount   string       `json:"sys_mod_count"`
	UpdatedBy  string       `json:"sys_updated_by"`
	CreatedBy  string       `json:"sys_created_by"`
	// UpdatedOn  time.Time    `json:"sys_updated_on"`
	// CreatedOn  time.Time    `json:"sys_created_on"`
}

type TicketDomain struct {
	Link  string `json:"link"`
	Value string `json:"value"`
}

type CreateTicketResponse struct {
	Result BaseTicket `json:"result"`
}

var ErrUnexpectedStatus = errors.New("returned with unexpected status")

func makeClient(creq goapp.CallRequest) (*Client, error) {
	oAuthConf, err := oauth2Config(creq)
	if err != nil {
		return nil, err
	}

	if creq.OAuth2User() == nil || creq.OAuth2User().Token == nil {
		return nil, errors.New("user account is not connected")
	}

	token := creq.OAuth2User().Token
	tokSrc := oAuthConf.TokenSource(creq.GoContext, token)
	return &Client{
		client:      oauth2.NewClient(creq.GoContext, tokSrc),
		tokenSource: tokSrc,
		original:    *token,
	}, nil
}

func (a *App) CreateIncident(c *Client, creq goapp.CallRequest, table string, v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	resp, err := c.client.Post(apiURL(creq, "/api/now/table/"+table), "application/json", bytes.NewReader(b))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	tok, _ := c.tokenSource.Token()
	if tok.AccessToken != c.original.AccessToken {
		if user := creq.OAuth2User(); user != nil {
			user.Token = tok
			err = a.StoreConnectedUser(creq, user)
			if err != nil {
				return "", err
			}
		}
	}
	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("%w: %v", ErrUnexpectedStatus, resp.Status)
	}

	var ticketResp CreateTicketResponse
	err = json.NewDecoder(resp.Body).Decode(&ticketResp)
	if err != nil {
		return "", errors.Wrap(err, "could not decode create ticket response")
	}

	return ticketResp.Result.ID, nil
}

func apiURL(creq goapp.CallRequest, apiPath string) string {
	return creq.Context.OAuth2.RemoteURL + apiPath
}
