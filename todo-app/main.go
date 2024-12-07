package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port, isPortDefined := os.LookupEnv("PORT")

	if !isPortDefined {
		port = "3000"
	}

	log.Printf("Server started in port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Printf("Failed to start server: %s", err)
		os.Exit(1)
	}
}
