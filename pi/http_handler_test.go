package pi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalculatorHandlerReturnsPiValue(t *testing.T) {
	req, err := http.NewRequest("GET", "/pi", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CalculatorHandler)
	handler.ServeHTTP(recorder, req)

	status := recorder.Code
	if status != http.StatusOK {
		t.Errorf("hello handler did not return 200: got %d", status)
	}

	var resp piAproximationResponse
	_ = json.NewDecoder(recorder.Body).Decode(&resp)

	value := fmt.Sprintf("%f", resp.Pi)
	if !strings.HasPrefix(value, "3.") {
		// %v prints values only, %#v prints keys and values
		t.Errorf("hello handler didn't return expected Pi value: %#v\n", resp)
	}
}
