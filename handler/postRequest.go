package handler

import (
	"bytes"
	"context"
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

func PostFormData(r models.Request, cancel <-chan struct{}, ctx context.Context) (*http.Response, error) {
	formData := url.Values{}
	for key, value := range r.Form {
		formData.Set(key, value)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, r.URL, strings.NewReader(formData.Encode()))
	startTime := time.Now()
	resp, err := http.DefaultClient.Do(req)
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

func PostJSONBody(r models.Request, cancel <-chan struct{}, ctx context.Context) (*http.Response, error) {
	bodyBytes, err := json.Marshal(r.Body)
	//log.Println("JSON body:", string(bodyBytes))
	if err != nil {
		log.Printf("Error marshalling body for %s: %v", r.URL, err)
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, r.URL, bytes.NewReader(bodyBytes))
	startTime := time.Now()
	resp, err := http.DefaultClient.Do(req)
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
