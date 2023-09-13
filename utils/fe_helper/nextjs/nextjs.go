package nextjs

import (
	"github.com/svetozar12/dragon-cli/constants"
	"github.com/svetozar12/dragon-cli/installers"
	"github.com/svetozar12/dragon-cli/utils"
)

func InitNextJSProject(projectName string) error {
	// Copy template files to the project directory.
	copyFolderContent := true
	err := utils.CopyTemplateFromRepo("frontend/with-nextjs", projectName+"/apps", copyFolderContent)
	if err != nil {
		return err
	}
	// Define and set project dependencies.
	if err := setProjectDependencies(); err != nil {
		return err
	}

	return nil
}

func setProjectDependencies() error {
	packageNames := []string{
		installers.NEXT,
		installers.REACT,
		installers.REACT_DOM,
		installers.TSLIB,
	}
	devPackageNames := []string{
		installers.NX_NEXT,
	}
	devPackageNames = append(devPackageNames, constants.CommonDevFe...)
	utils.SetDeps(packageNames, devPackageNames)
	return nil
}
