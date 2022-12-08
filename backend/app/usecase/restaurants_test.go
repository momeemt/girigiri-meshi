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
			location: model.Location{Latitude: 35.706028214316625, Longtitude: 139.71668341868383},
			time:     time.Date(2022, 12, 8, 21, 0, 0, 0, time.Local),
			prepareRR: func(m *mock_repository.MockRestaurant) {
				m.EXPECT().GetNearbyRestaurants(model.Location{
					Latitude:   35.706028214316625,
					Longtitude: 139.71668341868383,
				}).Return([]model.Restaurant{
					{
						Name: "Saizeriya Nishi-Waseda",
						Location: model.Location{
							Latitude:   35.70670800000001,
							Longtitude: 139.716813,
						},
						PhotoUrl: "https://maps.googleapis.com/maps/api/place/photo?maxwidth=1000&photo_reference=AW30NDyYirIzcXuADvUgJLalN2eTscRoAM6MZ6alsexTrg5glpLCYh9hPeH5CCowWM3K3fZeI-AypafFJBG1fkkpXzeMew8G2uJabkIKKQHKhhcQHip4YqHQG14PGDe0gpgEMiN_xQYjTuIEvdRaqIPL5Koh4ijNcB4SNHpXKI0P6Lh6EVZ5&key=AIzaSyD73WTJXHUol9u8BsgINXK0DkdfqiQurd8",
						PlaceId:  "ChIJu4I8HRyNGGAR0Zb6c2hWAoc",
						Rating:   3.5,
					},
					{
						Name: "ワンプラスダンプリング 西早稲田",
						Location: model.Location{
							Latitude:   35.7065364,
							Longtitude: 139.7174725,
						},
						PhotoUrl: "https://maps.googleapis.com/maps/api/place/photo?maxwidth=1000&photo_reference=AW30NDzOxFqJ8oQZGO4KRTSDU0j9hxJ3Df6dLAuV-JTd_F2aB4XCpIZf9yXoa8UzgZ6JPH-7Ra1Eel3kP5OS0tszYavGlJ0cKONSIlUBdsww7GsrKHfPCIXYjbvukftv72Xy-L1tBUUDZrNmctDEG_aIYpncGR647mVyGonhXAh_0cxHioaX&key=AIzaSyD73WTJXHUol9u8BsgINXK0DkdfqiQurd8",
						PlaceId:  "ChIJIdqmb92NGGAR3SetKe6rY2Q",
						Rating:   4.1,
					}, {
						Name: "Restaurant & Bar New school",
						Location: model.Location{
							Latitude:   35.8065153,
							Longtitude: 139.7178228,
						},
						PhotoUrl: "https://maps.googleapis.com/maps/api/place/photo?maxwidth=1000&photo_reference=AW30NDzjx3XypgDP_nPHHlfn4ALOPUXpSUVM0YYGSE_nQO54IcqAtf7GSZDxWOiWYnM1VbweJiVyNk-8XTxMK9Myf-bNTB4a-oRLNTl21JxY3-zd0dwPXlMftO6yD56jVhs6WRJxfcHQFM_Dw-yBe5WvIGNO58XcqO56MlQtvMyitpBIBkt-&key=AIzaSyD73WTJXHUol9u8BsgINXK0DkdfqiQurd8",
						PlaceId:  "ChIJAWs2PxyNGGARySvrTT7hkPw",
						Rating:   4.3,
					},
				}, nil)
				m.EXPECT().GetNextCloseTime(model.Restaurant{
					Name: "Saizeriya Nishi-Waseda",
					Location: model.Location{
						Latitude:   35.70670800000001,
						Longtitude: 139.716813,
					},
					PhotoUrl: "https://maps.googleapis.com/maps/api/place/photo?maxwidth=1000&photo_reference=AW30NDyYirIzcXuADvUgJLalN2eTscRoAM6MZ6alsexTrg5glpLCYh9hPeH5CCowWM3K3fZeI-AypafFJBG1fkkpXzeMew8G2uJabkIKKQHKhhcQHip4YqHQG14PGDe0gpgEMiN_xQYjTuIEvdRaqIPL5Koh4ijNcB4SNHpXKI0P6Lh6EVZ5&key=AIzaSyD73WTJXHUol9u8BsgINXK0DkdfqiQurd8",
					PlaceId:  "ChIJu4I8HRyNGGAR0Zb6c2hWAoc",
					Rating:   3.5,
				}, gomock.Any()).Return(time.Date(2022, 12, 8, 18, 00, 0, 0, time.Local), nil)
				m.EXPECT().GetNextCloseTime(model.Restaurant{
					Name: "ワンプラスダンプリング 西早稲田",
					Location: model.Location{
						Latitude:   35.7065364,
						Longtitude: 139.7174725,
					},
					PhotoUrl: "https://maps.googleapis.com/maps/api/place/photo?maxwidth=1000&photo_reference=AW30NDzOxFqJ8oQZGO4KRTSDU0j9hxJ3Df6dLAuV-JTd_F2aB4XCpIZf9yXoa8UzgZ6JPH-7Ra1Eel3kP5OS0tszYavGlJ0cKONSIlUBdsww7GsrKHfPCIXYjbvukftv72Xy-L1tBUUDZrNmctDEG_aIYpncGR647mVyGonhXAh_0cxHioaX&key=AIzaSyD73WTJXHUol9u8BsgINXK0DkdfqiQurd8",
					PlaceId:  "ChIJIdqmb92NGGAR3SetKe6rY2Q",
					Rating:   4.1,
				}, gomock.Any()).Return(time.Date(2022, 12, 8, 22, 00, 00, 00, time.Local), nil)
				m.EXPECT().GetNextCloseTime(model.Restaurant{
					Name: "Restaurant & Bar New school",
					Location: model.Location{
						Latitude:   35.8065153,
						Longtitude: 139.7178228,
					},
					PhotoUrl: "https://maps.googleapis.com/maps/api/place/photo?maxwidth=1000&photo_reference=AW30NDzjx3XypgDP_nPHHlfn4ALOPUXpSUVM0YYGSE_nQO54IcqAtf7GSZDxWOiWYnM1VbweJiVyNk-8XTxMK9Myf-bNTB4a-oRLNTl21JxY3-zd0dwPXlMftO6yD56jVhs6WRJxfcHQFM_Dw-yBe5WvIGNO58XcqO56MlQtvMyitpBIBkt-&key=AIzaSyD73WTJXHUol9u8BsgINXK0DkdfqiQurd8",
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
