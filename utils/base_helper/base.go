package basehelper

import (
	"fmt"
	"os/exec"
)

func CreateProjectDir(projectName string) error {
	cmd := exec.Command("cp", "-r", "template/base", projectName)

	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err.Error())
		return err
	}
	return nil
}
