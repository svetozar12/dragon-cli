package react

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/svetozar12/dragon-cli/constants"
	"github.com/svetozar12/dragon-cli/installers"
	"github.com/svetozar12/dragon-cli/utils"
)

func InitReactProject(projectName string) error {
	// Copy template files to the project directory.
	copyFolderContent := true
	err := utils.CopyTemplateFromRepo("frontend/with-react", projectName+"/apps", copyFolderContent)
	if err != nil {
		return err
	}

	// Define and set project dependencies.
	if err := setReactProjectDependencies(); err != nil {
		return err
	}
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

func setReactProjectDependencies() error {
	packageNames := []string{
		installers.SWC_HELPERS,
		installers.AXIOS,
		installers.REACT,
		installers.REACT_DOM,
		installers.REACT_ROUTER_DOM,
		installers.TSLIB,
	}
	devPackageNames := []string{}
	devPackageNames = append(devPackageNames, constants.CommonDevFe...)
	utils.SetDeps(packageNames, devPackageNames)
	return nil
}
