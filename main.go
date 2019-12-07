package main

import (
	"crypto/tls"
	"flag"
	"log"

	"github.com/samber/nft-http-api/nft"
)

var (
	Version string
	Build   string
)

var (
	listen     = flag.String("listen", "", "Address to listen on for API")
	unixSocket = flag.String("unix", "", "Unix socket to listen on for API")
	tlsKey     = flag.String("tls-key", "", "Name of the server key file (including full path) if the server requires TLS client authentication")
	tlsCert    = flag.String("tls-cert", "", "Name of the server certificate file (including full path) if the server requires TLS client authentication")
)

func loadCertificates() *tls.Certificate {
	if (*tlsKey != "") != (*tlsCert != "") {
		log.Fatal("TLS server key file and cert file should both be present")
	}
	if *tlsKey != "" && *tlsCert != "" {
		cert, err := tls.LoadX509KeyPair(*tlsCert, *tlsKey)
		if err != nil {
			log.Fatalf("Couldn't load TLS server key pair, err: %s", err)
		}
		return &cert
	}

	return nil
}

func main() {
	flag.Parse()
	_ = loadCertificates()

	_ = nft.Connect()

	SetupApi(*listen, *unixSocket, *tlsCert, *tlsKey)
}
