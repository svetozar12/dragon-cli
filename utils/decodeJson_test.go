package utils

import (
	"os"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	projectDir := "."
	jsonFilePath := projectDir + "/package.json"

	err := os.WriteFile(jsonFilePath, []byte(sampleJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to create a sample package.json file: %v", err)
	}

	defer func() {
		if err := os.Remove(jsonFilePath); err != nil {
			t.Errorf("Failed to delete temporary JSON file: %v", err)
		}
	}()

	json, file, err := DecodeJson(jsonFilePath)
	defer file.Close()

	if json["name"] != "dragon-cli-template" || err != nil {
		t.Errorf("DecodeJson() doesn't work correctly: %v", err)
	}
}

func TestSaveJsonFile(t *testing.T) {
	projectDir := "."
	jsonFilePath := projectDir + "/package.json"

	err := os.WriteFile(jsonFilePath, []byte(sampleJSON), 0644)
	if err != nil {
		t.Fatalf("Failed to create a sample package.json file: %v", err)
	}

	defer func() {
		if err := os.Remove(jsonFilePath); err != nil {
			t.Errorf("Failed to delete temporary JSON file: %v", err)
		}
	}()

	json, file, err := DecodeJson(jsonFilePath)
	defer file.Close()
	json["name"] = "test-dragon-cli"
	SaveJsonFile(file, json)
	newJson, _, _ := DecodeJson(jsonFilePath)
	if newJson["name"] == "dragon-cli-template" || err != nil {
		t.Errorf("SaveJsonFile() doesn't work correctly: %v", err)
	}
}
