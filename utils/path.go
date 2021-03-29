package utils

import "github.com/mattermost/mattermost-plugin-apps/apps"

type Path string

func (p Path) Submit() string {
	return string(p) + "/" + string(apps.CallTypeSubmit)
}

func (p Path) Form() string {
	return string(p) + "/" + string(apps.CallTypeForm)
}
