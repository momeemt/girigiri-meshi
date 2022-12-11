package infra

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"sort"
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
	if response.OpeningHours != nil {
		for _, v := range response.OpeningHours.Periods {
			if v.Open.Day == now.Weekday() {
				openTime, err := time.Parse("1504", v.Open.Time)
				if err != nil {
					return time.Time{}, errors.Wrap(err, "error parsing time string")
				}
				if openTime.After(now) {
					return time.Time{}, fmt.Errorf("restaurant not open")
				}
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
	} else {
		return time.Time{}, fmt.Errorf("no closing time defined")
	}

	return time.Time{}, fmt.Errorf("unknown error")
}

type restaurantWithIndex struct {
	model.Restaurant
	int
}

// GetNearbyRestaurants implements repository.Restaurant
func (g *googlePlacesApi) GetNearbyRestaurants(location model.Location, timeToSearch time.Time, isNow bool) ([]model.Restaurant, error) {
	client, err := maps.NewClient(maps.WithAPIKey(g.apiKey))
	if err != nil {
		return nil, errors.Wrap(err, "error creating new client")
	}
	var request *maps.NearbySearchRequest
	if isNow {
		request = &maps.NearbySearchRequest{
			Location: &maps.LatLng{Lat: location.Latitude, Lng: location.Longitude},
			Language: "ja",
			OpenNow:  true,
			RankBy:   maps.RankByDistance,
			Type:     maps.PlaceTypeRestaurant,
		}
	} else {
		request = &maps.NearbySearchRequest{
			Location: &maps.LatLng{Lat: location.Latitude, Lng: location.Longitude},
			Language: "ja",
			RankBy:   maps.RankByDistance,
			Type:     maps.PlaceTypeRestaurant,
		}
	}
	response, err := client.NearbySearch(context.Background(), request)
	if err != nil {
		return nil, errors.Wrap(err, "error getting response")
	}
	var results []model.Restaurant
	c := make(chan restaurantWithIndex)
	for i, v := range response.Results {
		go func(searchResult maps.PlacesSearchResult, index int) {
			c <- restaurantWithIndex{placesSearchResultToRestaurants(searchResult, g.apiKey), index}
		}(v, i)
	}
	var resultsNotSorted []restaurantWithIndex
	for range response.Results {
		result := <-c
		resultsNotSorted = append(resultsNotSorted, result)
	}
	sort.Slice(resultsNotSorted, func(i, j int) bool { return resultsNotSorted[i].int < resultsNotSorted[j].int })
	for _, v := range resultsNotSorted {
		results = append(results, v.Restaurant)
	}
	return results, nil
}

func (g *googlePlacesApi) GetRestaurantDetail(placeId string) (model.Restaurant, error) {
	client, err := maps.NewClient(maps.WithAPIKey(g.apiKey))
	if err != nil {
		return model.Restaurant{}, errors.Wrap(err, "error creating new client")
	}

	request := &maps.PlaceDetailsRequest{
		PlaceID:  placeId,
		Language: "ja",
	}
	response, err := client.PlaceDetails(context.Background(), request)
	if err != nil {
		return model.Restaurant{}, errors.Wrap(err, "error getting response")
	}
	returnRestaurant := model.Restaurant{
		Name: response.Name,
		Location: model.Location{
			Latitude:  response.Geometry.Location.Lat,
			Longitude: response.Geometry.Location.Lng,
		},
		PlaceId: response.PlaceID,
		Rating:  response.Rating,
		Reviews: func(r []maps.PlaceReview) []model.Review {
			var reviews []model.Review
			for _, v := range r {
				reviews = append(reviews, model.Review{
					AuthorName:      v.AuthorName,
					ProfilePhotoUrl: v.AuthorProfilePhoto,
					Rating:          v.Rating,
					Time:            time.Unix(int64(v.Time), 0),
					Text:            v.Text,
				})
			}
			return reviews
		}(response.Reviews),
		Url:              response.URL,
		UserRatingsTotal: response.UserRatingsTotal,
		Website:          response.Website,
	}
	c := make(chan struct {
		string
		error
	})
	for _, v := range response.Photos {
		go func(photoReference string, apiKey string, c chan struct {
			string
			error
		}) {
			photoUrl, err := getUrl(photoReference, apiKey)
			if err != nil {
				c <- struct {
					string
					error
				}{"", err}
			} else {
				c <- struct {
					string
					error
				}{photoUrl, nil}
			}
		}(v.PhotoReference, g.apiKey, c)
	}
	for range response.Photos {
		result := <-c
		if result.error != nil {
			continue
		} else {
			returnRestaurant.PhotoUrls = append(returnRestaurant.PhotoUrls, result.string)
		}
	}
	return returnRestaurant, nil
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
