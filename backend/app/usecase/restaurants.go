package usecase

import "github.com/momeemt/2000s/domain/model"

type Restaurants interface {
	// GetAvailableRestaurants は距離と閉店時刻を計算に入れて行くことができる飲食店を近い順番で並べて返す
	GetAvailableRestaurants(model.Location) ([]model.Restaurant, error)
}
