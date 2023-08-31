package behelper

import (
	"os/exec"

	"github.com/svetozar12/dragon-cli/constants"
	"github.com/svetozar12/dragon-cli/utils"
)

func initNodejsSwaggerProject(projectName string) error {
	cmd := exec.Command("cp", "-a", "template/backend/swagger/with-nodejs/.", projectName+"/apps")
	err := cmd.Run()
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
