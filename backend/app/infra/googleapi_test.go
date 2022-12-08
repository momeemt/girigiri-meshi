package infra

import (
	"fmt"
	"testing"
	"time"

	"github.com/momeemt/2000s/domain/model"
)

func Test_googlePlacesApi_GetNearbyRestaurants(t *testing.T) {
	tests := []struct {
		name     string
		g        googlePlacesApi
		want     []model.Restaurant
		wantErr  bool
		location model.Location
	}{
		{
			name:    "soga",
			g:       googlePlacesApi{},
			want:    []model.Restaurant{},
			wantErr: false,
			location: model.Location{
				Latitude:   35.5827517,
				Longtitude: 140.1327256,
			},
		},
		{
			name:     "waseda",
			g:        googlePlacesApi{},
			want:     []model.Restaurant{},
			wantErr:  false,
			location: model.Location{Latitude: 35.706028214316625, Longtitude: 139.71668341868383},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := googlePlacesApi{}
			got, err := g.GetNearbyRestaurants(tt.location)
			fmt.Printf("%+v\n", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("googlePlacesApi.GetNearbyRestaurants() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_googlePlacesApi_GetCloseTime(t *testing.T) {
	tests := []struct {
		name       string
		g          googlePlacesApi
		restaurant model.Restaurant
		time       time.Time
		wantErr    bool
	}{
		{
			name: "鐡",
			g:    googlePlacesApi{},
			restaurant: model.Restaurant{
				PlaceId: "ChIJeSTaRQWbImARznMsCGVeYcw",
			},
			time:    time.Now(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := googlePlacesApi{}
			got, err := g.GetNextCloseTime(tt.restaurant, tt.time)
			fmt.Printf("%+v\n", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("googlePlacesApi.GetNearbyRestaurants() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
