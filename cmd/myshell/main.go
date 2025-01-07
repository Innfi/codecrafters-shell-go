package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Argument struct {
	command string
	params  string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Wait for user input
		fmt.Fprint(os.Stdout, "$ ")

		arguments, err := readArguments(reader)
		if err != nil {
			fmt.Printf("read error\n")
			continue
		}

		handleCommand(arguments)
	}
}

// func
func readArguments(reader *bufio.Reader) (*Argument, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	input = strings.Replace(input, "\r", "", -1)
	input = strings.Replace(input, "\n", "", -1)

	command, params, _ := strings.Cut(input, " ")

	return &Argument{
		command,
		params,
	}, nil
}

func handleCommand(argument *Argument) {
	switch argument.command {
	case "exit":
		os.Exit(0)
	case "echo":
		fmt.Println(argument.params)
	case "type":
		handleCommandType(argument)
	default:
		fmt.Printf("%s: command not found\n", argument.command)
	}
}

func handleCommandType(argument *Argument) {
	builtIns := []string{"echo", "exit"}

	for idx := range builtIns {
		if builtIns[idx] == argument.params {
			fmt.Printf("%s is a shell builtin\n", argument.params)
			return
		}
	}

	fmt.Printf("%s: not found\n", argument.params)
}
