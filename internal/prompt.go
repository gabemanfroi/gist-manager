package internal

import (
	"github.com/AlecAivazis/survey/v2"
	"log"
)

type Content struct {
	ErrorMsg string
	Label    string
}

func GetInput(c Content, minLength int) string {

	prompt := &survey.Input{
		Message: c.Label,
	}

	input := ""

	err := survey.AskOne(prompt, &input, survey.WithValidator(survey.MinLength(minLength)))
	if err != nil {
		return ""
	}
	return input
}

func GetMultiline(c Content) string {
	prompt := &survey.Multiline{
		Message: c.Label,
	}

	input := ""

	err := survey.AskOne(prompt, &input, survey.WithValidator(survey.MinLength(1)))
	if err != nil {
		log.Fatal(err.Error())
	}
	return input
}
