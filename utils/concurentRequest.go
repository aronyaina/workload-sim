package utils

import (
	"log"
	"net/http"
	"sync"

	"github.com/aronyaina/workload-sim/models"
)

func HandleRequestsConcurrently(requests []models.Request, cancel <-chan struct{}) {
	var wg sync.WaitGroup

	for _, req := range requests {
		wg.Add(1)
		go func(r models.Request) {
			defer wg.Done()
			select {
			case <-cancel:
				return
			default:
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

				default:
					log.Printf("Unsupported method: %s for URL: %s", r.Method, r.URL)
				}
			}
		}(req)
	}
	wg.Wait()
}
