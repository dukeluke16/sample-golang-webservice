package config

import (
	"net/url"
	"os"
)

// BinaryVersion of Application
var BinaryVersion = "0.0.dev"

// LocationServicesURIKey enivronment variable key
const LocationServicesURIKey = "TRAVEL_LOCATIONS_URI"

// LocationServicesURIValue enivronment variable key
var LocationServicesURIValue string

// LocationServicesURI method to return cached environment config
func LocationServicesURI() string {
	if len(LocationServicesURIValue) != 0 {
		return LocationServicesURIValue
	}

	LocationServicesURIValue = os.Getenv(LocationServicesURIKey)
	if len(LocationServicesURIValue) == 0 {
		panic("Location Services Environment Variable not Configured!")
	}

	return LocationServicesURIValue
}

// ProxyURI method to return parsed environment config
func ProxyURI(proxyURIKey string) *url.URL {
	envValue := os.Getenv(proxyURIKey)
	if len(envValue) == 0 {
		return nil
	}

	uri, _ := url.Parse(envValue)
	return uri
}
