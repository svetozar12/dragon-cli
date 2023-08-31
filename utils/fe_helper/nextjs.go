package fehelper

import (
	"os/exec"

	"github.com/svetozar12/dragon-cli/constants"
	"github.com/svetozar12/dragon-cli/utils"
)

func initNextJSProject(projectName string) error {
	cmd := exec.Command("cp", "-a", "template/frontend/with-nextjs/.", projectName+"/apps")
	err := cmd.Run()
	if err != nil {
		return err
	}

	packageNames := []string{
		"next",
		"react",
		"react-dom",
		"tslib"}
	utils.SetDeps(packageNames)
	devPackageNames := []string{
		"@nx/next", "eslint-config-next"}
	devPackageNames = append(devPackageNames, constants.CommonDevFe...)
	utils.SetDevDeps(devPackageNames)
	return nil
}
