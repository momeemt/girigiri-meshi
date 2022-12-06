package main

import (
	"log"
	"net/http"

	"github.com/momeemt/2000s/handleFuncs"
)

func main() {
	log.Print("Server Starting...\n")
	http.HandleFunc("/hello", handleFuncs.GetMessage)

	http.ListenAndServe(":80", nil)
}
