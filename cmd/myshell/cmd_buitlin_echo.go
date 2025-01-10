package main

import "fmt"

func HandleCommandEcho(argument *Argument) {
	fmt.Println(argument.params)
}
