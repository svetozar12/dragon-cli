package utils

import (
	"fmt"
	"os"
	"os/exec"
)

var deps []string
var devDeps []string

func SetDeps(newDeps []string) {
	deps = append(deps, newDeps...)
	deps = removeDuplicates(deps)
}

func SetDevDeps(newDevDeps []string) {
	devDeps = append(devDeps, newDevDeps...)
	devDeps = removeDuplicates(devDeps)
}

func GetDeps() ([]string, []string) {
	return deps, devDeps
}

func removeDuplicates(input []string) []string {
	encountered := map[string]bool{} // Track encountered strings
	result := []string{}             // Slice to store unique strings

	for _, s := range input {
		if !encountered[s] {
			encountered[s] = true
			result = append(result, s)
		}
	}

	return result
}

func AddDependencyAndInstall(packageNames []string, isDev bool, projectName string) error {
	packageManager := "yarn"

	args := []string{"add"}
	if isDev {
		args = append(args, "-D")

	}
	args = append(args, packageNames...)

	// Add the dependencies to package.json
	addCmd := exec.Command(packageManager, args...)
	addCmd.Dir = projectName
	addCmd.Stdout = os.Stdout
	addCmd.Stderr = os.Stderr

	if err := addCmd.Run(); err != nil {
		return fmt.Errorf("error adding dependencies: %v", err)
	}

	// Install the dependencies
	installCmd := exec.Command(packageManager, "install")
	installCmd.Dir = projectName
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr

	if err := installCmd.Run(); err != nil {
		return fmt.Errorf("error installing dependencies: %v", err)
	}

	return nil
}
