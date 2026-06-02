package auth

import (
	"net/http"
	"reflect"
	"testing"
)


func TestGetAPIKey(t *testing.T) {
	auth := http.Header{}
	auth.Add("Authorization", "ApiKey secret")
	got, _ := GetAPIKey(auth)
	want := "secret"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetAPIKeyMissing(t *testing.T) {
	auth := http.Header{}
	_, err := GetAPIKey(auth)
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected: %v, got: %v", ErrNoAuthHeaderIncluded, err)
	}
}