package handler

import "net/http"

func HandleRestaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`[
	  {
	"name": "鐵 蘇我本店",
	"location": {
	  "latitude": 35.5827517,
	  "longtitude": 140.1327256
	},
	"closeTime": "2022-12-07T22:00:00+09:00",
	"photoUrl": "https://maps.googleapis.com/maps/api/place/photo?maxwidth=1000&photo_reference=AW30NDz7Dj5GkEsTnKGBetRfIK-GqhVNz_CAGyuOsX_UC8q_ZB_MLWr5VhOyjK-rzB91ZfNK3epikKXyiyguabpXOBWRypU4Nd18o9atheKz-Fv9OcQ6VYSXXqjJSeu8wJH0h-YWVMwjEHx2ezA7tvKv4tjRg9a8yLx3TKgJZEGhlx5-us4M&key=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"place_id": "ChIJeSTaRQWbImARznMsCGVeYcw",
	"rating": 3.7
	  }
	]`))
}
