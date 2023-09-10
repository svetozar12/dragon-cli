package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	t.Run("DecodeJson() - test correct behavior", func(t *testing.T) {
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
	})
	t.Run("DecodeJson() - test decoding error", func(t *testing.T) {
		projectDir := "."
		jsonFilePath := projectDir + "/package.json"

		err := os.WriteFile(jsonFilePath, []byte("invalid json"), 0644)
		if err != nil {
			t.Fatalf("Failed to create a sample package.json file: %v", err)
		}

		defer func() {
			if err := os.Remove(jsonFilePath); err != nil {
				t.Errorf("Failed to delete temporary JSON file: %v", err)
			}
		}()

		_, file, err := DecodeJson(jsonFilePath)
		defer file.Close()
		if err == nil {
			t.Errorf("DecodeJson() should throw a error")
		}
	})
}

func TestSaveJsonFile(t *testing.T) {
	t.Run("SaveJsonFile() - correct behavior", func(t *testing.T) {
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
	})
	t.Run("SaveJsonFile() - correct behavior", func(t *testing.T) {
		// Create a temporary file for testing
		tempFile, err := ioutil.TempFile("", "testfile.json")
		if err != nil {
			t.Fatalf("Failed to create a temporary file: %v", err)
		}
		defer tempFile.Close()
		defer os.Remove(tempFile.Name())

		data := map[string]interface{}{
			"key": func() {}, // This will trigger a marshal error
		}
		err = SaveJsonFile(tempFile, data)
		fmt.Println(err)
		if err == nil {
			t.Errorf("Expected an error, but got nil")
		} else {
			t.Logf("Error: %v", err)
		}

		// Close the file to trigger a write error
		tempFile.Close()

		data = map[string]interface{}{
			"key": 1, // This will trigger a marshal error
		}

		err = SaveJsonFile(tempFile, data)
		if err == nil {
			t.Errorf("Expected an error, but got nil")
		} else {
			t.Logf("Error: %v", err)
		}

	})
}
