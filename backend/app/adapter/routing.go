package adapter

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/momeemt/2000s/adapter/handler"
	"github.com/momeemt/2000s/infra"
	"github.com/momeemt/2000s/usecase"
	"github.com/rs/cors"
)

func Route() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/hello", handler.HandleHello)
	restaurantsDetailsHandler := handler.NewRestaurantsDetailsHandler(usecase.NewRestuarantsUsecase(infra.NewGooglePlacesApi(os.Getenv("GOOGLE_PLACES_API_KEY"))))
	router.HandleFunc("/restaurants/details", restaurantsDetailsHandler.HandleRestaurantsDetails)
	restaurantsHandler := handler.NewRestaurantsHandler(usecase.NewRestuarantsUsecase(infra.NewGooglePlacesApi(os.Getenv("GOOGLE_PLACES_API_KEY"))))
	router.HandleFunc("/restaurants", restaurantsHandler.HandleRestaurants).Methods("POST")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:3000", "https://girigirimeshi.netlify.app"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	// Insert the middleware
	corsRouter := c.Handler(router)
	return corsRouter
}
