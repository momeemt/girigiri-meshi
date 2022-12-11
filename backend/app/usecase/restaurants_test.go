//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/mock_$GOPACKAGE/mock_$GOFILE
package usecase

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/momeemt/2000s/domain/model"
	"github.com/momeemt/2000s/mock/mock_repository"
)

func Test_restaurantsUsecase_GetAvailableRestaurants(t *testing.T) {
	// EXPECT まわりが冗長 名前だけマッチとかしたいね
	tests := []struct {
		name      string
		wantErr   bool
		location  model.Location
		time      time.Time
		prepareRR func(m *mock_repository.MockRestaurant)
	}{
		{
			name:     "success",
			wantErr:  false,
			location: model.Location{Latitude: 35.706028214316625, Longitude: 139.71668341868383},
			time:     time.Date(2022, 12, 8, 21, 0, 0, 0, time.Local),
			prepareRR: func(m *mock_repository.MockRestaurant) {
				m.EXPECT().GetNearbyRestaurants(model.Location{
					Latitude:  35.706028214316625,
					Longitude: 139.71668341868383,
				}, gomock.Any(), gomock.Any()).Return([]model.Restaurant{
					{
						Name: "Saizeriya Nishi-Waseda",
						Location: model.Location{
							Latitude:  35.70670800000001,
							Longitude: 139.716813,
						},
						PhotoUrl: "https://lh3.googleusercontent.com/places/AJDFj40BrNvrjCwyFnyutzhWTBod6mrELN7sABx13W1H2pFs0ImIqrrtBaI5n_LQBMRvqdkybqkj8qCdtNwiCbfNsgAkUbkvs5luiMA=s1600-w1000",
						PlaceId:  "ChIJu4I8HRyNGGAR0Zb6c2hWAoc",
						Rating:   3.5,
					},
					{
						Name: "ワンプラスダンプリング 西早稲田",
						Location: model.Location{
							Latitude:  35.7065364,
							Longitude: 139.7174725,
						},
						PhotoUrl: "https://lh3.googleusercontent.com/places/AJDFj40BrNvrjCwyFnyutzhWTBod6mrELN7sABx13W1H2pFs0ImIqrrtBaI5n_LQBMRvqdkybqkj8qCdtNwiCbfNsgAkUbkvs5luiMA=s1600-w1000",
						PlaceId:  "ChIJIdqmb92NGGAR3SetKe6rY2Q",
						Rating:   4.1,
					}, {
						Name: "Restaurant & Bar New school",
						Location: model.Location{
							Latitude:  35.8065153,
							Longitude: 139.7178228,
						},
						PhotoUrl: "https://lh3.googleusercontent.com/places/AJDFj40BrNvrjCwyFnyutzhWTBod6mrELN7sABx13W1H2pFs0ImIqrrtBaI5n_LQBMRvqdkybqkj8qCdtNwiCbfNsgAkUbkvs5luiMA=s1600-w1000",
						PlaceId:  "ChIJAWs2PxyNGGARySvrTT7hkPw",
						Rating:   4.3,
					},
				}, nil)
				m.EXPECT().GetNextCloseTime(model.Restaurant{
					Name: "Saizeriya Nishi-Waseda",
					Location: model.Location{
						Latitude:  35.70670800000001,
						Longitude: 139.716813,
					},
					PhotoUrl: "https://lh3.googleusercontent.com/places/AJDFj40BrNvrjCwyFnyutzhWTBod6mrELN7sABx13W1H2pFs0ImIqrrtBaI5n_LQBMRvqdkybqkj8qCdtNwiCbfNsgAkUbkvs5luiMA=s1600-w1000",
					PlaceId:  "ChIJu4I8HRyNGGAR0Zb6c2hWAoc",
					Rating:   3.5,
				}, gomock.Any()).Return(time.Date(2022, 12, 8, 18, 00, 0, 0, time.Local), nil)
				m.EXPECT().GetNextCloseTime(model.Restaurant{
					Name: "ワンプラスダンプリング 西早稲田",
					Location: model.Location{
						Latitude:  35.7065364,
						Longitude: 139.7174725,
					},
					PhotoUrl: "https://lh3.googleusercontent.com/places/AJDFj40BrNvrjCwyFnyutzhWTBod6mrELN7sABx13W1H2pFs0ImIqrrtBaI5n_LQBMRvqdkybqkj8qCdtNwiCbfNsgAkUbkvs5luiMA=s1600-w1000",
					PlaceId:  "ChIJIdqmb92NGGAR3SetKe6rY2Q",
					Rating:   4.1,
				}, gomock.Any()).Return(time.Date(2022, 12, 8, 22, 00, 00, 00, time.Local), nil)
				m.EXPECT().GetNextCloseTime(model.Restaurant{
					Name: "Restaurant & Bar New school",
					Location: model.Location{
						Latitude:  35.8065153,
						Longitude: 139.7178228,
					},
					PhotoUrl: "https://lh3.googleusercontent.com/places/AJDFj40BrNvrjCwyFnyutzhWTBod6mrELN7sABx13W1H2pFs0ImIqrrtBaI5n_LQBMRvqdkybqkj8qCdtNwiCbfNsgAkUbkvs5luiMA=s1600-w1000",
					PlaceId:  "ChIJAWs2PxyNGGARySvrTT7hkPw",
					Rating:   4.3,
				}, gomock.Any()).Return(time.Date(2022, 12, 8, 22, 30, 0, 0, time.Local), nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mRR := mock_repository.NewMockRestaurant(ctrl)
			tt.prepareRR(mRR)
			r := &restaurantsUsecase{
				restaurantRepository: mRR,
			}
			got, err := r.GetAvailableRestaurants(tt.location, tt.time)
			fmt.Printf("%+v\n", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("restaurantsUsecase.GetAvailableRestaurants() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
