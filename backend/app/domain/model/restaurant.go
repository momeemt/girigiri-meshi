package model

import "time"

type Restaurant struct {
	Name      string
	Location  Location
	CloseTime time.Time
	PhotoUrl  string
	PlaceId   string //dependent to Google Maps API
	Rating    float32
}
