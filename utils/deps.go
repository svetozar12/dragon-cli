package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/svetozar12/dragon-cli/installers"
)

var deps []string
var devDeps []string

func SetDeps(newDeps []string, newDevDeps []string) {
	if len(newDeps) > 0 {
		deps = append(deps, newDeps...)
		deps = removeDuplicates(deps)

	}
	if len(newDevDeps) > 0 {
		devDeps = append(devDeps, newDevDeps...)
		devDeps = removeDuplicates(devDeps)
	}
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

func InstallDependencies(projectName string, packageManager string) error {
	cmd := exec.Command(packageManager, "install")
	cmd.Dir = projectName
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error installing dependencies: %v", err)
	}

	return nil
}

func AddDependency(packageList []string, isDev bool, projectName string) error {
	if len(packageList) < 1 {
		return nil
	}
	filePath := projectName + "/package.json"

	data, file, err := DecodeJson(filePath)
	if err != nil {
		return fmt.Errorf("Function DecodeJson() failed: %v", err)
	}
	defer file.Close()

	dependecies := data["dependencies"].(map[string]interface{})
	devDependecies := data["devDependencies"].(map[string]interface{})
	for _, value := range packageList {
		packageVersion := installers.DependencyVersionMap[value]
		if isDev {
			devDependecies[value] = packageVersion
		} else {
			dependecies[value] = packageVersion

		}

	}
	SaveJsonFile(file, data)
	return nil
}
