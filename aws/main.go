package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/gorilla/mux"
	"go.uber.org/zap/zapcore"

	"github.com/mattermost/mattermost-plugin-apps/apps"
	"github.com/mattermost/mattermost-plugin-apps/utils"

	function "github.com/mattermost/mattermost-app-servicenow/function"
)

func main() {
	r := mux.NewRouter()
	log := utils.MustMakeCommandLogger(zapcore.DebugLevel)
	function.Init(string(apps.DeployAWSLambda), r, log)
	http.Handle("/", r)

	lambda.Start(httpadapter.New(http.DefaultServeMux).Proxy)
}
