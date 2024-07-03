package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aronyaina/workload-sim/utils"
)

func main() {
	if len(os.Args) < 5 {
		log.Fatal("Usage: <main> <file_path> <limit_requests> <timeout_seconds> <requests_per_second>")
	}

	filePath := os.Args[1]
	limit := os.Args[2]
	t_timeout := os.Args[3]
	request_second := os.Args[4]

	l, err := strconv.Atoi(limit)
	if err != nil {
		log.Fatalf("Limit must be a number: %v", err)
	}

	t, err := strconv.Atoi(t_timeout)
	if err != nil {
		log.Fatalf("Timeout must be number: %v", err)
	}

	ts, err := strconv.Atoi(request_second)
	if err != nil {
		log.Fatalf("Request per second must be number: %v", err)
	}

	requests, err := utils.ParseJSONFile(filePath)
	if err != nil {
		log.Fatalf("Error parsing JSON file: %v", err)
	}

	timeout := time.After(time.Duration(t) * time.Second)
	cancel := make(chan struct{})

	if len(requests) > l {
		requests = requests[:l]
	}

	go func() {
		utils.HandleRequestsConcurrently(requests, cancel, ts)
	}()

	select {
	case <-timeout:
		close(cancel)
		log.Println("Timed out , Stopping request")
	}
}
