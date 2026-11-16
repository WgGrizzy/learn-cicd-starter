package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey abcdefg")
	key, err := GetAPIKey(headers)
	if err != nil {
		t.Error(err)
	}
	t.Log(key)
}
