package handler

import (
	"log"
	"sync"
	"time"

	"github.com/aronyaina/workload-sim/models"
)

func HandleRequestsConcurrently(requests []models.Request, cancel <-chan struct{}, requestsPerSecond int, timeout time.Duration) {
	var wg sync.WaitGroup
	var requestCount int
	var mu sync.Mutex

	limiter := time.Tick(time.Second / time.Duration(requestsPerSecond))
	timeoutCh := time.After(timeout)

	incrementRequestCount := func() {
		mu.Lock()
		requestCount++
		mu.Unlock()
	}

	logRequestCount := func() {
		mu.Lock()
		log.Printf("Total requests executed: %d", requestCount)
		mu.Unlock()
	}

	for _, req := range requests {
		wg.Add(1)
		go func(r models.Request) {
			defer wg.Done()
			select {
			case <-cancel:
				return
			case <-timeoutCh:
				return
			case <-limiter:
				incrementRequestCount()
				ExecuteRequest(r, cancel)
			}
		}(req)
	}
	select {
	case <-timeoutCh:
		logRequestCount()
		log.Println("Request timeout reached")
	case <-cancel:
		logRequestCount()
		log.Println("Request all executed")
	}
}
