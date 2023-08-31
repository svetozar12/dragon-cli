package behelper

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/svetozar12/dragon-cli/constants"
	"github.com/svetozar12/dragon-cli/utils"
)

func initNodejsSwaggerProject(projectName string) error {
	fmt.Println("BEFORE 1")

	cmd := exec.Command("cp", "-a", "template/backend/swagger/with-nodejs/.", projectName+"/apps")
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println("BEFORE 2")

	err = initNodejsSwaggerLib(projectName)
	if err != nil {
		return err
	}
	packageNames := []string{
		"axios",
		"express",
		"swagger-jsdoc",
		"swagger-ui-express",
		"tslib"}
	utils.SetDeps(packageNames)
	devPackageNames := []string{
		"@nx/express",
		"@nx/jest",
		"@nx/node",
		"@nx/webpack",
		"@types/express",
		"@types/jest",
		"@types/node",
		"jest",
		"jest-environment-node",
		"ts-jest",
		"ts-node"}
	devPackageNames = append(devPackageNames, constants.CommonDevFe...)
	utils.SetDevDeps(devPackageNames)
	return nil
}

func initNodejsSwaggerLib(projectName string) error {
	cmd := exec.Command("cp", "-a", "template/libs/swagger/with-nodejs/.", projectName+"/apps")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	fmt.Printf("Executing command: %s\n", cmd.String()) // Print the command
	if err != nil {
		return err
	}
	return nil
}
