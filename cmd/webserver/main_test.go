package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cheesydev/go-microservice/pi"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)
	handler.ServeHTTP(recorder, req)

	status := recorder.Code
	if status != http.StatusOK {
		t.Errorf("hello handler did not return 200: got %d", status)
	}

	body := recorder.Body.String()
	expected := "go-microservice"
	if !strings.Contains(body, expected) {
		t.Errorf("hello handler didn't return expected message: %s",
			expected)
	}
}

func TestNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/unknown", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)
	handler.ServeHTTP(recorder, req)

	status := recorder.Code
	if status != http.StatusNotFound {
		t.Errorf("hello handler did not return 404 as expected: got %d", status)
	}

	body := recorder.Body.String()
	expected := "oops not found"
	if !strings.Contains(body, expected) {
		t.Errorf("hello handler didn't return expected not found message: %s",
			expected)
	}
}

func TestPiValueIsThreePointSomething(t *testing.T) {
	req, err := http.NewRequest("GET", "/pi", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(pi.CalculatorHandler)
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
