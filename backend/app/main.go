package main

import (
	"log"
	"net/http"

	"github.com/momeemt/2000s/adapter"
)

func main() {
	router := adapter.Route()
	log.Printf("Listening on port :80")
	http.ListenAndServe(":80", router)
}
