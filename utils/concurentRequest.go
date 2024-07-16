package utils

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/aronyaina/workload-sim/models"
)

var RequestCount int

func HandleRequestsConcurrently(requests []models.Request, cancel <-chan struct{}, requestsPerSecond int) {
	var wg sync.WaitGroup

	limiter := time.Tick(time.Second / time.Duration(requestsPerSecond))

	for i, req := range requests {
		wg.Add(1)
		go func(r models.Request) {
			defer wg.Done()
			select {
			case <-cancel:
				return
			case <-limiter:
				switch r.Method {
				case "GET":
					_, err := GetData(r, cancel)
					if err != nil {
						log.Printf("Error creating GET request to %s: %v", r.URL, err)
					}
				case "POST":
					var resp *http.Response
					var err error

					if len(r.Form) > 0 {
						_, err = PostFormData(r, cancel)
					} else if r.Body != nil {
						_, err = PostJSONBody(r, cancel)
					} else {
						_ = resp
						log.Fatalln("Error parsing body ...")
					}

					if err != nil {
						log.Printf("Error creating POST request to %s: %v", r.URL, err)
						return
					}
				case "PUT":
					var resp *http.Response
					var err error

					if len(r.Form) > 0 {
						_, err = PutFormData(r, cancel)
					} else if r.Body != nil {
						_, err = PutJSONBody(r, cancel)
					} else {
						_ = resp
						log.Fatalln("Error parsing body ...")
					}

					if err != nil {
						log.Printf("Error creating PUT request to %s: %v", r.URL, err)
						return
					}
				case "DELETE":
					_, err := DeleteData(r, cancel)
					if err != nil {
						log.Printf("Error creating DELETE request to %s: %v", r.URL, err)
					}

				default:
					log.Printf("Unsupported method: %s for URL: %s", r.Method, r.URL)
				}
			}
		}(req)
		RequestCount = i + 1
	}
	wg.Wait()
}
