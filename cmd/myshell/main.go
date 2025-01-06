package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		// Wait for user input
		fmt.Fprint(os.Stdout, "$ ")

		arguments, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			os.Exit(1)
		}
		arguments = strings.Replace(arguments, "\r", "", -1)
		arguments = strings.Replace(arguments, "\n", "", -1)

		command, params, _ := strings.Cut(arguments, " ")
		if command == "echo" {
			fmt.Println(params)
			continue
		} else if arguments[0:4] == "exit" {
			os.Exit(0)
		}

		fmt.Printf("%s: command not found\n", arguments)
	}
}
