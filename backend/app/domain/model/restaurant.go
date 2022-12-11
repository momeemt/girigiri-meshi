package model

import "time"

type Restaurant struct {
	Name             string
	Location         Location
	CloseTime        time.Time
	PhotoUrl         string
	PhotoUrls        []string
	PlaceId          string //dependent to Google Maps API
	Rating           float32
	Reviews          []Review
	Url              string
	UserRatingsTotal int
	Website          string
}
