package main

import (
	"fmt"
	"os"
)

func HandleCommandCd(argument *Argument) {
	err := os.Chdir(argument.params)
	if err != nil {
		fmt.Println(err)
	}
}
