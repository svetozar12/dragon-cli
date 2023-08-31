package fehelper

import (
	"os/exec"

	"github.com/svetozar12/dragon-cli/constants"
	"github.com/svetozar12/dragon-cli/utils"
)

func initReactProject(projectName string) error {
	cmd := exec.Command("cp", "-a", "template/frontend/with-react/.", projectName+"/apps")
	err := cmd.Run()
	if err != nil {
		return err
	}

	packageNames := []string{"@swc/helpers",
		"axios",
		"react",
		"react-dom",
		"react-router-dom",
		"tslib"}
	utils.SetDeps(packageNames)
	utils.SetDevDeps(constants.CommonDevFe)
	return nil
}
