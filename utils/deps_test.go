package utils

import (
	"os"
	"reflect"
	"testing"

	"github.com/svetozar12/dragon-cli/installers"
)

func TestGetDeps(t *testing.T) {
	t.Run("TestGetDeps correct behavior", func(t *testing.T) {

		newDeps := "new deps"
		SetDeps([]string{newDeps}, []string{newDeps})
		deps, devDeps := GetDeps()
		if deps[len(deps)-1] != newDeps && devDeps[len(devDeps)-1] != newDeps {
			t.Fatalf("GetDeps() doesn't get correct dependencies")
		}
	})
}

func TestRemoveDuplicates(t *testing.T) {
	t.Run("TestRemoveDuplicates correct behavior", func(t *testing.T) {
		result := removeDuplicates([]string{"Same String 1234 !!", "Same String 1234 !!"})
		if len(result) > 1 {
			t.Fatalf("removeDuplicates() should remove strings that are the same")
		}
	})
	t.Run("TestRemoveDuplicates correct behavior with empty array", func(t *testing.T) {
		result := removeDuplicates([]string{})
		if len(result) > 0 {
			t.Fatalf("removeDuplicates() should return empty array")
		}
	})
}

func TestSetDeps(t *testing.T) {
	t.Run("TestSetDeps correct behavior", func(t *testing.T) {
		newDeps := "new deps"
		SetDeps([]string{newDeps}, []string{newDeps})
		if deps[0] != newDeps || devDeps[0] != newDeps {
			t.Fatalf("SetDeps() should set new dependency in array")
		}
	})
	t.Run("TestSetDeps shouldn't have two deps with same name", func(t *testing.T) {
		newDeps := "new deps"
		SetDeps([]string{newDeps, newDeps}, []string{newDeps, newDeps})
		if len(deps) > 1 || len(devDeps) > 1 {
			t.Fatalf("SetDeps() shouldn't have two deps with same name")
		}
	})
}

// Define your test cases
var AddDependencyTestCases = []struct {
	packageName     string
	isDev           bool
	expectedDeps    map[string]interface{}
	expectedDevDeps map[string]interface{} // Define this field for dev dependencies
}{
	{
		packageName: installers.NX,
		isDev:       false,
		expectedDeps: map[string]interface{}{
			installers.NX: installers.DependencyVersionMap[installers.NX],
		},
	},
	{
		packageName: installers.TYPES_MONGOOSE,
		isDev:       true,
		expectedDevDeps: map[string]interface{}{
			installers.TYPES_MONGOOSE: installers.DependencyVersionMap[installers.TYPES_MONGOOSE],
		},
	},
}

var sampleJSON = `{
	"name": "dragon-cli-template",
	"version": "0.0.0",
	"license": "MIT",
	"scripts": {},
	"private": true,
	"dependencies": {},
	"devDependencies": {}
  }
  `

func TestAddDependency(t *testing.T) {
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
	for _, tc := range AddDependencyTestCases {
		t.Run(tc.packageName, func(t *testing.T) {

			AddDependency([]string{tc.packageName}, tc.isDev, projectDir)
			data, _, err := DecodeJson(jsonFilePath)
			if err != nil {
				t.Fatalf("DecodeJson() failed: %v", err)
			}

			dependencies := data["dependencies"].(map[string]interface{})
			devDependencies := data["devDependencies"].(map[string]interface{})

			// Check if the dependencies match the expected result
			for pkg, version := range tc.expectedDeps {
				if dependencies[pkg] != version {
					t.Errorf("Dependency %s does not match the expected version: got %s, expected %s", pkg, dependencies[pkg], version)
				}
			}

			// If it's a dev dependency, check devDependencies too
			if tc.isDev {
				for pkg, version := range tc.expectedDevDeps {
					if devDependencies[pkg] != version {
						t.Errorf("DevDependency %s does not match the expected version: got %s, expected %s", pkg, devDependencies[pkg], version)
					}
				}
			}
		})
	}
	t.Run("TestAddDependency() test if dependency list is empty", func(t *testing.T) {
		beforeAddDeps := devDeps
		AddDependency([]string{}, true, "")

		if !reflect.DeepEqual(beforeAddDeps, devDeps) {
			t.Errorf("devDeps shouldn't have changed")
		}
	})
	t.Run("TestAddDependency() test if package.json is missing", func(t *testing.T) {
		err := AddDependency([]string{"test"}, true, "invalid")
		if err == nil {
			t.Errorf("AddDependency() should throw error")
		}
	})
}

func TestInstallDependency(t *testing.T) {
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
	AddDependency([]string{installers.NX}, false, ".")
	InstallDependencies(".", "yarn")
	nodeModules := "./node_modules"
	filePath := "./yarn.lock"
	if _, err := os.Stat(nodeModules); os.IsNotExist(err) {
		t.Errorf("Folder %s does not exist.\n", nodeModules)
	}

	// Remove the folder
	if err := os.RemoveAll(nodeModules); err != nil {
		t.Errorf("Failed to delete folder %s: %v\n", nodeModules, err)
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("File %s does not exist.\n", filePath)
	}

	// Remove the file
	if err := os.Remove(filePath); err != nil {
		t.Errorf("Failed to delete file %s: %v\n", filePath, err)
	}
	t.Run("TestInstallDependency test with invalid install path", func(t *testing.T) {
		err := InstallDependencies("invalid", "yarn")
		if err == nil {
			t.Errorf("InstallDependencies() should throw error")
		}
	})
}
