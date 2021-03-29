package constants

import (
	"time"

	"github.com/mattermost/mattermost-app-servicenow/utils"
)

const (
	ManifestPath    = "/manifest"
	InstallPath     = "/install"
	BindingsPath    = "/bindings"
	StaticAssetPath = "/static"

	OAuthPath         = "/oauth"
	OAuthConnectPath  = "/connect"
	OAuthCompletePath = "/complete"

	BindingPathCreate         utils.Path = "/create"
	BindingPathConnect        utils.Path = "/connect"
	BindingPathDisconnect     utils.Path = "/disconnect"
	BindingPathConfigureOAuth utils.Path = "/configure/oauth"

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
)
