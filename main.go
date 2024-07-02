package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/aronyaina/workload-sim/utils"
)

func main() {

	if len(os.Args) < 4 {
		log.Fatal("Usage: go run main.go <file_path> <limit> <timeout_seconds>")
	}

	filePath := os.Args[1]
	limit := os.Args[2]
	t_timeout := os.Args[3]

	l, err := strconv.Atoi(limit)
	if err != nil {
		log.Fatalf("Value must be a number: %v", err)
	}

	t, err := strconv.Atoi(t_timeout)
	if err != nil {
		log.Fatalf("Value must be number: %v", err)
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
		utils.HandleRequestsConcurrently(requests, cancel)
	}()

	select {
	case <-timeout:
		close(cancel)
		log.Println("Timed out , Stopping request")
	}
}
