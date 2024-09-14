package auth

import (
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetAPIKeySuccess(t *testing.T) {
	authHeader := httptest.NewRecorder().Header()
	authHeader.Add("Authorization", "ApiKey 12345")

	got, _ := GetAPIKey(authHeader)
	want := "12345"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
