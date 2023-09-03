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
	installDeps := utils.GetBooleanInput("Do you want to install dependencies ?")
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
	utils.AddDependency(deps, false, projectName)
	utils.AddDependency(devDeps, true, projectName)
	if installDeps {
		err := utils.InstallDependencies(projectName, "yarn")
		if err != nil {
			panic("Function InstallDependencies() failed" + err.Error())
		}
	}

}
