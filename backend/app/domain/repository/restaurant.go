package repository

import (
	"time"

	"github.com/momeemt/2000s/domain/model"
)

type Restaurant interface {
	GetNearbyRestaurants(model.Location) ([]model.Restaurant, error)
	GetCloseTime(model.Restaurant) (time.Time, error)
}
