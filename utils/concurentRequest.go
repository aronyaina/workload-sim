package utils

import (
	"log"
	"net/http"
	"sync"
)

func HandleRequestsConcurrently(requests []Request) {
	var wg sync.WaitGroup

	for _, req := range requests {
		wg.Add(1)
		go func(r Request) {
			defer wg.Done()
			err := performRequest(r)
			if err != nil {
				log.Printf("Error performing request to %s: %v", r.URL, err)
			}
		}(req)
	}
	wg.Wait()
}

func performRequest(r Request) error {
	var req *http.Request
	var err error

	return nil
}
