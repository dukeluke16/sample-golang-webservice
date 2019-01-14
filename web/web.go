package web

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/newrelic/go-agent"
	"github.com/dukeluke16/sample-golang-webservice/config"
	"github.com/dukeluke16/sample-golang-webservice/logger"
)

// EnableNewRelic enivronment variable key
var EnableNewRelic string

// listenAndServe for mocking out the http.ListenAndServe
var listenAndServe = http.ListenAndServe

// wrapHandleFunc for managing NewRelic bypass
var wrapHandleFunc = newrelic.WrapHandleFunc

var certDirectory = "./certs/"

// bypassNewRelic for environments not configured for NewRelic
func bypassNewRelic(app newrelic.Application, pattern string, handler func(http.ResponseWriter, *http.Request)) (string, func(http.ResponseWriter, *http.Request)) {
	return pattern, func(w http.ResponseWriter, r *http.Request) { http.HandlerFunc(handler).ServeHTTP(w, r) }
}

func configureNewRelic() newrelic.Application {
	if EnableNewRelic != "true" {
		wrapHandleFunc = bypassNewRelic
		logger.Warning.Println("New Relic is bypassed!")
	}

	datacenter := strings.ToUpper(os.Getenv("NEW_RELIC_DATACENTER"))
	environment := strings.ToUpper(os.Getenv("NEW_RELIC_ENVIRONMENT"))
	roletypeid := strings.ToUpper(os.Getenv("NEW_RELIC_ROLETYPEID"))
	appname := fmt.Sprintf("%v-%v1-%v", datacenter, environment, roletypeid)

	newRelicConfig := newrelic.NewConfig(appname, os.Getenv("NEW_RELIC_LICENSE_KEY"))

	// Configure Certificate Trust Chain for New Relic
	certFile, _ := os.Open(filepath.Join(certDirectory, "newrelic.pem"))
	certData, _ := ioutil.ReadAll(certFile)
	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM(certData)
	tlsConfig := &tls.Config{
		RootCAs: roots,
	}

	// Configure Transport
	newRelicProxy := config.ProxyURI("NEW_RELIC_PROXY")
	logger.Info.Println("Configuring New Relic for Certificate Trust Chain & Proxy.")
	newRelicConfig.Transport = &http.Transport{
		Proxy:           http.ProxyURL(newRelicProxy),
		TLSClientConfig: tlsConfig,
	}

	newRelicConfig.Labels["Application"] = os.Getenv("NEW_RELIC_APP_NAME")
	newRelicConfig.Labels["Datacenter"] = datacenter
	newRelicConfig.Labels["Envrionment"] = environment
	newRelicConfig.Labels["RoleTypeID"] = roletypeid

	app, _ := newrelic.NewApplication(newRelicConfig)

	return app
}

// Start the web service and return the error if any
func Start() error {
	newRelicApp := configureNewRelic()

	http.HandleFunc(wrapHandleFunc(newRelicApp, HealthPath, HealthGetHandler))
	http.HandleFunc(wrapHandleFunc(newRelicApp, EvaluatePath, EvaluatePostHandler))

	logger.Info.Println("Application Version: ", config.BinaryVersion)

	port := ":4001"
	logger.Info.Println("Starting server on ", port)
	return listenAndServe(port, nil)
}
