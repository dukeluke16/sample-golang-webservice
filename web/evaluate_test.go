package web

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/dukeluke16/sample-golang-webservice/config"
	"github.com/dukeluke16/sample-golang-webservice/logger"
)

func TestEvaluateResponse_HTTPMethodFailure(t *testing.T) {
	w := setupRequestAndServe(http.MethodGet, nil, nil)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}

	expected := http.StatusText(http.StatusMethodNotAllowed) + "\n"
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}

func TestEvaluateResponse_NilRequestBody(t *testing.T) {
	w := setupPostRequestAndServe(nil, nil)

	if w.Code != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusNoContent)
	}

	if len(w.Body.Bytes()) > 0 {
		t.Errorf("Unexpected Body Content")
	}
}

func TestEvaluateResponse_EmptyRequestBody(t *testing.T) {
	w := setupPostRequestAndServe(strings.NewReader(`[]`), nil)

	if w.Code != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusNoContent)
	}

	if len(w.Body.Bytes()) > 0 {
		t.Errorf("Unexpected Body Content")
	}
}

func TestEvaluateResponse_BadRequestBody(t *testing.T) {
	w := setupPostRequestAndServe(strings.NewReader(`["foo": "bar"]`), nil)

	if w.Code != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusBadRequest)
	}

	if w.Body.String() != "Bad Request\n" {
		t.Errorf("Unexpected body response: got %v  want %v", w.Body.String(), "Bad Request\n")
	}
}

func TestEvaluateResponseSuccessSEA(t *testing.T) {
	ts := setupFakeServerUSA()
	w := setupPostRequestAndServe(strings.NewReader(`["sea"]`), nil)
	defer ts.Close()

	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}

	type hazardousGoodsPolicyModel struct {
		Code  string
		Alert string
		Title string
		Body  []string
	}
	var parsedData []hazardousGoodsPolicyModel
	if err := json.Unmarshal(w.Body.Bytes(), &parsedData); err != nil {
		t.Errorf("Parsing error.")
	}

	if len(parsedData) != 1 {
		t.Errorf("Array size error.")
	}

	policy := parsedData[0]
	expectedCode := "US"
	if policy.Code != expectedCode {
		t.Errorf("handler returned unexpected code: got %v want %v",
			policy.Code, expectedCode)
	}

	expectedAlert := "By completing this booking, you agree to the fare rules and restrictions and hazardous goods policy."
	if policy.Alert != expectedAlert {
		t.Errorf("handler returned unexpected alert: got %v want %v",
			policy.Alert, expectedAlert)
	}

	expectedTitle := "Hazardous Materials Restrictions"
	if policy.Title != expectedTitle {
		t.Errorf("handler returned unexpected title: got %v want %v",
			policy.Title, expectedTitle)
	}

	expectedBody := []string{"Federal law forbids the carriage of hazardous materials aboard aircraft in your luggage or on your person. A violation can result in five years' imprisonment and penalties of $250,000 or more (49 U.S.C. 5124). Hazardous materials include explosives, compressed gases, flammable liquids and solids, oxidizers, poisons, corrosives and radioactive materials. Examples: Paints, lighter fluid, fireworks, tear gases, oxygen bottles, and radio-pharmaceuticals.",
		"There are special exceptions for small quantities (up to 70 ounces total) of medicinal and toilet articles carried in your luggage and certain smoking materials carried on your person. For further information contact your airline representative."}

	if policy.Body[0] != expectedBody[0] || policy.Body[1] != expectedBody[1] {
		t.Errorf("handler returned unexpected body: got \n%v \n\nwant \n%v",
			policy.Body, expectedBody)
	}
}

func TestEvaluateResponseSuccessLCY(t *testing.T) {
	ts := setupFakeServerEMEA()
	w := setupPostRequestAndServe(strings.NewReader(`["lcy"]`), nil)
	defer ts.Close()

	if w.Code != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusNoContent)
	}
}

func TestEvaluateResponseSuccessLCYMismatch(t *testing.T) {
	ts := setupFakeServer(`{
			  "airports": [
			    {
			      "isMajor": true,
			      "name": "London City Arpt",
			      "rank": 1,
			      "point": {
			        "lon": 0.04977,
			        "lat": 51.5032
			      },
			      "country": {
			        "code": "GB",
			        "regionId": 2,
			        "name": "United Kingdom",
			        "alternateIds": []
			      },
			      "city": {
			        "name": "London",
			        "alternateIds": []
			      },
			      "state": {
			        "name": "England",
			        "alternateIds": []
			      },
			      "alternateIds": [
			        {
			          "source": "IATA",
			          "code": "foobar"
			        }
			      ],
			      "language": "en-us"
			    }
			  ],
			  "hubs": [],
			  "carRentals": [],
			  "heliports": [],
			  "railStations": [],
			  "busStations": [],
			  "summary": {
			    "hasError": false,
			    "hasWarning": false,
			    "messages": []
			  }
			}`)

	w := setupPostRequestAndServe(strings.NewReader(`["lcy"]`), nil)
	defer ts.Close()

	if w.Code != http.StatusServiceUnavailable {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusServiceUnavailable)
	}

	if w.Body.String() != "Service Unavailable\n" {
		t.Errorf("Unexpected body response: got %v  want %v", w.Body.String(), "Service Unavailable\n")
	}
}

func TestEvaluateResponseFailureBadJSON(t *testing.T) {
	ts := setupFakeServer(`foobar {jack} "be" "nimble"`)
	w := setupPostRequestAndServe(strings.NewReader(`["lcy"]`), nil)
	defer ts.Close()

	if w.Code != http.StatusServiceUnavailable {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusServiceUnavailable)
	}

	if w.Body.String() != "Service Unavailable\n" {
		t.Errorf("Unexpected body response: got %v  want %v", w.Body.String(), "Service Unavailable\n")
	}
}

func TestEvaluateResponseDefaultPolicyNotFound(t *testing.T) {
	ts := setupFakeServerUSA()
	defaultDataFolder = "../badDataFolder/"
	w := setupPostRequestAndServe(strings.NewReader(`["sea"]`), nil)
	defer ts.Close()

	if w.Code != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusInternalServerError)
	}

	if w.Body.String() != "Internal Server Error\n" {
		t.Errorf("Unexpected body response: got %v  want %v", w.Body.String(), "Internal Server Error\n")
	}
}

func TestEvaluateResponseServiceNotAvailable(t *testing.T) {
	setServiceEndpoint("http://localhost")
	w := setupPostRequestAndServe(strings.NewReader(`["lcy"]`), nil)

	if w.Code != http.StatusServiceUnavailable {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusServiceUnavailable)
	}

	if w.Body.String() != "Service Unavailable\n" {
		t.Errorf("Unexpected body response: got %v  want %v", w.Body.String(), "Service Unavailable\n")
	}
}

func init() {
	resetServiceEndpoint()
	logger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	defaultDataFolder = "../data/"
}

func resetServiceEndpoint() {
	os.Clearenv()
	config.LocationServicesURIValue = ""
}

func setupFakeServerEMEA() *httptest.Server {
	return setupFakeServer(`{
			  "airports": [
			    {
			      "isMajor": true,
			      "name": "London City Arpt",
			      "rank": 1,
			      "point": {
			        "lon": 0.04977,
			        "lat": 51.5032
			      },
			      "country": {
			        "code": "GB",
			        "regionId": 2,
			        "name": "United Kingdom",
			        "alternateIds": []
			      },
			      "city": {
			        "name": "London",
			        "alternateIds": []
			      },
			      "state": {
			        "name": "England",
			        "alternateIds": []
			      },
			      "alternateIds": [
			        {
			          "source": "IATA",
			          "code": "LCY"
			        }
			      ],
			      "language": "en-us"
			    }
			  ],
			  "hubs": [],
			  "carRentals": [],
			  "heliports": [],
			  "railStations": [],
			  "busStations": [],
			  "summary": {
			    "hasError": false,
			    "hasWarning": false,
			    "messages": []
			  }
			}`)
}

func setupFakeServerUSA() *httptest.Server {
	return setupFakeServer(`{
			"airports": [
				{
					"isMajor": true,
					"name": "Seattle-Tacoma",
					"rank": 1,
					"point": {
						"lon": -122.2,
						"lat": 47.4
					},
					"country": {
						"code": "US",
						"regionId": 2,
						"name": "United States",
						"alternateIds": []
					},
					"city": {
						"name": "Seattle",
						"alternateIds": []
					},
					"state": {
						"name": "Washington",
						"alternateIds": []
					},
					"alternateIds": [
						{
							"source": "IATA",
							"code": "SEA"
						}
					],
					"language": "en-us"
				}
			],
			"hubs": [],
			"carRentals": [],
			"heliports": [],
			"railStations": [],
			"busStations": [],
			"summary": {
				"hasError": false,
				"hasWarning": false,
				"messages": []
			}
		}`)
}

func setupFakeServer(data string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept-Language", "application/json")
		fmt.Fprintln(w, data)
	}))
	setServiceEndpoint(ts.URL)
	return ts
}

func setupPostRequestAndServe(dataReader io.Reader, tag *string) *httptest.ResponseRecorder {
	return setupRequestAndServe(http.MethodPost, dataReader, tag)
}

func setupRequestAndServe(method string, dataReader io.Reader, tag *string) *httptest.ResponseRecorder {
	r, _ := http.NewRequest(method, EvaluatePath, dataReader)
	r.Header.Set("Authorization", "Bearer abc123")
	r.Header.Set("correlationid", "123456789")
	if tag != nil {
		r.Header.Set("Accept-Language", *tag)
	}

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(EvaluatePostHandler)
	handler.ServeHTTP(w, r)

	return w
}

func setServiceEndpoint(endpoint string) {
	resetServiceEndpoint()
	os.Setenv(config.LocationServicesURIKey, endpoint)
}
