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

	// instrument (wrap) handler with a duration observer
	http.HandleFunc("/pi", promhttp.InstrumentHandlerDuration(
		promauto.NewSummaryVec(
			prometheus.SummaryOpts{
				Name:       "pi_requests_durations_seconds",
				Help:       "Duration distributions of pi http requests",
				Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
			},
			// partition the counter by http status code
			// (code and method are the supported partition labels)
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

	var p pi.PiCalculator = pi.Leibniz{}

	fmt.Fprintf(w, "%v\n", p.CalculatePi())
}
