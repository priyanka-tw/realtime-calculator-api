package main

import (
	"log"
	"net/http"
	"os"
	"realtime-calculator-api/router"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Println("initializing routes..")
	engine := router.InitializeRouter()

	log.Println("starting server on :8080")
	err := http.ListenAndServe(":"+port, engine)
	if err != nil {
		log.Fatal("unable to start server, err: ", err)
	}

}
