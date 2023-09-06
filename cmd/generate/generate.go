package generate

import (
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/svetozar12/dragon-cli/constants"
	"github.com/svetozar12/dragon-cli/utils"
	basehelper "github.com/svetozar12/dragon-cli/utils/base_helper"
	behelper "github.com/svetozar12/dragon-cli/utils/be_helper"
	fehelper "github.com/svetozar12/dragon-cli/utils/fe_helper"
)

func Generate() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting user's home directory:", err)
		panic("Function RenameDefaultNames() failed" + err.Error())
	}

	tmpRepoDir := filepath.Join(usr.HomeDir, "dragon-cli-tmp")

	go func() {
		err := utils.CloneTemplateRepo(tmpRepoDir)
		if err != nil {
			panic("Template Repository isn't available at the moment or the is some problem with the cli tool")
		}
	}()
	defer func() {
		utils.DeleteTemplateRepo(tmpRepoDir)
	}()
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
	err = basehelper.CreateProjectDir(projectName)
	if err != nil {
		panic("Function CreateProjectDir() failed" + err.Error())
	}
	err = behelper.InitBeProject(projectName, beFramework)
	if err != nil {
		panic("Function InitBeProject() failed" + err.Error())
	}
	err = fehelper.InitFeProject(projectName, feFramework)
	if err != nil {
		panic("Function InitFeProject() failed" + err.Error())
	}
	err = basehelper.RenameDefaultNames(projectName)
	if err != nil {
		panic("Function RenameDefaultNames() failed" + err.Error())
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
