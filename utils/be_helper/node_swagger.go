package behelper

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/svetozar12/dragon-cli/constants"
	"github.com/svetozar12/dragon-cli/utils"
)

func initNodejsSwaggerProject(projectName string) error {
	copyFolderContent := true
	err := utils.CopyTemplateFromRepo("template/backend/swagger/with-nodejs", projectName+"/apps", copyFolderContent)
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

	// Define and set development dependencies.
	if err := setNodejsSwaggerDevelopmentDependencies(); err != nil {
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
	utils.SetDeps(packageNames)
	return nil
}

func setNodejsSwaggerDevelopmentDependencies() error {
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
	utils.SetDevDeps(devPackageNames)
	return nil
}

func copyTemplateFiles(projectName, relativeSourcePath string) error {
	// Get the current working directory.
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working directory: %v", err)
	}

	// Construct the full source and destination paths.
	fullSourcePath := currentDir + "/" + relativeSourcePath
	fullDestinationPath := projectName + "/apps"

	// Create the `cp` command with the full source and destination paths.
	cmd := exec.Command("cp", "-a", fullSourcePath+"/.", fullDestinationPath)

	// Run the `cp` command.
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error copying template files: %v", err)
	}

	return nil
}

func initNodejsSwaggerLib(projectName string) error {
	copyFolderContent := true
	err := utils.CopyTemplateFromRepo("template/libs/swagger/with-nodejs", projectName+"/libs", copyFolderContent)
	if err != nil {
		return err
	}
	modifyJson(projectName)
	return nil
}

func modifyJson(projectName string) {
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
