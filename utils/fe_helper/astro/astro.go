package astro

import (
	"fmt"
	"os/exec"

	"github.com/svetozar12/dragon-cli/utils"
)

func InitAstroProject(projectName string) error {
	// Copy template files to the project directory.
	if err := copyTemplateFiles(projectName, "template/frontend/with-astro"); err != nil {
		return err
	}
	// Define and set project dependencies.
	if err := setProjectDependencies(); err != nil {
		return err
	}

	return nil
}

func copyTemplateFiles(projectName, relativeSourcePath string) error {
	// Construct the full source and destination paths.
	fullDestinationPath := projectName + "/apps"
	// Create the `cp` command with the full source and destination paths.
	cmd := exec.Command("cp", "-a", relativeSourcePath+"/.", fullDestinationPath)

	// Run the `cp` command.
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error copying template files: %v", err)
	}

	return nil
}

func setProjectDependencies() error {
	packageNames := []string{
		"astro",
	}
	utils.SetDeps(packageNames)
	return nil
}
