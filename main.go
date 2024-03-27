// ovpnd
package main

import (
	"flag"
	"log"
	"net/http"
)

const bodyUnauthorized = `<?xml version="1.0" encoding="UTF-8"?>
<Error>
<Type>Authorization Required</Type>
<Synopsis>REST method failed</Synopsis>
<Message>Invalid username or password</Message>
</Error>`

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "Address to listen on")
	flag.Parse()
	http.HandleFunc("/rest/GetUserlogin", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, bodyUnauthorized, http.StatusUnauthorized)
	})
	log.Fatal(http.ListenAndServe(*addr, nil))
}
