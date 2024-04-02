package main

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockProfiles struct {
	Reader io.ReadCloser
	Error  error
}

func (m MockProfiles) Open(username, password string) (io.ReadCloser, error) {
	return m.Reader, m.Error
}

func TestHandler(t *testing.T) {
	tests := []struct {
		name     string
		profiles Opener
		wantCode int
	}{
		{
			"authentication failed",
			MockProfiles{nil, fmt.Errorf("hello")},
			401,
		},
		{
			"authentication successful",
			MockProfiles{io.NopCloser(strings.NewReader("hello")), nil},
			200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/rest/GetUserlogin", nil)
			req.SetBasicAuth("some username", "some password")
			w := httptest.NewRecorder()
			Handler{tt.profiles}.ServeHTTP(w, req)
			resp := w.Result()
			if resp.StatusCode != tt.wantCode {
				t.Fatalf("Status Code: got %d want %d", resp.StatusCode, tt.wantCode)
			}
		})
	}
}
