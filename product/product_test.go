package product

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProducRatingHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/rating", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(RatingHandler)
	handler.ServeHTTP(recorder, req)

	var resp ratingResponse
	_ = json.NewDecoder(recorder.Body).Decode(&resp)

	if resp.Rating != "4" {
		t.Errorf("rating handler returned wrong value.\n")
	}
}
