package utils

import "github.com/fatih/color"

func GetErrorColor(message string) string {
	return color.New(color.FgRed).SprintFunc()(message)
}

func GetSuccessColor(message string) string {
	return color.New(color.FgGreen).SprintFunc()(message)
}
