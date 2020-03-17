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
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go-microservice v%s\n", gomicroservice.Version())
}
