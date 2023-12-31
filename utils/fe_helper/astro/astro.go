package astro

import (
	"github.com/svetozar12/dragon-cli/installers"
	"github.com/svetozar12/dragon-cli/utils"
)

func InitAstroProject(projectName string) error {
	// Copy template files to the project directory.
	copyFolderContent := true
	err := utils.CopyTemplateFromRepo("frontend/with-astro", projectName+"/apps", copyFolderContent)
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
		installers.ASTRO,
	}
	utils.SetDeps(packageNames, []string{})
	return nil
}
