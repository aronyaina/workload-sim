package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aronyaina/workload-sim/handler"
	"github.com/aronyaina/workload-sim/utils"
)

func main() {
	if len(os.Args) < 5 {
		log.Fatal("Usage: <main> <file_path> <limit_requests> <timeout_seconds> <requests_per_second>")
	}

	filePath := os.Args[1]
	limit, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Limit must be a number: %v", err)
	}

	timeout, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalf("Timeout must be a number: %v", err)
	}

	requestsPerSecond, err := strconv.Atoi(os.Args[4])
	if err != nil {
		log.Fatalf("Requests per second must be a number: %v", err)
	}

	requests, err := utils.ParseJSONFile(filePath)
	if err != nil {
		log.Fatalf("Error parsing JSON file: %v", err)
	}

	if len(requests) > limit {
		requests = requests[:limit]
	}

	cancel := make(chan struct{})
	go handler.HandleRequestsConcurrently(requests, cancel, requestsPerSecond, time.Duration(timeout)*time.Second)

	select {}
}
