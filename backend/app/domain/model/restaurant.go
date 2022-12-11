package model

import "time"

type Restaurant struct {
	Name                 string
	Location             Location
	CloseTime            time.Time
	PhotoUrl             string
	PhotoUrls            []string
	PlaceId              string //dependent to Google Maps API
	Rating               float32
	Reviews              []Review
	ServesBeer           bool
	ServesBranch         bool
	ServesBreakfast      bool
	ServesDinner         bool
	ServesLunch          bool
	ServesVegetarianFood bool
	ServesWine           bool
	Takeout              bool
	Url                  string
	UserRatingsTotal     int
	Website              string
}
