package utils

import (
	"log"
	"net/http"
	"time"

	"github.com/aronyaina/workload-sim/models"
)

func GetData(r models.Request) (*http.Response, error) {

	startTime := time.Now()
	resp, err := http.Get(r.URL)
	if err != nil {
		log.Printf("Error performing GET request to %s: %v", r.URL, err)
	}
	duration := time.Since(startTime)
	log.Println("GET Time taken:", r.URL, duration)

	return resp, err
}
