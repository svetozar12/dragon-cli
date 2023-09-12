package behelper

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/svetozar12/dragon-cli/installers"
	"github.com/svetozar12/dragon-cli/utils"
)

func initGolangSwaggerProject(projectName string) error {
	copyFolderContent := true
	err := utils.CopyTemplateFromRepo("backend/with-golang", projectName+"/apps", copyFolderContent)
	if err != nil {
		return fmt.Errorf("error creating project directory: %v", err)
	}

	// Initialize the Node.js Swagger library.
	if err := initGolangSwaggerLib(projectName); err != nil {
		return err
	}

	// Define and set project dependencies.
	if err := setGolangSwaggerProjectDependencies(projectName); err != nil {
		return err
	}

	return nil
}

func setGolangSwaggerProjectDependencies(projectName string) error {
	err := utils.CopyTemplateFromRepo("/go.sum", projectName+"/go.sum", false)
	if err != nil {
		return fmt.Errorf("error creating project directory: %v", err)
	}
	utils.SetDeps([]string{installers.NX_GO, installers.OPEN_API_GENERATOR_CLI}, []string{installers.TYPES_NODE})
	return nil
}

func initGolangSwaggerLib(projectName string) error {
	copyFolderContent := true
	err := utils.CopyTemplateFromRepo("libs/with-golang", projectName+"/libs", copyFolderContent)
	if err != nil {
		return err
	}
	modifyGolangTsconfig(projectName)
	return nil
}

func modifyGolangTsconfig(projectName string) {
	// Open the JSON file for reading and writing
	filePath := projectName + "/tsconfig.base.json"
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Decode the existing JSON data
	var data map[string]interface{}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Update the paths map under compilerOptions
	compilerOptions, ok := data["compilerOptions"].(map[string]interface{})
	if ok {
		paths, ok := compilerOptions["paths"].(map[string]interface{})
		if !ok {
			paths = make(map[string]interface{})
			compilerOptions["paths"] = paths
		}
		paths[projectName+"/shared/sdk"] = []interface{}{"libs/shared/sdk/src/index.ts"}
	}

	// Rewind the file pointer to the beginning
	file.Seek(0, 0)
	encodedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	if _, err := file.Write(encodedData); err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}
