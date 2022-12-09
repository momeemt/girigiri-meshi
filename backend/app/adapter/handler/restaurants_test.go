package handler

import (
	"bytes"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/momeemt/2000s/domain/model"
	"github.com/momeemt/2000s/mock/mock_usecase"
	"github.com/momeemt/2000s/testutil"
)

func TestHandleRestaurants(t *testing.T) {
	tests := []struct {
		name      string
		postBody  string
		prepareRU func(m *mock_usecase.MockRestaurants)
	}{
		{
			name:     "success",
			postBody: `{"latitude":35.58276308282412,"longitude":140.1326089828801}`,
			prepareRU: func(m *mock_usecase.MockRestaurants) {
				m.EXPECT().GetAvailableRestaurants(model.Location{
					Latitude:  35.58276308282412,
					Longitude: 140.1326089828801,
				}, gomock.Any()).Return([]model.Restaurant{{
					Name: "鐵 蘇我本店",
					Location: model.Location{
						Latitude:  35.5827517,
						Longitude: 140.1327256,
					},
					CloseTime: time.Date(2022, 12, 7, 22, 00, 00, 00, time.Local),
					PhotoUrl:  "https://maps.googleapis.com/maps/api/place/photo?maxwidth=1000&photo_reference=AW30NDz7Dj5GkEsTnKGBetRfIK-GqhVNz_CAGyuOsX_UC8q_ZB_MLWr5VhOyjK-rzB91ZfNK3epikKXyiyguabpXOBWRypU4Nd18o9atheKz-Fv9OcQ6VYSXXqjJSeu8wJH0h-YWVMwjEHx2ezA7tvKv4tjRg9a8yLx3TKgJZEGhlx5-us4M&key=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
					PlaceId:   "ChIJeSTaRQWbImARznMsCGVeYcw",
					Rating:    3.7,
				}}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mRU := mock_usecase.NewMockRestaurants(ctrl)
			tt.prepareRU(mRU)
			req, err := http.NewRequest(http.MethodPost, "/restaurants", bytes.NewBufferString(tt.postBody))
			req.Header.Set("Content-Type", "application/json; charset=UTF-8")
			if err != nil {
				panic(err)
			}
			handler := NewRestaurantsHandler(mRU)
			err = testutil.TestRequest(handler.HandleRestaurants, req)
			if err != nil {
				t.Errorf("got error %v\n", err)
			}
		})
	}
}
