package basehelper

import (
	"os/exec"

	"github.com/svetozar12/dragon-cli/utils"
)

func CreateProjectDir(projectName string) error {
	cmd := exec.Command("cp", "-r", "template/base", projectName)

	// Run the command
	err := cmd.Run()
	if err != nil {
		return err
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
