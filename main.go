package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"01-Login/platform/authenticator"
	"01-Login/platform/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to laod the env vars: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to intialize the authenticator: %v", err)
	}

	rtr := router.New(auth)

	log.Print("Server listening on http://localhost:3000")
	if err := http.ListenAndServe("0.0.0.0:3000", rtr); err != nil {
		log.Fatalf("Error loading HTTP Server: %v", err)
	}
}
