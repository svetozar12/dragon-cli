package fehelper

import (
	"github.com/svetozar12/dragon-cli/utils/fe_helper/astro"
	"github.com/svetozar12/dragon-cli/utils/fe_helper/nextjs"
	"github.com/svetozar12/dragon-cli/utils/fe_helper/react"
)

func InitFeProject(projectName string, framework string) error {
	switch framework {
	case "React(with vite)":
		err := react.InitReactProject(projectName)
		if err != nil {
			return err
		}
	case "Nextjs":
		err := nextjs.InitNextJSProject(projectName)
		if err != nil {
			return err
		}
	case "Astro":
		err := astro.InitAstroProject(projectName)
		if err != nil {
			return err
		}
	default:
		return nil
	}
	return nil
}
