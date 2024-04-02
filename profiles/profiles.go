// Package profiles provides access to OpenVPN connection profiles.
package profiles

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	profileSuffix  = ".ovpn"
	passwordSuffix = ".txt"
)

// Profiles represents a database of profiles paths.
type Profiles map[index]string

type index struct {
	username, password string
}

// ReadDir returns all Profiles from the given directory, where each profile
// has a corresponding password file.
func ReadDir(dir string) (Profiles, error) {
	profiles := Profiles{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return profiles, err
	}
	for _, entry := range entries {
		suffix := filepath.Ext(entry.Name())
		if suffix != profileSuffix || entry.IsDir() {
			continue
		}
		name := strings.TrimSuffix(entry.Name(), suffix)
		passwordFile, err := os.Open(filepath.Join(dir, name+passwordSuffix))
		if err != nil {
			continue
		}
		defer passwordFile.Close()
		password, err := io.ReadAll(passwordFile)
		passwordStr := strings.TrimSuffix(string(password), "\n")
		profiles[index{name, passwordStr}] = filepath.Join(dir, entry.Name())
	}
	return profiles, nil
}

// Open returns a profile reader for the given username and password.
// If the login is invalid, an error is returned.
func (p Profiles) Open(username, password string) (io.ReadCloser, error) {
	path, ok := p[index{username, password}]
	if !ok {
		return nil, fmt.Errorf("Authentication failed")
	}
	return os.Open(path)
}
