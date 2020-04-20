package product

import (
	"encoding/json"
	"net/http"
)

type ratingResponse struct {
	ProductId string `json:product_id`
	Rating    string `json:rating`
}

func RatingHandler(w http.ResponseWriter, r *http.Request) {

	resp := ratingResponse{
		ProductId: "10",
		Rating:    "4",
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
