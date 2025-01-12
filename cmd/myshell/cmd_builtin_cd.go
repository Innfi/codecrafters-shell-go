package main

import (
	"fmt"
	"os"
)

func HandleCommandCd(argument *Argument) {
	effectivePath := toEffectivePath(argument)
	// fmt.Printf("effectivePath: %s\n", effectivePath)
	if err := os.Chdir(effectivePath); err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", effectivePath)
	}
}

func toEffectivePath(argument *Argument) string {
	if argument.params[0] != '~' {
		return argument.params
	}

	home := os.Getenv("HOME")
	if len(argument.params) > 1 {
		return fmt.Sprintf("%s%s", home, argument.params[1:])
	}

	return home
}
