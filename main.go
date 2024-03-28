// ovpnd
package main

import (
	"flag"
	"log"
	"net/http"
)

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
	handler := Handler{db}
	http.Handle("/rest/GetUserlogin", handler)
	http.Handle("/rest/GetAutologin", handler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
