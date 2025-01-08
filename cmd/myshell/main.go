package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
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
	builtIns := []string{"echo", "exit", "type"}
	for idx := range builtIns {
		if builtIns[idx] == argument.params {
			fmt.Printf("%s is a shell builtin\n", argument.params)
			return
		}
	}

	delimiter, divider := getPathTokens()
	envPath := os.Getenv("PATH")
	pathArray := strings.Split(envPath, delimiter)

	for _, pathElem := range pathArray {
		effectivePath := fmt.Sprintf("%s%s%s", pathElem, divider, argument.params)
		// effectivePath = strings.Replace(effectivePath, "\\", "\\\\", -1)
		// effectivePath = strings.Replace(effectivePath, " ", "\\ ", -1)
		// fmt.Printf("effectivePath: %s\n", effectivePath)

		if _, err := os.Stat(effectivePath); !os.IsNotExist(err) {
			fmt.Printf("%s is %s\n", argument.params, effectivePath)
			return
		} else {
			fmt.Printf("%s\n", err)
		}
	}

	fmt.Printf("%s: not found\n", argument.params)
}

func getPathTokens() (string, string) {
	os := runtime.GOOS

	if os == "windows" {
		return ";", "\\"
	}

	return ":", "/"
}
