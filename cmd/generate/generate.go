package generate

import (
	"fmt"

	"github.com/svetozar12/dragon-cli/constants"
	"github.com/svetozar12/dragon-cli/utils"
	basehelper "github.com/svetozar12/dragon-cli/utils/base_helper"
	behelper "github.com/svetozar12/dragon-cli/utils/be_helper"
	fehelper "github.com/svetozar12/dragon-cli/utils/fe_helper"
)

func Generate() {
	projectName := utils.GetInput("Project Name")
	beFramework := utils.GetCheckbox(
		constants.BeFrameworkLabel,
		constants.BeFrameworkList,
	)
	feFramework := utils.GetCheckbox(
		constants.FeFrameworkLabel,
		constants.FeFrameworkList,
	)
	err := basehelper.CreateProjectDir(projectName)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	err = basehelper.RenameDefaultNames(projectName)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	err = behelper.InitBeProject(projectName, beFramework)
	if err != nil {
		panic(err)
	}
	err = fehelper.InitFeProject(projectName, feFramework)
	if err != nil {
		panic(err)
	}
	deps, devDeps := utils.GetDeps()
	err = utils.AddDependencyAndInstall(deps, false, projectName)
	if err != nil {
		panic(err)
	}
	err = utils.AddDependencyAndInstall(devDeps, false, projectName)
	if err != nil {
		panic(err)
	}
}
