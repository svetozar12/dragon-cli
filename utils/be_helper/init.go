package behelper

func InitBeProject(projectName string, framework string) error {
	switch framework {
	case "nodejs":
		err := initNodejsSwaggerProject(projectName)
		if err != nil {
			return err
		}
	default:
		return nil
	}
	return nil
}
