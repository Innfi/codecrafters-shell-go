package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		os.Exit(1)
	}

	if command[0:7] == "invalid" {
		// fmt.Fprint(os.Stdout, "invalid_command: command not found")
		fmt.Printf("%s: command not found\n", command)
	}
}
