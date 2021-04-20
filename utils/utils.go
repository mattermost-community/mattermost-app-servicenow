package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/mattermost/mattermost-plugin-apps/apps"
)

const AppsPluginID = "com.mattermost.apps"

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

func GetIconURL(name string, cc *apps.Context) string {
	return strings.TrimRight(cc.MattermostSiteURL, "/") + cc.AppPath + "/static/" + name
}

func GetAppsPluginURL(siteURL string) string {
	return fmt.Sprintf("%s/plugins/%s", siteURL, AppsPluginID)
}

func DumpObject(c interface{}) {
	b, _ := json.MarshalIndent(c, "", "    ")
	log.Printf("%s\n", string(b))
}
