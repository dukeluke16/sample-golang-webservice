package config

import (
	"os"
	"testing"
)

func TestBinaryVersionAccess(t *testing.T) {
	expected := "0.0.dev"
	if BinaryVersion != expected {
		t.Errorf("BinaryVersion does not match: got %v want %v",
			BinaryVersion, expected)
	}
}

func TestLocationServicesURISuccess(t *testing.T) {
	os.Clearenv()
	LocationServicesURIValue = ""

	expected := "TEST"
	os.Setenv(LocationServicesURIKey, expected)
	actual := LocationServicesURI()
	if actual != expected {
		t.Errorf("LocationServicesURI does not match: got %v want %v",
			actual, expected)
	}
}

func TestLocationServicesURICacheSuccess(t *testing.T) {
	os.Clearenv()
	LocationServicesURIValue = ""

	expected := "TEST"
	os.Setenv(LocationServicesURIKey, expected)
	actual1 := LocationServicesURI()
	if actual1 != expected {
		t.Errorf("LocationServicesURI does not match: got %v want %v",
			actual1, expected)
	}

	actual2 := LocationServicesURI()
	if actual2 != expected {
		t.Errorf("LocationServicesURI does not match: got %v want %v",
			actual2, expected)
	}
}

func TestLocationServicesURIFailure(t *testing.T) {
	os.Clearenv()
	LocationServicesURIValue = ""

	defer func() {
		if recover() == nil {
			t.Errorf("LocationServicesURI failed to detect no configuration!")
		}
	}()
	LocationServicesURI()
}

func TestProxyURISuccess(t *testing.T) {
	os.Clearenv()

	expected := "http://TEST"
	os.Setenv("ProxyURIKey", expected)
	actual := ProxyURI("ProxyURIKey")
	if actual.Host != "TEST" {
		t.Errorf("ProxyURI does not match: got %v want %v",
			actual, expected)
	}
}

func TestProxyURINotFoundFailure(t *testing.T) {
	os.Clearenv()

	actual := ProxyURI("ProxyURIKeyBad")
	if actual != nil {
		t.Errorf("ProxyURI does not match: got %v want %v",
			*actual, nil)
	}
}

func TestProxyURIParseFailure(t *testing.T) {
	os.Clearenv()

	expected := "foobie 123 !##$%^&"
	os.Setenv("ProxyURIKey", "")
	actual := ProxyURI("ProxyURIKey")
	if actual != nil {
		t.Errorf("ProxyURI does not match: got %v want %v",
			actual, expected)
	}
}
