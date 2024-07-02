package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aronyaina/workload-sim/models"
)

func PostFormData(r models.Request, cancel <-chan struct{}) (*http.Response, error) {
	formData := url.Values{}
	for key, value := range r.Form {
		formData.Set(key, value)
	}

	startTime := time.Now()
	resp, err := http.Post(r.URL, "application/x-www-form-urlencoded", strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}

	duration := time.Since(startTime)
	log.Println("POST "+r.URL+" Time taken:", duration)

	if err := CheckCancelation(cancel); err != nil {
		return nil, err
	}
	return resp, nil
}

func PostJSONBody(r models.Request, cancel <-chan struct{}) (*http.Response, error) {
	bodyBytes, err := json.Marshal(r.Body)
	//log.Println("JSON body:", string(bodyBytes))
	if err != nil {
		log.Printf("Error marshalling body for %s: %v", r.URL, err)
		return nil, err
	}

	startTime := time.Now()
	resp, err := http.Post(r.URL, "application/json", bytes.NewReader(bodyBytes))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	duration := time.Since(startTime)
	log.Println("POST "+r.URL+" Time taken:", duration)

	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusNotFound {
		log.Printf("404 Not Found: %s", string(respBody))
		return nil, errors.New("404 Not Found: " + string(respBody))
	} else if resp.StatusCode == http.StatusBadRequest {
		log.Printf("400 Bad Request: %s", string(respBody))
		return nil, fmt.Errorf("400 Bad Request: %s", string(respBody))
	} else if resp.StatusCode >= 400 {
		log.Printf("HTTP %d: %s", resp.StatusCode, string(respBody))
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(respBody))
	}

	if err := CheckCancelation(cancel); err != nil {
		return nil, err
	}

	return resp, nil
}
