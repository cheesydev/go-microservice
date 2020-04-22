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

  // TODO: get param from querystring
  id := "10"

  // TODO: check err
  rating, _ := ratingByProductId(id)

	resp := ratingResponse{
		ProductId: id,
		Rating:    rating,
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
