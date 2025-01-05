package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		os.Exit(1)
	}

	command = strings.Replace(command, "\r", "", -1)
	command = strings.Replace(command, "\n", "", -1)

	fmt.Printf("%s: command not found", command)
}
