package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetA(t *testing.T) {
	validHeader := http.Header{}
	validHeader.Set("Authorization", "Bearer Q8rX2fA5vL9nTz3EwK1cYdU6oHbJmZpR")
	validToken, _ := GetAPIKey(validHeader)

	noAPIKey := http.Header{}
	noAPIKey.Set("Authorization", "Q8rX2fA5vL9nTz3EwK1cYdU6oHbJmZpR")

	noAuth := http.Header{}
	noAuth.Set("Authorization", "")

	cases := []struct {
		name        string
		InputHeader http.Header
		ExitToken   string
		wantErr     bool
	}{
		{
			name:        "Valid Header",
			InputHeader: validHeader,
			ExitToken:   validToken,
			wantErr:     false,
		},
		{
			name:        "No APIKey Bearer",
			InputHeader: noAPIKey,
			ExitToken:   "",
			wantErr:     true,
		},
		{
			name:        "No auth provided",
			InputHeader: noAuth,
			ExitToken:   "",
			wantErr:     true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			token, err := GetAPIKey(c.InputHeader)
			if (err != nil) != c.wantErr {
				fmt.Println(c.name)
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if token != c.ExitToken {
				t.Errorf("GetAPIKey() token = %v, ExitToken %v", token, c.ExitToken)
				return
			}
		})
	}
}
