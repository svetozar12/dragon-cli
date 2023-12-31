package utils

import (
	"github.com/AlecAivazis/survey/v2"
)

func GetCheckbox(label string, opts []string) string {
	res := ""
	prompt := &survey.Select{
		Message: label,
		Options: opts,
	}
	survey.AskOne(prompt, &res)

	return res
}

func GetInput(label string) string {
	res := ""
	prompt := &survey.Input{Message: label}
	survey.AskOne(prompt, &res)
	return res
}

func GetBooleanInput(label string) string {
	res := false
	prompt := &survey.Confirm{Message: label}
	survey.AskOne(prompt, &res)
	if res {
		return "true"
	}
	return "false"
}
