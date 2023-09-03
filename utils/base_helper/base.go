package basehelper

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/svetozar12/dragon-cli/utils"
)

func CreateProjectDir(projectName string) error {
	// Specify the relative path to the source base template directory.
	relativeSourcePath := "template/base"

	// Get the current working directory.
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working directory: %v", err)
	}

	// Construct the full source and destination paths.
	fullSourcePath := currentDir + "/" + relativeSourcePath
	fullDestinationPath := projectName

	// Create the `cp` command with the full source and destination paths.
	cmd := exec.Command("cp", "-r", fullSourcePath, fullDestinationPath)

	// Run the `cp` command.
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error creating project directory: %v", err)
	}

	return nil
}

func RenameDefaultNames(projectName string) error {
	cmd := exec.Command("find", ".", "-type", "f", "-exec", "sed", "-e", "s/@dragon-cli-template/"+projectName+"/g", "-i", "{}", "+")
	cmd.Dir = projectName
	// Run the command
	err := cmd.Run()
	if err != nil {
		return err
	}
	devDeps := []string{
		"nx",
		"@nx/workspace",
	}
	utils.SetDevDeps(devDeps)
	return nil
}
