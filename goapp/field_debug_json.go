package goapp

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/utils"
)

var fieldDebugJSON = apps.Field{
	Type:        apps.FieldTypeBool,
	Name:        "json",
	Description: "Output JSON",
}

func (creq CallRequest) AppendDebugJSON(in []apps.Field) []apps.Field {
	if !creq.Context.DeveloperMode {
		return in
	}

	return append(in, fieldDebugJSON)
}

func (creq CallRequest) Respond(message string, v interface{}) apps.CallResponse {
	outJSON := creq.BoolValue("json")
	if outJSON {
		message += "----\n"
		message += utils.JSONBlock(v)
	}
	return apps.NewTextResponse(message)
}
