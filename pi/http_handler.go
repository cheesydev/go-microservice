package pi

import (
	"fmt"
	"net/http"
)

// Handler function to serve aproximations of the  pi value.
func CalculatorHandler(w http.ResponseWriter, r *http.Request) {
	var p PiCalculator = Leibniz{}
	fmt.Fprintf(w, "%v\n", p.CalculatePi())
}
