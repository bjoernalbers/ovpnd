//go:build ignore

package main

import (
	"net/http"
	"os/exec"
	"testing"
	"time"
)

func TestIntegration(t *testing.T) {
	tests := []struct {
		name     string
		username string
		password string
		path     string
		want     int
	}{
		{
			"invalid url",
			"johndoe",
			"secret",
			"/missing",
			404,
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
				t.Fatalf("failed to build reqquest: %v", err)
			}
			req.SetBasicAuth(tt.username, tt.password)
			response, err := client.Do(req)
			if err != nil {
				t.Fatalf("request failed: %v", err)
			}
			if response.StatusCode != tt.want {
				t.Fatalf("response = %d, want: %d", response.StatusCode, tt.want)
			}
		})
	}
}
