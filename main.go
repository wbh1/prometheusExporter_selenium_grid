package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus"

)


func main() {

	selCollector := newSeleniumCollector()
	prometheus.MustRegister(selCollector)
	
	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	print("Listening on port 8080... \n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
