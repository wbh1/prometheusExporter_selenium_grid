package main

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"time"
)

type hubResponse struct {
	Slots        slotCounts `json:"slotCounts"`
	NewSession   float64    `json:"newSessionRequestCount"`
}

type slotCounts struct {
	Free  float64 `json:"free"`
	Total float64 `json:"total"`
}


func calcMetrics() (float64, float64, float64) {
	url := "http://selenium.liberty.edu:4444/grid/api/hub"	

	interwebs := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := interwebs.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonOutput := hubResponse{}
	jsonErr := json.Unmarshal(body, &jsonOutput)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}


	x := float64(jsonOutput.NewSession)
	y := float64(jsonOutput.Slots.Free)
	z := float64(jsonOutput.Slots.Total)
	return x, y, z
}