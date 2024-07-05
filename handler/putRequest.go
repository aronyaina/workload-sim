package handler

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

func PutFormData(r models.Request, cancel <-chan struct{}) (*http.Response, error) {
	formData := url.Values{}
	for key, value := range r.Form {
		formData.Set(key, value)
	}

	startTime := time.Now()

	req, err := http.NewRequest(http.MethodPut, r.URL, strings.NewReader(formData.Encode()))
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	duration := time.Since(startTime)
	log.Println("PUT "+r.URL+" Time taken:", duration)

	if err := CheckCancelation(cancel); err != nil {
		return nil, err
	}
	return resp, nil
}

func PutJSONBody(r models.Request, cancel <-chan struct{}) (*http.Response, error) {
	bodyBytes, err := json.Marshal(r.Body)
	//log.Println("JSON body:", string(bodyBytes))
	if err != nil {
		log.Printf("Error marshalling body for %s: %v", r.URL, err)
		return nil, err
	}

	startTime := time.Now()
	req, err := http.NewRequest(http.MethodPut, r.URL, bytes.NewReader(bodyBytes))
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	duration := time.Since(startTime)
	log.Println("PUT "+r.URL+" Time taken:", duration)

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
