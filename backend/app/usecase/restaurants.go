//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/mock_$GOPACKAGE/mock_$GOFILE
package usecase

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/momeemt/2000s/domain/model"
	"github.com/momeemt/2000s/domain/repository"
	"github.com/pkg/errors"
)

type Restaurants interface {
	// GetAvailableRestaurants は距離と閉店時刻を計算に入れて行くことができる飲食店を近い順番で並べて返す
	GetAvailableRestaurants(model.Location, time.Time) ([]model.Restaurant, error)
	GetRestaurantDetail(placeId string) (model.Restaurant, error)
}

func NewRestuarantsUsecase(restaurantRepository repository.Restaurant) Restaurants {
	return &restaurantsUsecase{
		restaurantRepository: restaurantRepository,
	}
}

type restaurantsUsecase struct {
	restaurantRepository repository.Restaurant
}

type restaurantWithIndex struct {
	restaurant model.Restaurant
	index      int
}

// GetAvailableRestaurants implements Restaurants
func (r *restaurantsUsecase) GetAvailableRestaurants(location model.Location, now time.Time) ([]model.Restaurant, error) {
	diff := func(dur time.Duration) time.Duration {
		if dur < 0 {
			return -dur
		} else {
			return dur
		}
	}(time.Time.Sub(now, time.Now()))
	restaurants, err := r.restaurantRepository.GetNearbyRestaurants(location, now, diff < time.Minute*10)
	var returnRestaurants []restaurantWithIndex
	if err != nil {
		return nil, errors.Wrap(err, "error while getting nearby restaurants")
	}
	c := make(chan struct {
		restaurantWithIndex
		error
	})
	for i, v := range restaurants {
		go func(restaurant model.Restaurant, index int, c chan struct {
			restaurantWithIndex
			error
		}) {
			closeTime, err := r.restaurantRepository.GetNextCloseTime(restaurant, now)
			if err != nil {
				c <- struct {
					restaurantWithIndex
					error
				}{restaurantWithIndex{}, err}
			} else {
				restaurant.CloseTime = closeTime
				c <- struct {
					restaurantWithIndex
					error
				}{restaurantWithIndex{
					restaurant: restaurant,
					index:      index,
				}, nil}
			}
		}(v, i, c)
	}

	for range restaurants {
		result := <-c
		if result.error != nil {
			// log したい
			fmt.Printf("%v\n", result.error)
			continue
		} else {
			v := result.restaurantWithIndex.restaurant
			duration, err := time.ParseDuration(fmt.Sprint(distance(location.Latitude, location.Longitude, v.Location.Latitude, v.Location.Longitude)/4) + "h")
			if err != nil {
				fmt.Printf("%v\n", result.error)
				continue
			}
			//徒歩で時速4kmとして直線距離で計算した到着時刻+30分後に着かない場合弾く
			arrivalTime := now.Add(duration + 30*time.Minute)
			if arrivalTime.Before(v.CloseTime) {
				returnRestaurants = append(returnRestaurants, result.restaurantWithIndex)
			}
		}
	}

	sort.Slice(returnRestaurants, func(i, j int) bool { return returnRestaurants[i].index < returnRestaurants[j].index })
	var rRestaurants []model.Restaurant
	for _, v := range returnRestaurants {
		rRestaurants = append(rRestaurants, v.restaurant)
	}
	return rRestaurants, nil
}

func (r *restaurantsUsecase) GetRestaurantDetail(placeId string) (model.Restaurant, error) {
	detailedRestaurant, err := r.restaurantRepository.GetRestaurantDetail(placeId)
	if err != nil {
		errors.Wrap(err, "error while getting detailed restaurants")
	}
	return detailedRestaurant, nil
}

//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
//:::                                                                         :::
//:::  This routine calculates the distance between two points (given the     :::
//:::  latitude/longitude of those points). It is being used to calculate     :::
//:::  the distance between two locations using GeoDataSource (TM) products  :::
//:::                                                                         :::
//:::  Definitions:                                                           :::
//:::    South latitudes are negative, east longitudes are positive           :::
//:::                                                                         :::
//:::  Passed to function:                                                    :::
//:::    lat1, lon1 = Latitude and Longitude of point 1 (in decimal degrees)  :::
//:::    lat2, lon2 = Latitude and Longitude of point 2 (in decimal degrees)  :::
//:::    unit = the unit you desire for results                               :::
//:::           where: 'M' is statute miles (default)                         :::
//:::                  'K' is kilometers                                      :::
//:::                  'N' is nautical miles                                  :::
//:::                                                                         :::
//:::  Worldwide cities and other features databases with latitude longitude  :::
//:::  are available at https://www.geodatasource.com                         :::
//:::                                                                         :::
//:::  For enquiries, please contact sales@geodatasource.com                  :::
//:::                                                                         :::
//:::  Official Web site: https://www.geodatasource.com                       :::
//:::                                                                         :::
//:::               GeoDataSource.com (C) All Rights Reserved 2022            :::
//:::                                                                         :::
//:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

func distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515
	dist = dist * 1.609344

	return dist
}
