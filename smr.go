package main


import (
	"StudyMonorepo/Command"
	"fmt"
	"strings"
	"bufio"
	"os"

)

var CommandMap = map[string] Command.ICommand{}



func Execute(cmd string, key string ,args []string) {
	if command, ok := CommandMap[cmd]; ok {
		command.Execute(key, args)
	} else {
		fmt.Println("Command not found")
	}
}

func ParseCommand(input string) (string, string, []string) {
    parts := strings.Fields(input)

    switch len(parts) {
    case 0:
        return "", "", nil
    case 1:
        return parts[0], "", nil
    case 2:
        return parts[0], parts[1], nil
    default:
        return parts[0], parts[1], parts[2:]
    }
}

func main() {
	CommandMap["create"] = Command.CommandFactory(1)
	CommandMap["delete"] = Command.CommandFactory(2)

	for {
		scanner := bufio.NewScanner(os.Stdin)
		
		var input string

		fmt.Print("Enter command: ")
		scanner.Scan()
		input = scanner.Text()
		Execute(ParseCommand(input))
	}


}

