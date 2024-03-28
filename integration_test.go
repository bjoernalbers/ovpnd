//go:build ignore

package main

import (
	"io"
	"net/http"
	"os/exec"
	"testing"
	"time"
)

func TestIntegration(t *testing.T) {
	const unauthorizedResponse = `<?xml version="1.0" encoding="UTF-8"?>
<Error>
<Type>Authorization Required</Type>
<Synopsis>REST method failed</Synopsis>
<Message>Invalid username or password</Message>
</Error>
`
	tests := []struct {
		name     string
		username string
		password string
		path     string
		wantCode int
		wantBody string
	}{
		{
			"unknown path",
			"",
			"",
			"/missing",
			404,
			"404 page not found\n",
		},
		{
			"wrong username",
			"wronguser",
			"secret",
			"/rest/GetUserlogin",
			401,
			unauthorizedResponse,
		},
		{
			"wrong password",
			"johndoe",
			"wrongpassword",
			"/rest/GetUserlogin",
			401,
			unauthorizedResponse,
		},
		{
			"right username and password",
			"johndoe",
			"secret",
			"/rest/GetUserlogin",
			200,
			"content of profile\n",
		},
	}
	cmd := exec.Command("./ovpnd")
	if err := cmd.Start(); err != nil {
		t.Fatalf("%s failed: %v", cmd, err)
	}
	time.Sleep(1 * time.Second) // wait for server to start
	defer cmd.Process.Kill()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{}
			req, err := http.NewRequest("GET", "http://127.0.0.1:8080"+tt.path, nil)
			if err != nil {
				t.Fatalf("Failed to build request: %v", err)
			}
			req.SetBasicAuth(tt.username, tt.password)
			response, err := client.Do(req)
			if err != nil {
				t.Fatalf("Request failed: %v", err)
			}
			defer response.Body.Close()
			if got := response.StatusCode; got != tt.wantCode {
				t.Fatalf("Status Code = %d, want: %d", got, tt.wantCode)
			}
			if got, _ := io.ReadAll(response.Body); string(got) != tt.wantBody {
				t.Fatalf("Body:\ngot:  %q\nwant: %q", got, tt.wantBody)
			}
		})
	}
}
