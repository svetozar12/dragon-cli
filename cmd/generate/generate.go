package generate

import (
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/svetozar12/dragon-cli/constants"
	"github.com/svetozar12/dragon-cli/utils"
	basehelper "github.com/svetozar12/dragon-cli/utils/base_helper"
	behelper "github.com/svetozar12/dragon-cli/utils/be_helper"
	fehelper "github.com/svetozar12/dragon-cli/utils/fe_helper"
)

func Generate(cmd *cobra.Command, args []string) {
	// clone repo
	branch, _ := cmd.Flags().GetString("branch")
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting user's home directory:", err)
		panic("Function RenameDefaultNames() failed" + err.Error())
	}
	tmpRepoDir := filepath.Join(usr.HomeDir, "dragon-cli-tmp")

	go func() {
		err := utils.CloneTemplateRepo(tmpRepoDir, branch)
		if err != nil {
			panic("Template Repository isn't available at the moment or the is some problem with the cli tool")
		}
	}()
	defer func() {
		utils.DeleteTemplateRepo(tmpRepoDir)
	}()

	// flags
	projectName, _ := cmd.Flags().GetString("projectName")
	beFramework, _ := cmd.Flags().GetString("beFramework")
	feFramework, _ := cmd.Flags().GetString("feFramework")
	installDeps, _ := cmd.Flags().GetString("installDeps")

	if projectName == "" {
		projectName = utils.GetInput("Project Name")
	}
	if beFramework == "" {
		beFramework = utils.GetCheckbox(
			constants.BeFrameworkLabel,
			constants.BeFrameworkList,
		)
	}
	if feFramework == "" {
		feFramework = utils.GetCheckbox(
			constants.FeFrameworkLabel,
			constants.FeFrameworkList,
		)
	}
	if installDeps == "" {
		installDeps = utils.GetBooleanInput("Do you want to install dependencies ?")
	}
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
	if installDeps == "true" {
		err := utils.InstallDependencies(projectName, "yarn")
		if err != nil {
			panic("Function InstallDependencies() failed" + err.Error())
		}
	}

}
