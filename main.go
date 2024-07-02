package main

import (
	"log"

	"github.com/aronyaina/workload-sim/utils"
)

func main() {
	filePath := "./seeder/projects_test.json"
	requests, err := utils.ParseJSONFile(filePath)
	if err != nil {
		log.Fatalf("Error parsing JSON file: %v", err)
	}

	limit := 100
	if len(requests) > limit {
		requests = requests[:limit]
	}

	utils.HandleRequestsConcurrently(requests)
}
