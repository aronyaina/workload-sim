package utils

import (
	"encoding/json"
	"os"
)

type Request struct {
	URL    string            `json:"url"`
	Method string            `json:"method"`
	Params map[string]string `json:"params,omitempty"`
	Form   map[string]string `json:"form,omitempty"`
}

func ParseJSONFile(filePath string) ([]Request, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var requests []Request
	err = json.Unmarshal(data, &requests)
	if err != nil {
		return nil, err
	}

	return requests, nil
}
