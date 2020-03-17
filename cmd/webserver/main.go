package main

import (
	"fmt"
	"log"
	"net/http"

	gomicroservice "github.com/cheesydev/go-microservice"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	v := gomicroservice.Version()
	fmt.Printf("Running go-microservice v%s, listening at port 8080\n", v)

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/pi", piHandler)
	http.Handle("/metrics", promhttp.Handler())
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
	pi := (4. / 1) - (4. / 3) + (4. / 5) - (4. / 7) + (4. / 9) - (4. / 11) + (4. / 13) - (4. / 15)

	fmt.Fprintf(w, "%v\n", pi)
}
