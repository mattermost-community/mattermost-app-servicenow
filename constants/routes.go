package constants

import "time"

const (
	ManifestPath = "/manifest"
	InstallPath  = "/install"
	BindingsPath = "/bindings"

	OAuthPath         = "/oauth"
	OAuthConnectPath  = "/connect"
	OAuthCompletePath = "/complete"

	BindingPathCreate         = "/create"
	BindingPathConnect        = "/connect"
	BindingPathDisconnect     = "/disconnect"
	BindingPathConfigureOAuth = "/configure/oauth"

	LocationCreate         = "create"
	LocationConnect        = "connect"
	LocationDisconnect     = "disconnect"
	LocationConfigure      = "configure"
	LocationConfigureOAuth = "oauth"

	ConfigFile = "config.json"
	TokenFile  = "tokens.json"

	CommandTrigger = "com.mattermost.servicenow"

	AppSecret = "1234"

	OAuthStateTTL      = 5 * time.Minute
	OAuthStateGCTicker = 30 * time.Second

	MattermostURL = "http://localhost:8065"

	TableIDGetField = "table"
)
