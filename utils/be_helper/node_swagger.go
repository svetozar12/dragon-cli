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
	cmd := exec.Command("cp", "-a", "template/backend/swagger/with-nodejs/.", projectName+"/apps")
	err := cmd.Run()
	if err != nil {
		return err
	}
	err = initNodejsSwaggerLib(projectName)
	if err != nil {
		return err
	}
	packageNames := []string{
		"axios",
		"express",
		"swagger-jsdoc",
		"swagger-ui-express",
		"tslib",
		"@openapitools/openapi-generator-cli"}
	utils.SetDeps(packageNames)
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
	}
	devPackageNames = append(devPackageNames, constants.CommonDevFe...)
	utils.SetDevDeps(devPackageNames)
	return nil
}

func initNodejsSwaggerLib(projectName string) error {
	cmd := exec.Command("cp", "-a", "template/libs/swagger/with-nodejs/.", projectName+"/libs")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	fmt.Printf("Executing command: %s\n", cmd.String()) // Print the command
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

	fmt.Println("Key-value pair added successfully.")
}
