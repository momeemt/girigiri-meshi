package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/momeemt/2000s/adapter/handler/apiio"
	"github.com/momeemt/2000s/domain/model"
	"github.com/momeemt/2000s/usecase"
)

type RestaurantsDetails struct {
	restaurantsUsecase usecase.Restaurants
}

func NewRestaurantsDetailsHandler(restaurantUsecase usecase.Restaurants) RestaurantsDetails {
	return RestaurantsDetails{
		restaurantsUsecase: restaurantUsecase,
	}
}

func (restaurantsDetails *RestaurantsDetails) HandleRestaurantsDetails(w http.ResponseWriter, r *http.Request) {
	placeId, ok := r.URL.Query()["placeId"]
	if !ok {
		ReturnErr(fmt.Errorf("no placeId parameter"), w)
	}
	restaurant, err := restaurantsDetails.restaurantsUsecase.GetRestaurantDetail(placeId[0])
	if err != nil {
		ReturnErr(err, w)
		return
	}
	response := apiio.RestaurantDetail{
		PhotoUrls: &restaurant.PhotoUrls,
		PlaceId:   &restaurant.PlaceId,
		Reviews: func(r model.Restaurant) *[]apiio.Review {
			var result []apiio.Review
			for i, v := range r.Reviews {
				result = append(result, apiio.Review{
					AuthorName:      v.AuthorName,
					ProfilePhotoUrl: &r.Reviews[i].ProfilePhotoUrl,
					Rating:          v.Rating,
					Text:            v.Text,
					Time:            &r.Reviews[i].Time,
				})
			}
			return &result
		}(restaurant),
		Url:              &restaurant.Url,
		UserRatingsTotal: &restaurant.UserRatingsTotal,
		Website:          &restaurant.Website,
	}
	body, err := json.Marshal(response)
	if err != nil {
		ReturnErr(err, w)
		return
	}
	w.Write(body)
}
