package auth

import (
	"net/http"
	"testing"
)

func TestAuthHeaderEmpty(t *testing.T) {
	headers:= http.Header{}
	headers.Set("Authorization", "")
	
	_, error := GetAPIKey(headers)
	
	if error != ErrNoAuthHeaderIncluded {
		t.Errorf("got %d, wanted %q", error, ErrNoAuthHeaderIncluded)
	}
}

func TestAuthHeader(t *testing.T) {
	headers:= http.Header{}
	headers.Set("Authorization", "ApiKey FakeKeyValue")

	_, error := GetAPIKey(headers)
	
	if error != nil {
		t.Errorf("got %d, wanted %q", error, "malformed authorization header")
	}
}

func TestApiKeyValue(t *testing.T) {
	headers:= http.Header{}
	headers.Set("Authorization", "ApiKey FakeKeyValue")

	apiKey, error := GetAPIKey(headers)
	
	if apiKey == "FakeKeyValue" {
		t.Errorf("got %d, wanted %q", error, "FakeKeyValue")
	}
}