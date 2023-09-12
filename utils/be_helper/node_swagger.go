package behelper

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/svetozar12/dragon-cli/constants"
	"github.com/svetozar12/dragon-cli/utils"
)

func initNodejsSwaggerProject(projectName string) error {
	copyFolderContent := true
	err := utils.CopyTemplateFromRepo("backend/with-golang", projectName+"/apps", copyFolderContent)
	if err != nil {
		return fmt.Errorf("error creating project directory: %v", err)
	}

	// Initialize the Node.js Swagger library.
	if err := initNodejsSwaggerLib(projectName); err != nil {
		return err
	}

	// Define and set project dependencies.
	if err := setNodejsSwaggerProjectDependencies(); err != nil {
		return err
	}

	return nil
}

func setNodejsSwaggerProjectDependencies() error {
	packageNames := []string{
		"axios",
		"express",
		"swagger-jsdoc",
		"swagger-ui-express",
		"tslib",
		"@openapitools/openapi-generator-cli",
		"mongoose",
		"zod",
	}
	devPackageNames := []string{
		"@nx/express",
		"@nx/jest",
		"@nx/node",
		"@nx/webpack",
		"@types/express",
		"@types/jest",
		"@types/node",
		"jest",
		"jest-environment-node",
		"ts-jest",
		"ts-node",
		"@types/swagger-ui-express",
		"@types/swagger-jsdoc",
		"@types/mongoose",
	}
	devPackageNames = append(devPackageNames, constants.CommonDevFe...)
	utils.SetDeps(packageNames, devPackageNames)
	return nil
}

func initNodejsSwaggerLib(projectName string) error {
	copyFolderContent := true
	err := utils.CopyTemplateFromRepo("libs/with-nodejs", projectName+"/libs", copyFolderContent)
	if err != nil {
		return err
	}
	modifyNodejsTsconfig(projectName)
	return nil
}

func modifyNodejsTsconfig(projectName string) {
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
