package web

import (
	"fmt"
	"io"
	"net/http"

	"github.com/dukeluke16/sample-golang-webservice/config"
)

// HealthPath for endpoint
var HealthPath = "/health"

// HealthGetHandler for handling routed requests
func HealthGetHandler(w http.ResponseWriter, r *http.Request) {
	packageVersion := fmt.Sprintf("Service Version: %v", config.BinaryVersion)
	io.WriteString(w, packageVersion)
}
