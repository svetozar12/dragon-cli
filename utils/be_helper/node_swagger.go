package behelper

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/svetozar12/dragon-cli/constants"
	"github.com/svetozar12/dragon-cli/installers"
	"github.com/svetozar12/dragon-cli/utils"
)

func initNodejsSwaggerProject(projectName string) error {
	copyFolderContent := true
	err := utils.CopyTemplateFromRepo("backend/with-nodejs", projectName+"/apps", copyFolderContent)
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
		installers.AXIOS,
		installers.SWAGGER_JSDOC,
		installers.SWAGGER_UI_EXPRESS,
		installers.TSLIB,
		installers.OPEN_API_GENERATOR_CLI,
		installers.MONGOOSE,
		installers.ZOD,
	}
	devPackageNames := []string{
		installers.NX_EXPRESS,
		installers.NX_JEST,
		installers.NX_NODE,
		installers.NX_WEBPACK,
		installers.TYPES_EXPRESS,
		installers.TYPES_JEST,
		installers.TYPES_NODE,
		installers.JEST,
		installers.JEST_ENV_NODE,
		installers.TS_JEST,
		installers.TS_NODE,
		installers.TYPES_SWAGGER_UI_EXPRESS,
		installers.TYPES_SWAGGER_JSDOC,
		installers.TYPES_MONGOOSE,
	}
	devPackageNames = append(devPackageNames, constants.CommonDevFe...)
	utils.SetDeps(packageNames, devPackageNames)
	return nil
}

func initNodejsSwaggerLib(projectName string) error {
	copyFolderContent := true
	err := utils.CopyTemplateFromRepo("libs", projectName+"/libs", copyFolderContent)
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
