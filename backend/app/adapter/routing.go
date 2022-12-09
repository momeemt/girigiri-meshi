package adapter

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/momeemt/2000s/adapter/handler"
	"github.com/momeemt/2000s/infra"
	"github.com/momeemt/2000s/usecase"
)

func Route() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/hello", handler.HandleHello)
	restaurantsHandler := handler.NewRestaurantsHandler(usecase.NewRestuarantsUsecase(infra.NewGooglePlacesApi(os.Getenv("GOOGLE_PLACES_API_KEY"))))
	router.HandleFunc("/restaurants", restaurantsHandler.HandleRestaurants).Methods("POST")
	return router
}
