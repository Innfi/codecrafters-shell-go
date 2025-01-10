package main

import (
	"fmt"
	"os"
	"os/exec"
)

func HandleCommandDefault(argument *Argument) {
	pathArray, divider := GetPathAndDivider()

	for _, pathElem := range pathArray {
		effectivePath := fmt.Sprintf("%s%s%s", pathElem, divider, argument.command)

		if _, err := os.Stat(effectivePath); !os.IsNotExist(err) {
			cmd := exec.Command(argument.command, argument.params)
			cmd.Stdout = os.Stdout

			if runErr := cmd.Run(); runErr != nil {
				fmt.Println(runErr)
			}
			return
		}
	}

	fmt.Printf("%s: command not found\n", argument.command)
}
