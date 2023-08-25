package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var (
		promPort = flag.Int("prom.port", 9150, "port to expose prometheus metrics")
	)
	flag.Parse()

	reg := prometheus.NewRegistry()
	// reg.MustRegister(collectors.NewGoCollector())

	mux := http.NewServeMux()
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	mux.Handle("/metrics", promHandler)

	port := fmt.Sprintf(":%d", *promPort)
	log.Printf("Exporter is started on %q/metrics", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Error, the Exporter can't be run : %s", err)
	}

}
