package mattermostclient

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"
)

var ErrUnexpectedStatus = errors.New("returned with unexpected status")

type MMClient struct {
	client *http.Client
	botID  string
	token  string
	url    string
}

func NewMMClient(botID, botAccessToken, baseURL string) *MMClient {
	return &MMClient{
		client: http.DefaultClient,
		botID:  botID,
		token:  botAccessToken,
		url:    baseURL,
	}
}

func (c *MMClient) asBot(f func(mmclient *model.Client4, botUserID string) error) error {
	mmClient := model.NewAPIv4Client(c.url)
	mmClient.SetToken(c.token)

	return f(mmClient, c.botID)
}

// func (c *MMClient) dm(userID string, post *model.Post) (*model.Post, error) {
// 	var createdPost *model.Post
// 	err := c.asBot(func(mmclient *model.Client4, botUserID string) error {
// 		var res *model.Response
// 		post.UserId = botUserID

// 		createdPost, res = mmclient.CreatePost(post)
// 		if res.StatusCode != http.StatusCreated {
// 			if res.Error != nil {
// 				return res.Error
// 			}
// 			return fmt.Errorf("returned with status %d", res.StatusCode)
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return createdPost, nil
// }

func (c *MMClient) GetUser(userID string) (*model.User, error) {
	var user *model.User

	err := c.asBot(func(mmclient *model.Client4, botUserID string) error {
		var res *model.Response
		user, res = mmclient.GetUser(userID, "")
		if res.StatusCode != http.StatusOK {
			if res.Error != nil {
				return res.Error
			}
			return fmt.Errorf("%w: %d", ErrUnexpectedStatus, res.StatusCode)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}
