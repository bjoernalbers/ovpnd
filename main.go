// ovpnd
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const bodyUnauthorized = `<?xml version="1.0" encoding="UTF-8"?>
<Error>
<Type>Authorization Required</Type>
<Synopsis>REST method failed</Synopsis>
<Message>Invalid username or password</Message>
</Error>`

type Profile struct {
	Password, Content string
}

type ProfileHandler struct {
	db map[string]Profile
}

func (p ProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, bodyUnauthorized, http.StatusUnauthorized)
		return
	}
	profile, ok := p.db[username]
	if !ok {
		http.Error(w, bodyUnauthorized, http.StatusUnauthorized)
		return
	}
	if profile.Password != password {
		http.Error(w, bodyUnauthorized, http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, profile.Content)
}

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "Address to listen on")
	flag.Parse()
	handler := ProfileHandler{map[string]Profile{"johndoe": Profile{"secret", "content of profile\n"}}}
	http.Handle("/rest/GetUserlogin", handler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
