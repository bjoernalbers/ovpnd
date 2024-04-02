package profiles

import (
	"io"
	"testing"
)

func TestOpen(t *testing.T) {
	tests := []struct {
		name     string
		username string
		password string
		want     string
		wantErr  bool
	}{
		{
			"invalid username and password",
			"wrong",
			"wrong",
			"",
			true,
		},
		{
			"invalid password",
			"johndoe",
			"wrong",
			"",
			true,
		},
		{
			"invalid username",
			"wrong",
			"secret",
			"",
			true,
		},
		{
			"valid username and password",
			"johndoe",
			"secret",
			"content of profile\n",
			false,
		},
	}
	p, err := ReadDir("../testdata")
	if err != nil {
		t.Fatalf("failed to load testdata: %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := p.Open(tt.username, tt.password)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Open() error = %v, wantErr: %v", err, tt.wantErr)
			}
			if tt.wantErr == true {
				return
			}
			got, err := io.ReadAll(r)
			if err != nil {
				t.Fatalf("failed to read profile: %v", err)
			}
			defer r.Close()
			if string(got) != tt.want {
				t.Fatalf("got: %q, want: %q", got, tt.want)
			}
		})
	}
}
