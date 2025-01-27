package main

import (
	"fmt"
	"strings"
)

func HandleCommandEcho(argument *Argument) {
	tokenArray := ToTokenArrayRevised(argument.params)

	fmt.Println(strings.Join(tokenArray, ""))
}
