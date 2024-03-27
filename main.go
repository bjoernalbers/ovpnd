// ovpnd
package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:8080", "Address to listen on")
	flag.Parse()
	log.Fatal(http.ListenAndServe(*addr, nil))
}
