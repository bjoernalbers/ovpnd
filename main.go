// ovpnd
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/bjoernalbers/ovpnd/profiles"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("ovpnd: ")
}

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "Address to listen on")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("Please give directory with connection profiles (.ovpn) and password files (*.txt) as argument!")
	}
	p, err := profiles.ReadDir(args[0])
	if err != nil {
		log.Fatal(err)
	}
	handler := Handler{p}
	http.Handle("/rest/GetUserlogin", handler)
	http.Handle("/rest/GetAutologin", handler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
