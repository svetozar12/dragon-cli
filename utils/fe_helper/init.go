package fehelper

func InitFeProject(projectName string, framework string) error {
	switch framework {
	case "React(with vite)":
		err := initReactProject(projectName)
		if err != nil {
			return err
		}
	case "Nextjs":
		err := initNextJSProject(projectName)
		if err != nil {
			return err
		}
	default:
		return nil
	}
	return nil
}
