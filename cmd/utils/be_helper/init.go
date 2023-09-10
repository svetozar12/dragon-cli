package behelper

func InitBeProject(projectName string, framework string) error {
	switch framework {
	case "nodejs":
		if err := initNodejsSwaggerProject(projectName); err != nil {
			return err
		}
	case "golang":
		if err := initGolangSwaggerProject(projectName); err != nil {
			return err
		}

	default:
		return nil
	}
	return nil
}
