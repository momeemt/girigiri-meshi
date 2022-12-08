package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/momeemt/2000s/adapter/handler/apiio"
	"github.com/momeemt/2000s/domain/model"
	"github.com/momeemt/2000s/usecase"
)

type Restaurants struct {
	restaurantsUsecase usecase.Restaurants
}

func NewRestaurantsHandler(restaurantsUsecase usecase.Restaurants) Restaurants {
	return Restaurants{
		restaurantsUsecase: restaurantsUsecase,
	}
}

func (restaurants *Restaurants) HandleRestaurants(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var locationApiio apiio.Location
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ReturnErr(err, w)
		return
	}
	json.Unmarshal(body, &locationApiio)
	w.Header().Set("Content-Type", "application/json")
	availableRestaurants, err := restaurants.restaurantsUsecase.GetAvailableRestaurants(
		model.Location{
			Latitude:   locationApiio.Latitude,
			Longtitude: locationApiio.Longtitude,
		}, time.Now())
	if err != nil {
		ReturnErr(err, w)
		return
	}
	var response []apiio.Restaurant
	for _, v := range availableRestaurants {
		response = append(response, apiio.Restaurant{
			CloseTime: v.CloseTime,
			Location: apiio.Location{
				Latitude:   v.Location.Latitude,
				Longtitude: v.Location.Longtitude,
			},
			Name:     v.Name,
			PhotoUrl: &v.PhotoUrl,
			PlaceId:  v.PlaceId,
			Rating:   func(f float64) *float64 { return &f }(float64(v.Rating)),
		})
	}
	body, err = json.Marshal(response)
	if err != nil {
		ReturnErr(err, w)
		return
	}
	w.Write(body)
}
