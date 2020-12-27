package main

import (
	"log"
	"net/http"
	"realtime-calculator-api/router"
)

func main() {
	log.Println("initializing routes..")
	engine := router.InitializeRouter()

	log.Println("starting server on :8080")
	err := http.ListenAndServe(":8080", engine)
	if err != nil{
		log.Fatal("unable to start server, err: ", err)
	}

}
