package nextjs

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/svetozar12/dragon-cli/constants"
	"github.com/svetozar12/dragon-cli/utils"
)

func InitNextJSProject(projectName string) error {
	// Copy template files to the project directory.
	if err := copyTemplateFiles(projectName); err != nil {
		return err
	}

	// Define and set project dependencies.
	if err := setProjectDependencies(); err != nil {
		return err
	}

	// Define and set development dependencies.
	if err := setDevelopmentDependencies(); err != nil {
		return err
	}

	return nil
}

func copyTemplateFiles(projectName string) error {
	// Get the current working directory.
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working directory: %v", err)
	}

	// Specify the relative path to the source directory.
	relativeSourcePath := "template/frontend/with-nextjs"

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

func setProjectDependencies() error {
	packageNames := []string{
		"next",
		"react",
		"react-dom",
		"tslib",
	}
	utils.SetDeps(packageNames)
	return nil
}

func setDevelopmentDependencies() error {
	devPackageNames := []string{
		"@nx/next",
		"eslint-config-next",
	}
	devPackageNames = append(devPackageNames, constants.CommonDevFe...)
	utils.SetDevDeps(devPackageNames)
	return nil
}
