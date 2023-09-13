package basehelper

import (
	"fmt"
	"os/exec"

	"github.com/svetozar12/dragon-cli/installers"
	"github.com/svetozar12/dragon-cli/utils"
)

func CreateProjectDir(projectName string) error {
	copyFolderContent := false
	utils.SetDeps([]string{installers.TYPESCRIPT}, []string{})
	err := utils.CopyTemplateFromRepo("base", projectName, copyFolderContent)
	if err != nil {
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

	return nil
}

func InitGit(projectName string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = projectName
	// Run the command
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
