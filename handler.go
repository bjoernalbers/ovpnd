package main

import (
	"io"
	"net/http"
	"os"
)

const (
	bodyUnauthorized = `<?xml version="1.0" encoding="UTF-8"?>
<Error>
<Type>Authorization Required</Type>
<Synopsis>REST method failed</Synopsis>
<Message>Invalid username or password</Message>
</Error>`
	bodyError = `<?xml version="1.0" encoding="UTF-8"?>
<Error>
<Type>Server Error</Type>
<Synopsis>REST method failed</Synopsis>
<Message>Failed to load profile</Message>
</Error>`
)

type Handler struct {
	db database
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, bodyUnauthorized, http.StatusUnauthorized)
		return
	}
	profile, ok := h.db[username]
	if !ok {
		http.Error(w, bodyUnauthorized, http.StatusUnauthorized)
		return
	}
	if profile.Password != password {
		http.Error(w, bodyUnauthorized, http.StatusUnauthorized)
		return
	}
	file, err := os.Open(profile.Path)
	if err != nil {
		http.Error(w, bodyError, 500)
		return
	}
	defer file.Close()
	io.Copy(w, file)
}
