package mattermostclient

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-app-servicenow/config"
	"github.com/mattermost/mattermost-server/v5/model"
)

func asBot(f func(mmclient *model.Client4, botUserID string) error) error {
	conf := config.Mattermost()
	mmClient := model.NewAPIv4Client(conf.MattermostURL)
	mmClient.SetToken(conf.BotAccessToken)

	return f(mmClient, conf.BotID)
}

func dm(userID string, post *model.Post) (*model.Post, error) {
	var createdPost *model.Post
	err := asBot(func(mmclient *model.Client4, botUserID string) error {
		var res *model.Response
		post.UserId = botUserID

		createdPost, res = mmclient.CreatePost(post)
		if res.StatusCode != http.StatusCreated {
			if res.Error != nil {
				return res.Error
			}
			return fmt.Errorf("returned with status %d", res.StatusCode)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return createdPost, nil
}

func GetUser(userID string) (*model.User, error) {
	var user *model.User
	err := asBot(func(mmclient *model.Client4, botUserID string) error {
		var res *model.Response
		user, res = mmclient.GetUser(userID, "")
		if res.StatusCode != http.StatusOK {
			if res.Error != nil {
				return res.Error
			}
			return fmt.Errorf("returned with status %d", res.StatusCode)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}
