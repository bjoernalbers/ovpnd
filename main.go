// ovpnd
package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const bodyUnauthorized = `<?xml version="1.0" encoding="UTF-8"?>
<Error>
<Type>Authorization Required</Type>
<Synopsis>REST method failed</Synopsis>
<Message>Invalid username or password</Message>
</Error>`

type Profile struct {
	Path, Password string
}

type database map[string]Profile

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, bodyUnauthorized, http.StatusUnauthorized)
		return
	}
	profile, ok := db[username]
	if !ok {
		http.Error(w, bodyUnauthorized, http.StatusUnauthorized)
		return
	}
	if profile.Password != password {
		http.Error(w, bodyUnauthorized, http.StatusUnauthorized)
		return
	}
	file, _ := os.Open(profile.Path)
	// TODO: Handle error!
	defer file.Close()
	io.Copy(w, file)
}

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

func main() {
	dir := flag.String("dir", "", "Directory with configuration profiles (.ovpn) and password files (.txt)")
	addr := flag.String("addr", "127.0.0.1:8080", "Address to listen on")
	flag.Parse()
	if *dir == "" {
		log.Fatal("dir not set")
	}
	db, err := buildDatabase(*dir)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/rest/GetUserlogin", db)
	http.Handle("/rest/GetAutologin", db)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
