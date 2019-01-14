package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckResponses(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, HealthPath, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthGetHandler)

	handler.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}

	expected := "Service Version: 0.0.dev"
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}
