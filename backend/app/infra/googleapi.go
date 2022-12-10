package infra

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/momeemt/2000s/domain/model"
	"github.com/momeemt/2000s/domain/repository"
	"github.com/pkg/errors"
	"googlemaps.github.io/maps"
)

type googlePlacesApi struct {
	apiKey string
}

func NewGooglePlacesApi(apiKey string) repository.Restaurant {
	return &googlePlacesApi{
		apiKey: apiKey,
	}
}

// GetCloseTime implements repository.Restaurant
func (g *googlePlacesApi) GetNextCloseTime(restaurant model.Restaurant, now time.Time) (time.Time, error) {
	client, err := maps.NewClient(maps.WithAPIKey(g.apiKey))
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
func (g *googlePlacesApi) GetNearbyRestaurants(location model.Location) ([]model.Restaurant, error) {
	client, err := maps.NewClient(maps.WithAPIKey(g.apiKey))
	if err != nil {
		return nil, errors.Wrap(err, "error creating new client")
	}
	request := &maps.NearbySearchRequest{
		Location: &maps.LatLng{Lat: location.Latitude, Lng: location.Longitude},
		Language: "ja",
		OpenNow:  true,
		RankBy:   maps.RankByDistance,
		Type:     maps.PlaceTypeRestaurant,
	}
	response, err := client.NearbySearch(context.Background(), request)
	if err != nil {
		return nil, errors.Wrap(err, "error getting response")
	}
	var results []model.Restaurant
	c := make(chan model.Restaurant)
	for _, v := range response.Results {
		go func(searchResult maps.PlacesSearchResult) {
			c <- placesSearchResultToRestaurants(searchResult, g.apiKey)
		}(v)
	}
	for range response.Results {
		result := <-c
		results = append(results, result)
	}
	return results, nil
}

func placesSearchResultToRestaurants(searchResult maps.PlacesSearchResult, apiKey string) model.Restaurant {
	result := model.Restaurant{
		Name: searchResult.Name,
		Location: model.Location{
			Latitude:  searchResult.Geometry.Location.Lat,
			Longitude: searchResult.Geometry.Location.Lng,
		},
		PlaceId: searchResult.PlaceID,
		Rating:  searchResult.Rating,
	}
	if len(searchResult.Photos) > 0 {
		photoUrl, err := getUrl(searchResult.Photos[0].PhotoReference, apiKey)
		if err == nil {
			result.PhotoUrl = photoUrl
		}
	}
	return result
}

func getUrl(photoReference string, apiKey string) (string, error) {
	url := "https://maps.googleapis.com/maps/api/place/photo?maxwidth=1000&photo_reference=" + photoReference + "&key=" + apiKey
	for i := 0; i < 10; i++ {
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}}
		res, err := client.Get(url)
		if err != nil {
			return "", err
		}
		if res.StatusCode == 200 {
			return url, nil
		} else {
			url = res.Header.Get("Location")
		}
	}
	return "", fmt.Errorf("unable to get url")
}
