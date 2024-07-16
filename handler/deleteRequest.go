package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/aronyaina/workload-sim/models"
)

func DeleteData(r models.Request, cancel <-chan struct{}, ctx context.Context) (*http.Response, error) {
	startTime := time.Now()

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, r.URL, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	duration := time.Since(startTime)
	log.Println("DELETE "+r.URL+" Time taken:", duration)

	if err := CheckCancelation(cancel); err != nil {
		return nil, err
	}
	return resp, nil
}
