package handler

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/momeemt/2000s/testutil"
)

func TestHandleRestaurants(t *testing.T) {
	tests := []struct {
		name     string
		postBody string
	}{
		{
			name:     "success",
			postBody: `{"latitude":35.58276308282412,"longtitude":140.1326089828801}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodPost, "/restaurants", bytes.NewBufferString(tt.postBody))
			req.Header.Set("Content-Type", "application/json; charset=UTF-8")
			if err != nil {
				panic(err)
			}
			err = testutil.TestRequest(HandleRestaurants, req)
			if err != nil {
				t.Errorf("got error %v\n", err)
			}
		})
	}
}
