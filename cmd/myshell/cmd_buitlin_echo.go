package main

import (
	"fmt"
	"strings"
)

func HandleCommandEcho(argument *Argument) {
	tokenArray := ToTokenArray(argument.params)

	fmt.Println(strings.Join(tokenArray, ""))
}
