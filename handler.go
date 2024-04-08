package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Opener interface {
	Open(username, password string) (io.ReadCloser, error)
}

type Handler struct {
	profiles Opener
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		h.SendUnauthorized(w)
		return
	}
	reader, err := h.profiles.Open(username, password)
	if err != nil {
		h.SendUnauthorized(w)
		return
	}
	defer reader.Close()
	io.Copy(w, reader)
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

type LogHandler struct {
	next http.Handler
}

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (s *StatusRecorder) WriteHeader(status int) {
	s.Status = status
	s.ResponseWriter.WriteHeader(status)
}

func (l LogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	recorder := &StatusRecorder{w, http.StatusOK}
	l.next.ServeHTTP(recorder, r)
	username, _, _ := r.BasicAuth()
	duration := time.Since(start)
	log.Printf("%s %s (%s) %d %s\n", r.Method, r.URL, username, recorder.Status, duration)
}
