//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/mock_$GOPACKAGE/mock_$GOFILE
package repository

import (
	"time"

	"github.com/momeemt/2000s/domain/model"
)

type Restaurant interface {
	GetNearbyRestaurants(model.Location, time.Time, bool) ([]model.Restaurant, error)
	GetNextCloseTime(model.Restaurant, time.Time) (time.Time, error)
	GetRestaurantDetail(placeId string) (model.Restaurant, error)
}
