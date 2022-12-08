package adapter

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/momeemt/2000s/adapter/handler"
	"github.com/momeemt/2000s/infra"
	"github.com/momeemt/2000s/usecase"
)

func Route() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/hello", handler.HandleHello)
	restaurantsHandler := handler.NewRestaurantsHandler(usecase.NewRestuarantsUsecase(infra.NewGooglePlacesApi()))
	router.HandleFunc("/restaurants", restaurantsHandler.HandleRestaurants)
	return router
}
