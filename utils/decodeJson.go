package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func DecodeJson(filePath string) (map[string]interface{}, *os.File, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("Error opening file: %v", err)
	}

	// Decode the existing JSON data
	var data map[string]interface{}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, nil, fmt.Errorf("Error decoding JSON: %v", err)
	}
	return data, file, nil
}

func SaveJsonFile(file *os.File, data map[string]interface{}) error {
	file.Seek(0, 0)
	encodedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("Function MarshalIndent() failed: %v", err)
	}
	if _, err := file.Write(encodedData); err != nil {
		return fmt.Errorf("Function Write() failed: %v", err)
	}
	return nil
}
