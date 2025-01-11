package main

import (
	"fmt"
	"os"
)

func HandleCommandCd(argument *Argument) {
	if err := os.Chdir(argument.params); err != nil {
		// fmt.Println(err)
		fmt.Printf("cd: %s: no such file or directory\n", argument.params)
	}
}
