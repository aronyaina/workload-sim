package utils

import (
	"encoding/json"
	"os"

	"github.com/aronyaina/workload-sim/models"
)

func ParseJSONFile(filePath string) ([]models.Request, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var requests []models.Request
	err = json.Unmarshal(data, &requests)
	if err != nil {
		return nil, err
	}

	return requests, nil
}
