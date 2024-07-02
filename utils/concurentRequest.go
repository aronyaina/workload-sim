package utils

import (
	"log"
	"net/http"
	"sync"

	"github.com/aronyaina/workload-sim/models"
)

func HandleRequestsConcurrently(requests []models.Request) {
	var wg sync.WaitGroup

	for _, req := range requests {
		wg.Add(1)
		go func(r models.Request) {
			defer wg.Done()

			switch r.Method {
			case "GET":
				_, err := GetData(r)
				if err != nil {
					log.Printf("Error creating POST request to %s: %v", r.URL, err)
				}
			case "POST":
				var resp *http.Response
				var err error

				if len(r.Form) > 0 {
					_, err = PostFormData(r)
				} else if r.Body != nil {
					_, err = PostJSONBody(r)
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
		}(req)
	}
	wg.Wait()
}
