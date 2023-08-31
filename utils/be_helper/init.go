package behelper

func InitBeProject(projectName string, framework string) error {
	switch framework {
	case "typescript(swagger)":
		err := initNodejsSwaggerProject(projectName)
		if err != nil {
			return err
		}
	default:
		return nil
	}
	return nil
}
