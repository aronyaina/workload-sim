package handler

import (
	"context"
	"log"
	"time"

	"github.com/aronyaina/workload-sim/models"
	"github.com/aronyaina/workload-sim/utils"
)

func ExecuteRequest(r models.Request, cancel <-chan struct{}, timeout time.Duration, ctx context.Context) {

	switch r.Method {
	case "GET":
		_, err := GetData(r, cancel, ctx)
		if err != nil {
			log.Printf("Error executing GET request to %s: %v", r.URL, err)
		}
	case "POST":
		err := handlePostRequest(r, cancel, ctx)
		if err != nil {
			log.Printf("Error executing POST request to %s: %v", r.URL, err)
		}
	case "PUT":
		err := handlePutRequest(r, cancel, ctx)
		if err != nil {
			log.Printf("Error executing PUT request to %s: %v", r.URL, err)
		}
	case "DELETE":
		_, err := DeleteData(r, cancel, ctx)
		if err != nil {
			log.Printf("Error executing DELETE request to %s: %v", r.URL, err)
		}
	default:
		log.Printf("Unsupported method: %s for URL: %s", r.Method, r.URL)
	}
}

func handlePostRequest(r models.Request, cancel <-chan struct{}, ctx context.Context) error {
	if len(r.Form) > 0 {
		_, err := PostFormData(r, cancel, ctx)
		if err != nil {
			return utils.HandleContextError(err)
		}
		return utils.HandleContextError(err)
	}
	if r.Body != nil {
		_, err := PostJSONBody(r, cancel, ctx)
		if err != nil {
			return utils.HandleContextError(err)
		}
	}
	log.Fatalln("Error parsing POST body ...")
	return nil
}

func handlePutRequest(r models.Request, cancel <-chan struct{}, ctx context.Context) error {
	if len(r.Form) > 0 {
		_, err := PutFormData(r, cancel, ctx)
		if err != nil {
			return utils.HandleContextError(err)
		}
	}
	if r.Body != nil {
		_, err := PutJSONBody(r, cancel, ctx)
		if err != nil {
			return utils.HandleContextError(err)
		}
	}
	log.Fatalln("Error parsing PUT body ...")
	return nil
}
