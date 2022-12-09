package infra

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/momeemt/2000s/domain/model"
)

func Test_googlePlacesApi_GetNearbyRestaurants(t *testing.T) {
	tests := []struct {
		name     string
		wantErr  bool
		location model.Location
	}{
		{
			name:    "soga",
			wantErr: false,
			location: model.Location{
				Latitude:   35.5827517,
				Longtitude: 140.1327256,
			},
		},
		{
			name:     "waseda",
			wantErr:  false,
			location: model.Location{Latitude: 35.706028214316625, Longtitude: 139.71668341868383},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := googlePlacesApi{
				apiKey: os.Getenv("GOOGLE_PLACES_API_KEY"),
			}
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
		restaurant model.Restaurant
		time       time.Time
		wantErr    bool
	}{
		{
			name: "鐡",
			restaurant: model.Restaurant{
				PlaceId: "ChIJeSTaRQWbImARznMsCGVeYcw",
			},
			time:    time.Now(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := googlePlacesApi{
				apiKey: os.Getenv("GOOGLE_PLACES_API_KEY"),
			}
			got, err := g.GetNextCloseTime(tt.restaurant, tt.time)
			fmt.Printf("%+v\n", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("googlePlacesApi.GetNearbyRestaurants() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_getUrl(t *testing.T) {
	type args struct {
		photoReference string
		apiKey         string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				photoReference: "AW30NDz7Dj5GkEsTnKGBetRfIK-GqhVNz_CAGyuOsX_UC8q_ZB_MLWr5VhOyjK-rzB91ZfNK3epikKXyiyguabpXOBWRypU4Nd18o9atheKz-Fv9OcQ6VYSXXqjJSeu8wJH0h-YWVMwjEHx2ezA7tvKv4tjRg9a8yLx3TKgJZEGhlx5-us4M",
				apiKey:         os.Getenv("GOOGLE_PLACES_API_KEY"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getUrl(tt.args.photoReference, tt.args.apiKey)
			fmt.Println(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
