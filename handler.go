package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Handler struct {
	db database
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		h.SendUnauthorized(w)
		return
	}
	profile, ok := h.db[username]
	if !ok {
		h.SendUnauthorized(w)
		return
	}
	if profile.Password != password {
		h.SendUnauthorized(w)
		return
	}
	file, err := os.Open(profile.Path)
	if err != nil {
		h.SendServerError(w)
		return
	}
	defer file.Close()
	io.Copy(w, file)
}

func (h Handler) SendUnauthorized(w http.ResponseWriter) {
	body := XmlError{Type: "Authorization Required", Message: "Invalid username or password"}
	http.Error(w, body.String(), http.StatusUnauthorized)
}

func (h Handler) SendServerError(w http.ResponseWriter) {
	body := XmlError{Type: "Internal Server Error", Message: "Failed to load profile"}
	http.Error(w, body.String(), http.StatusInternalServerError)
}

type XmlError struct {
	Type, Message string
}

func (x XmlError) String() string {
	const str = `<?xml version="1.0" encoding="UTF-8"?>
<Error>
<Type>%s</Type>
<Synopsis>REST method failed</Synopsis>
<Message>%s</Message>
</Error>`
	return fmt.Sprintf(str, x.Type, x.Message)
}
