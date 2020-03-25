package pi

import (
	"encoding/json"
	"net/http"
)

type piAproximationResponse struct {

	// fields should be exported (uppercase) to be json encoded
	Pi     float64 `json:pi`
	Method string  `json:method`
}

// Handler function to serve aproximations of the  pi value.
func CalculatorHandler(w http.ResponseWriter, r *http.Request) {

	var p PiCalculator = Leibniz{}
	var value float64 = p.CalculatePi()

	resp := piAproximationResponse{
		Pi:     value,
		Method: "leibniz",
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
