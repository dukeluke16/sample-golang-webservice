package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/text/language"

	"github.com/newrelic/go-agent"
	"github.com/dukeluke16/sample-golang-webservice/config"
	"github.com/dukeluke16/sample-golang-webservice/logger"
)

var defaultDataFolder = "../data/"

const hazardousGoodsPolicyPath = "/hazardousGoodsPolicy.json"

const contentType = "application/json; charset=UTF-8"

var serverLanguageMatcher = language.NewMatcher([]language.Tag{
	language.Und,                  // clearly identify Undefined Tag
	language.Bulgarian,            // bg
	language.Czech,                // cs
	language.Danish,               // da
	language.German,               // de
	language.Greek,                // el
	language.English,              // en
	language.AmericanEnglish,      // en-US
	language.BritishEnglish,       // en-GB
	language.Spanish,              // es
	language.EuropeanSpanish,      // es-ES
	language.LatinAmericanSpanish, // es-LA
	language.Finnish,              // fi
	language.CanadianFrench,       // fr-CA
	language.French,               // fr
	language.Croatian,             // hr
	language.Hungarian,            // hu
	language.Italian,              // it
	language.Japanese,             // ja
	language.Korean,               // ko
	language.Lithuanian,           // lt
	language.Latvian,              // lv
	language.Dutch,                // nl
	language.Norwegian,            // no
	language.Polish,               // pl
	language.Portuguese,           // pt
	language.BrazilianPortuguese,  // pt-BR
	language.EuropeanPortuguese,   // pt-PT
	language.Romanian,             // ro
	language.Russian,              // ru
	language.Slovak,               // sk
	language.Swedish,              // sv
	language.Turkish,              // tr
	language.TraditionalChinese,   // zh-Hant, zh-CN
	language.SimplifiedChinese,    // zh-Hans, zh-TW
})

// EvaluatePath for endpoint
var EvaluatePath = "/policy/hazardousgoods/evaluate"

// EvaluatePostHandler for handling routed requests
func EvaluatePostHandler(w http.ResponseWriter, r *http.Request) {
	// Enforce POST Only
	if r.Method != http.MethodPost {
		genericStatusResponseError(w, r, http.StatusMethodNotAllowed)
		return
	}

	// Parse Accept-Language
	tag := parseAcceptLanguageHeader(r)
	if tag == language.Und {
		if r.Header.Get("Accept-Language") != "" {
			genericStatusResponseError(w, r, http.StatusNotAcceptable)
			return
		}
		tag = language.AmericanEnglish
	}

	evaluateLogicHandler(w, r, tag)
}

// evaluateLogicHandler for handling routed requests
func evaluateLogicHandler(w http.ResponseWriter, r *http.Request, tag language.Tag) {
	currentTransaction, _ := w.(newrelic.Transaction)

	// Empty Body receives Empty Response
	if r.Body == nil {
		emptyResponse(w, tag)
		return
	}

	bytes, _ := ioutil.ReadAll(r.Body)

	// Empty Body receives Empty Response
	body := string(bytes)
	if body == "" || body == "[]" {
		emptyResponse(w, tag)
		return
	}

	// Deserialize Array of AirportCodes
	// Parsing Error receives BadRequest Response
	var airportCodes []string
	err := json.Unmarshal(bytes, &airportCodes)
	if err != nil {
		genericStatusResponseError(w, r, http.StatusBadRequest)
		return
	}

	// Check if any AirportCode is inside USA
	airportInsideUSA := false
	for _, airportCode := range airportCodes {
		result, serviceError := checkAirportLocationCode(airportCode, currentTransaction)
		if serviceError != nil {
			genericStatusResponseError(w, r, http.StatusServiceUnavailable)
			return
		}

		airportInsideUSA = result == "US"
		if airportInsideUSA {
			break
		}
	}

	// Decide if policy is applicable
	if airportInsideUSA {
		hazardousGoodsPolicyResponse(w, r, tag)
		return
	}

	emptyResponse(w, tag)
	return
}

func buildErrorMessage(message string, r *http.Request) string {
	var buffer bytes.Buffer
	buffer.WriteString(strings.Replace(message, "\n", " |", -1))

	for k, v := range r.Header {
		buffer.WriteString(fmt.Sprint(" Request Header - key: ", k, " value: ", v, " |"))
	}

	return buffer.String()
}

func checkAirportLocationCode(inputCode string, txn newrelic.Transaction) (string, error) {
	airportCode := strings.ToUpper(inputCode)

	requestTemplate := `{
    "request": {
      "language": "en-US",
      "locationTypes": [
        "airport"
      ],
      "maxResultCount": 1,
      "searchText": "%s",
      "includeMinorLocations": true,
      "sources": ["IATA"]
    }
  }`
	apiRequestBody := fmt.Sprintf(requestTemplate, airportCode)

	client := &http.Client{}
	client.Transport = newrelic.NewRoundTripper(txn, nil)
	req, _ := http.NewRequest(http.MethodPost, config.LocationServicesURI(), strings.NewReader(apiRequestBody))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	// decode the json
	var j map[string]interface{}
	err = json.Unmarshal(body, &j)
	if err != nil {
		return "", err
	}

	airport := j["airports"].([]interface{})[0]
	returnAirportCode := airport.(map[string]interface{})["alternateIds"].([]interface{})[0].(map[string]interface{})["code"].(string)
	returnCountryCode := airport.(map[string]interface{})["country"].(map[string]interface{})["code"].(string)

	if strings.ToUpper(returnAirportCode) == airportCode {
		return returnCountryCode, nil
	}
	return "", errors.New("mismatched IATA code returned")
}

func hazardousGoodsPolicyResponse(w http.ResponseWriter, r *http.Request, tag language.Tag) {
	defaultResponse, err := getHazardousGoodsPolicy(tag)
	if err != nil {
		genericStatusResponseError(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Language", tag.String())
	io.WriteString(w, string(defaultResponse))
}

func emptyResponse(w http.ResponseWriter, tag language.Tag) {
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Language", tag.String())
	w.WriteHeader(http.StatusNoContent)
}

func getHazardousGoodsPolicy(tag language.Tag) (b []byte, err error) {
	policyFile := defaultDataFolder + tag.String() + hazardousGoodsPolicyPath
	return ioutil.ReadFile(policyFile)
}

func genericStatusResponseError(w http.ResponseWriter, r *http.Request, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
	message := fmt.Sprintln(statusCode, ":", http.StatusText(statusCode))
	errorMessage := buildErrorMessage(message, r)
	if logger.Initialized {
		logger.Error.Println(errorMessage)
	}
}

func parseAcceptLanguageHeader(r *http.Request) (tag language.Tag) {
	t, _, _ := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	parsedTag, _, _ := serverLanguageMatcher.Match(t...)
	return parsedTag
}
