package infra

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/momeemt/2000s/domain/model"
	"github.com/momeemt/2000s/domain/repository"
	"github.com/pkg/errors"
	"googlemaps.github.io/maps"
)

type googlePlacesApi struct{}

// GetCloseTime implements repository.Restaurant
func (googlePlacesApi) GetNextCloseTime(restaurant model.Restaurant, now time.Time) (time.Time, error) {
	client, err := maps.NewClient(maps.WithAPIKey("AIzaSyD73WTJXHUol9u8BsgINXK0DkdfqiQurd8"))
	if err != nil {
		return time.Time{}, errors.Wrap(err, "error creating new client")
	}

	request := &maps.PlaceDetailsRequest{
		PlaceID: restaurant.PlaceId,
		Fields:  []maps.PlaceDetailsFieldMask{maps.PlaceDetailsFieldMaskOpeningHours},
	}
	response, err := client.PlaceDetails(context.Background(), request)
	if err != nil {
		return time.Time{}, errors.Wrap(err, "error getting response")
	}
	for _, v := range response.OpeningHours.Periods {
		if v.Open.Day == now.Weekday() {
			closeTime, err := time.Parse("1504", v.Close.Time)
			if err != nil {
				return time.Time{}, errors.Wrap(err, "error parsing time string")
			}
			closeWeekday := v.Close.Day
			// 次の closeWeekday 曜日はいつか
			if dist := closeWeekday - now.Weekday(); dist < 0 {
				return time.Date(now.Year(), now.Month(), now.Day()+7-int(math.Abs(float64(dist))), closeTime.Hour(), closeTime.Minute(), 0, 0, now.Location()), nil
			} else {
				return time.Date(now.Year(), now.Month(), now.Day()+int(dist), closeTime.Hour(), closeTime.Minute(), 0, 0, now.Location()), nil
			}
		}
	}
	return time.Time{}, fmt.Errorf("no closing time defined")
}

// GetNearbyRestaurants implements repository.Restaurant
func (googlePlacesApi) GetNearbyRestaurants(location model.Location) ([]model.Restaurant, error) {
	client, err := maps.NewClient(maps.WithAPIKey("AIzaSyD73WTJXHUol9u8BsgINXK0DkdfqiQurd8"))
	if err != nil {
		return nil, errors.Wrap(err, "error creating new client")
	}
	request := &maps.NearbySearchRequest{
		Location: &maps.LatLng{Lat: location.Latitude, Lng: location.Longtitude},
		OpenNow:  true,
		RankBy:   maps.RankByDistance,
		Type:     maps.PlaceTypeRestaurant,
	}
	response, err := client.NearbySearch(context.Background(), request)
	if err != nil {
		return nil, errors.Wrap(err, "error getting response")
	}
	var results []model.Restaurant
	for _, v := range response.Results {
		result := model.Restaurant{
			Name: v.Name,
			Location: model.Location{
				Latitude:   v.Geometry.Location.Lat,
				Longtitude: v.Geometry.Location.Lng,
			},
			PlaceId: v.PlaceID,
			Rating:  v.Rating,
		}
		if len(v.Photos) > 0 {
			result.PhotoUrl = "https://maps.googleapis.com/maps/api/place/photo?maxwidth=1000&photo_reference=" + v.Photos[0].PhotoReference + "&key=AIzaSyD73WTJXHUol9u8BsgINXK0DkdfqiQurd8"
		}
		results = append(results, result)
	}
	return results, nil
}

func NewGooglePlacesApi() repository.Restaurant {
	return googlePlacesApi{}
}
