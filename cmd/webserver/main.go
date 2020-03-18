package main

import (
	"fmt"
	"log"
	"net/http"

	gomicroservice "github.com/cheesydev/go-microservice"
	"github.com/cheesydev/go-microservice/pi"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	v := gomicroservice.Version()

	http.HandleFunc("/", helloHandler)

	// example from https://github.com/prometheus/client_golang/blob/master/prometheus/promauto/auto.go
	http.HandleFunc("/pi", promhttp.InstrumentHandlerCounter(
		promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "pi_requests_total",
				Help: "Total number of pi http requests",
			},
			// partition the counter by http status code (a supported label)
			[]string{"code"},
		),
		http.HandlerFunc(piHandler),
	))

	http.Handle("/metrics", promhttp.Handler())

	fmt.Printf("Running go-microservice v%s, listening at port 8080\n", v)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		w.Write([]byte("oops not found"))
		return
	}
	fmt.Fprintf(w, "go-microservice v%s\n", gomicroservice.Version())
}

func piHandler(w http.ResponseWriter, r *http.Request) {

	// Leibniz formula
	pi := pi.LeibnizPi()

	fmt.Fprintf(w, "%v\n", pi)
}
