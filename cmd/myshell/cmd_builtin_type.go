package main

import (
	"fmt"
	"os"
)

func HandleCommandType(argument *Argument) {
	builtIns := []string{"echo", "exit", "type", "pwd", "cd"}
	for idx := range builtIns {
		if builtIns[idx] == argument.params {
			fmt.Printf("%s is a shell builtin\n", argument.params)
			return
		}
	}

	pathArray, divider := GetPathAndDivider()
	for _, pathElem := range pathArray {
		effectivePath := fmt.Sprintf("%s%s%s", pathElem, divider, argument.params)

		if _, err := os.Stat(effectivePath); !os.IsNotExist(err) {
			fmt.Printf("%s is %s\n", argument.params, effectivePath)
			return
		}
	}

	fmt.Printf("%s: not found\n", argument.params)
}
