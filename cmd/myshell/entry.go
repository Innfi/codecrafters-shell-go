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
	case "cd":
		HandleCommandCd(argument)
	case "cat":
		HandleCommandCat(argument)
	default:
		HandleCommandDefault(argument)
	}
}
