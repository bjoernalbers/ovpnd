// ovpnd
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/bjoernalbers/ovpnd/profiles"
)

const (
	DefaultAddr    = ":80"
	DefaultTLSAddr = ":443"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("ovpnd: ")
}

func main() {
	addr := flag.String("addr", "", fmt.Sprintf("Address to listen on (default %s or %s with -no-tls)", DefaultTLSAddr, DefaultAddr))
	noTLS := flag.Bool("no-tls", false, "Disable TLS if behing TLS proxy")
	cert := flag.String("cert", "", "TLS certificate file (required unless running with -no-tls)")
	key := flag.String("key", "", "TLS key file (required unless running with -no-tls)")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("Please give directory with connection profiles (.ovpn) and password files (*.txt) as argument!")
	}
	p, err := profiles.ReadDir(args[0])
	if err != nil {
		log.Fatal(err)
	}
	if *addr == "" {
		if *noTLS {
			*addr = DefaultAddr
		} else {
			*addr = DefaultTLSAddr
		}
	}
	handler := Handler{p}
	http.Handle("/rest/GetUserlogin", handler)
	http.Handle("/rest/GetAutologin", handler)
	if *noTLS {
		log.Fatal(http.ListenAndServe(*addr, nil))
	} else {
		if *cert == "" || *key == "" {
			log.Fatal("both -cert and -key are required if not running with -no-tls")
		}
		log.Fatal(http.ListenAndServeTLS(*addr, *cert, *key, nil))
	}
}
