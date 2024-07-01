package main

import (
	"log"

	"github.com/aronyaina/workload-sim/utils"
)

func main() {
	filePath := "./seeder/first.json"
	requests, err := utils.ParseJSONFile(filePath)
	if err != nil {
		log.Fatalf("Error parsing JSON file: %v", err)
	}
	utils.HandleRequestsConcurrently(requests)
}
