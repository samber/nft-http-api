package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupApi(listen string, unixSocket string, tlsCertFile string, tlsKeyFile string) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Recovery())

	registerRoutes(router)

	var err error
	if listen != "" && tlsKeyFile != "" && tlsCertFile != "" {
		err = router.RunTLS(listen, tlsCertFile, tlsKeyFile)
	} else if unixSocket != "" {
		err = router.RunUnix(unixSocket)
	} else if listen != "" {
		err = router.Run(listen)
	} else {
		log.Fatalf("Please provide --listen or --unix\n")
	}

	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
