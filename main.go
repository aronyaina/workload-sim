package main

import (
	"log"
	"time"

	"github.com/aronyaina/workload-sim/utils"
)

func main() {
	filePath := "./seeder/projects_test.json"
	requests, err := utils.ParseJSONFile(filePath)
	if err != nil {
		log.Fatalf("Error parsing JSON file: %v", err)
	}

	timeout := time.After(15 * time.Second)
	cancel := make(chan struct{})
	limit := 100
	if len(requests) > limit {
		requests = requests[:limit]
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
