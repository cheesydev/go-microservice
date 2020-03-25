package pi

import (
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

	body := recorder.Body.String()
	if !strings.HasPrefix(body, "3.") {
		t.Errorf("hello handler didn't return expected Pi value: %s\n", body)
	}
}
