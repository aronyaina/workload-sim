package handler

import (
	"log"

	"github.com/aronyaina/workload-sim/models"
)

func ExecuteRequest(r models.Request, cancel <-chan struct{}) {
	var err error

	switch r.Method {
	case "GET":
		_, err = GetData(r, cancel)
	case "POST":
		err = handlePostRequest(r, cancel)
	case "PUT":
		err = handlePutRequest(r, cancel)
	case "DELETE":
		_, err = DeleteData(r, cancel)
	default:
		log.Printf("Unsupported method: %s for URL: %s", r.Method, r.URL)
		return
	}

	if err != nil {
		log.Printf("Error executing %s request to %s: %v", r.Method, r.URL, err)
	}
}

func handlePostRequest(r models.Request, cancel <-chan struct{}) error {
	if len(r.Form) > 0 {
		_, err := PostFormData(r, cancel)
		return err
	}
	if r.Body != nil {
		_, err := PostJSONBody(r, cancel)
		return err
	}
	log.Fatalln("Error parsing POST body ...")
	return nil
}

func handlePutRequest(r models.Request, cancel <-chan struct{}) error {
	if len(r.Form) > 0 {
		_, err := PutFormData(r, cancel)
		return err
	}
	if r.Body != nil {
		_, err := PutJSONBody(r, cancel)
		return err
	}
	log.Fatalln("Error parsing PUT body ...")
	return nil
}
