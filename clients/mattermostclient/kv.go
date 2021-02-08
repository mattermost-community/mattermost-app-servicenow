package mattermostclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type KVClient struct {
	client *http.Client
	token  string
	url    string
}

func NewKVClient(botAccessToken string, baseURL string) *KVClient {
	return &KVClient{
		client: http.DefaultClient,
		token:  botAccessToken,
		url:    baseURL,
	}
}

func (c *KVClient) KVSet(key string, v interface{}) error {
	body, err := json.Marshal(v)
	if err != nil {
		return err
	}

	req, err := c.newRequest(key, bytes.NewReader(body), http.MethodPost)
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%w, %s", ErrUnexpectedStatus, res.Status)
	}

	return nil
}

func (c *KVClient) KVGet(key string, out interface{}) error {
	req, err := c.newRequest(key, nil, http.MethodGet)
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%w, %s", ErrUnexpectedStatus, res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(&out)
	if err != nil {
		return err
	}

	return nil
}

func (c *KVClient) KVDelete(key string) error {
	req, err := c.newRequest(key, nil, http.MethodDelete)
	if err != nil {
		return err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%w, %s", ErrUnexpectedStatus, res.Status)
	}

	return nil
}

func (c *KVClient) newRequest(key string, body io.Reader, method string) (*http.Request, error) {
	req, err := http.NewRequest(method, c.getKVURL(key), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "BEARER"+" "+c.token)

	req.ParseForm()
	req.Form.Add("prefix", "servicenow")

	return req, nil
}

func (c *KVClient) getKVURL(key string) string {
	return c.url + "/plugins/com.mattermost.apps/api/v1/kv/" + url.PathEscape(key)
}
