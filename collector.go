package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Define a structs for the collector that contains pointers
// to the Prometheus descriptors for each metric

type seleniumCollector struct {
	session_requests *prometheus.Desc
	slots_free *prometheus.Desc
	slots_total *prometheus.Desc
}

// Must create a constructor for collector to initialize
// each descriptor and return a pointer to the collector
func newSeleniumCollector() *seleniumCollector {
	return &seleniumCollector{
		session_requests: prometheus.NewDesc("selenium_session_requests",
						  "Number of Selenium requests waiting for a session",
						  nil, nil,
						),
		slots_free: prometheus.NewDesc("selenium_slots_free",
						  "Number of Selenium requests currently occupying a slot",
						  nil, nil,
						),
		slots_total: prometheus.NewDesc("selenium_slots_total",
						  "Number of Selenium slots total",
						  nil, nil,
						),
	}
}

// All collectors must implement the Describe function
// Writes all descriptors to the prometheus desc channel
func (collector *seleniumCollector) Describe(ch chan<- *prometheus.Desc) {
	// Should be updated with each metric created
	ch <- collector.session_requests
	ch <- collector.slots_free
	ch <- collector.slots_total
}

// Collect implements the required collect function
func (collector *seleniumCollector) Collect(ch chan<- prometheus.Metric) {
	// This is where the logic goes in order to set the metric values
	val_sessions, val_slots_used, val_total_slots := calcMetrics()

	ch <- prometheus.MustNewConstMetric(collector.session_requests, prometheus.GaugeValue, val_sessions)
	ch <- prometheus.MustNewConstMetric(collector.slots_free, prometheus.GaugeValue, val_slots_used)
	ch <- prometheus.MustNewConstMetric(collector.slots_total, prometheus.GaugeValue, val_total_slots)
}