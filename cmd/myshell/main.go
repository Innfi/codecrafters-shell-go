package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Wait for user input
		fmt.Fprint(os.Stdout, "$ ")

		arguments, err := ReadArguments(reader)
		if err != nil {
			fmt.Printf("read error\n")
			continue
		}

		HandleCommand(arguments)
	}
}
