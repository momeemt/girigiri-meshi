package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/momeemt/2000s/adapter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	router := adapter.Route()
	log.Printf("Listening on port :80")
	http.ListenAndServe(":80", router)
}
