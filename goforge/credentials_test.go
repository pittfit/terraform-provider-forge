package goforge

import (
	"io"
	"net/http"
	"reflect"
	"testing"
)

// testCredentialsList
func TestCredentialsList(t *testing.T) {
	tc := SetUpTestClient(t)
	defer tc.TearDown()

	tc.Server.Mux.Handle("/api/v1/credentials", CredentialsListSuccessfulResponse(t))

	result, err := tc.Client.CredentialsList()
	if err != nil {
		t.Errorf("Error getting credentials: %v", err)
	}

	expected := []Credential{
		Credential{ID: 1, Type: "test", Name: "Personal"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CredentialsList returned %+v, expected %+v", result, expected)
	}
}

func CredentialsListSuccessfulResponse(t *testing.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testCommonHeaders(t, r)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		io.WriteString(w, `{"credentials":[{"id":1,"type":"test","name":"Personal"}]}`)
	})
}
