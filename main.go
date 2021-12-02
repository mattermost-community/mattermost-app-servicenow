package main

import (
	"embed"
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-plugin-apps/apps"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"

	"github.com/mattermost/mattermost-app-servicenow/routers/mattermost"
)

const (
	baseURLPosition = 1
	addressPosition = 2
)

//go:embed manifest.json
var manifestSource []byte //nolint: gochecknoglobals

//go:embed static
var staticAssets embed.FS //nolint: gochecknoglobals

func main() {
	var manifest apps.Manifest

	err := json.Unmarshal(manifestSource, &manifest)
	if err != nil {
		panic("failed to load manfest: " + err.Error())
	}

	localMode := os.Getenv("LOCAL") == "true"

	var verbose bool

	flag.BoolVarP(&verbose, "verbose", "v", false, "help message for flagname")

	flag.Parse()

	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	// Init routers
	r := mux.NewRouter()
	r.Use(logRequest)

	mattermost.Init(r, &manifest, staticAssets, localMode)

	http.Handle("/", r)

	if localMode {
		baseURL := os.Args[baseURLPosition]

		addr := ":3000"
		if len(os.Args) > addressPosition {
			addr = os.Args[addressPosition]
		}

		manifest.HTTP.RootURL = baseURL

		_ = http.ListenAndServe(addr, nil)

		return
	}

	lambda.Start(httpadapter.New(r).Proxy)
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithField("method", r.Method).WithField("url", r.URL.Path).Debug("Received HTTP request")

		next.ServeHTTP(w, r)
	})
}
