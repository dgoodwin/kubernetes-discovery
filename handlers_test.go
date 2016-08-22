package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: add a token lookup interface for mocking
const testTokenId = "TOKENID"
const testToken = "TOKEN"

func TestClusterInfoIndex(t *testing.T) {
	tests := map[string]struct {
		url       string
		expStatus int
	}{
		"no token": {
			"/cluster-info/v1/",
			http.StatusForbidden,
		},
		"valid token": {
			fmt.Sprintf("/cluster-info/v1/?token-id=%s.%s", testTokenId, testToken),
			http.StatusOK,
		},
		"invalid token": {
			fmt.Sprintf("/cluster-info/v1/?token-id=JUNK", testTokenId, testToken),
			http.StatusForbidden,
		},
	}

	for name, test := range tests {
		t.Logf("Running test: %s", name)
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest("GET", test.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ClusterInfoIndex)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.expStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.expStatus)
		}

		// If we were expecting valid status validate the body:
		if test.expStatus == http.StatusOK {
			var ci ClusterInfo
			err := json.Unmarshal(rr.Body.Bytes(), &ci)
			if err != nil {
				t.Errorf("Unable to marshall response to JSON: error=%s body=%s", err, rr.Body.String())
				continue
			}
			if ci.RootCertificates == "" {
				t.Error("No root certificates in response")
			}
		}
	}
}
