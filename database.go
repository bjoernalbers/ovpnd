package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Profile struct {
	Path, Password string
}

type database map[string]Profile

func buildDatabase(dir string) (database, error) {
	db := database{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return db, err
	}
	for _, entry := range entries {
		suffix := filepath.Ext(entry.Name())
		if suffix != ".ovpn" {
			continue
		}
		name := strings.TrimSuffix(entry.Name(), suffix)
		passwordFile, err := os.Open(filepath.Join(dir, name+".txt"))
		if err != nil {
			continue
		}
		defer passwordFile.Close()
		password, err := io.ReadAll(passwordFile)
		db[name] = Profile{filepath.Join(dir, entry.Name()), strings.TrimSuffix(string(password), "\n")}
	}
	return db, nil
}