package main

import (
	"fmt"
	"os"
)

func HandleCommandCat(argument *Argument) {
	tokenArray := ToTokenArrayRevised(argument.params)

	for _, token := range tokenArray {
		printFile(token)
	}
}

func printFile(path string) {
	if _, err := os.Stat(path); err != nil {
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return
	}

	fmt.Printf("%s", string(data))
}
