package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func DecodeJson(filePath string) (map[string]interface{}, *os.File, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, err
	}

	// Decode the existing JSON data
	var data map[string]interface{}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, nil, err
	}
	return data, file, nil
}

func SaveJsonFile(file *os.File, data map[string]interface{}) {
	file.Seek(0, 0)
	encodedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic("Function MarshalIndent() failed" + err.Error())
	}
	if _, err := file.Write(encodedData); err != nil {
		panic("Function Write() failed" + err.Error())
	}

	fmt.Println("Key-value pair added successfully.")
}
