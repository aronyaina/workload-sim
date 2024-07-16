package handler

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/aronyaina/workload-sim/models"
)

func HandleRequestsConcurrently(requests []models.Request, timeout time.Duration, requestsPerSecond int, cancel <-chan struct{}) {
	var wg sync.WaitGroup
	var requestCount int
	var mu sync.Mutex

	limiter := time.Tick(time.Second / time.Duration(requestsPerSecond))
	done := make(chan struct{})

	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	defer cancelFunc()

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
		select {
		case <-cancel:
			log.Println("Request cancellation requested")
			return
		case <-ctx.Done():
			logRequestCount()
			log.Println("Context timeout reached")
			return
		default:
			wg.Add(1)
			go func(r models.Request) {
				defer wg.Done()

				select {
				case <-cancel:
					return
				case <-ctx.Done():
					return
				default:
					select {
					case <-cancel:
						return
					case <-ctx.Done():
						return
					case <-limiter:
						incrementRequestCount()
						ExecuteRequest(r, cancel, timeout, ctx)
					}
				}
			}(req)
		}
	}
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-ctx.Done():
		logRequestCount()
		log.Println("Context timeout reached")
	case <-cancel:
		logRequestCount()
		log.Println("Request cancellation requested")
	case <-done:
		logRequestCount()
		log.Println("All requests executed")
	}
}
