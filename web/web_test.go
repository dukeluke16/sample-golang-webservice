package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	newrelic "github.com/newrelic/go-agent"
)

func TestStart(t *testing.T) {
	certDirectory = "../certs/"
	listenAndServe = func(addr string, handler http.Handler) error {
		return nil
	}

	result := Start()
	if result != nil {
		t.Error("Failure to startup!")
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
}

func TestBypassNewRelic(t *testing.T) {
	expectedPath := "/testPath"
	r, err := http.NewRequest(http.MethodGet, expectedPath, nil)
	r.Header.Set("Authorization", "Bearer abc123")
	r.Header.Set("correlationid", "123456789")
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	newRelicConfig := newrelic.NewConfig("Your_App_Name", "__YOUR_NEW_RELIC_LICENSE_KEY__")
	app, _ := newrelic.NewApplication(newRelicConfig)

	path, bypassHandler := bypassNewRelic(app, expectedPath, testHandler)

	if path != expectedPath {
		t.Errorf("handler returned wrong path: got %v want %v", path, expectedPath)
	}

	handler := http.HandlerFunc(bypassHandler)

	handler.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}
}
