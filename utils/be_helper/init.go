package behelper

import "fmt"

func InitBeProject(projectName string, framework string) error {
	fmt.Println(projectName, framework)
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
