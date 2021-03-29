package utils

import (
	"strings"

	"github.com/mattermost/mattermost-plugin-apps/apps"
)

func GetStringFromMapInterface(in map[string]interface{}, key, def string) string {
	if len(in) == 0 {
		return def
	}

	v, ok := in[key]
	if !ok {
		return def
	}

	out, ok := v.(string)
	if !ok {
		return def
	}

	return out
}

func GetIconURL(siteURL, name string, appID apps.AppID) string {
	return strings.TrimRight(siteURL, "/") + "/plugins/com.mattermost.apps/api/v1/static/" + string(appID) + "/" + name
}
