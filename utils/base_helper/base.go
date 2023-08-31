package basehelper

import (
	"os/exec"
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
	// cmd := exec.Command("find", ".", "-type", "f", "-exec", "sed", "-e", "s/@dragon-cli/template/"+projectName+"/g", "-i.bak", "'{}'", projectName)
	cmd := exec.Command("find", ".", "-type", "f", "-exec", "sed", "-e", "s/@dragon-cli-template/"+projectName+"/g", "-i", "{}", "+")
	cmd.Dir = projectName
	// Run the command
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
