package main

func HandleCommand(argument *Argument) {
	switch argument.command {
	case "exit":
		HandleCommandExit()
	case "echo":
		HandleCommandEcho(argument)
	case "type":
		HandleCommandType(argument)
	case "pwd":
		HandleCommandPwd()
	default:
		HandleCommandDefault(argument)
	}
}
