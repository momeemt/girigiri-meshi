package adapter

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/momeemt/2000s/adapter/handler"
)

func Route() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/hello", handler.HandleHello)
	return router
}
