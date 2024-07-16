package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/aronyaina/workload-sim/models"
	"github.com/aronyaina/workload-sim/utils"
)

func GetData(r models.Request, cancel <-chan struct{}, ctx context.Context) (*http.Response, error) {
	startTime := time.Now()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.URL, nil)
	if err != nil {
		return nil, utils.HandleContextError(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error performing GET request to %s: %v", r.URL, err)
	}
	duration := time.Since(startTime)
	log.Println("GET "+r.URL+" Time taken:", duration)

	if err := CheckCancelation(cancel); err != nil {
		return nil, err
	}
	return resp, nil
}
