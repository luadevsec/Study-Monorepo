package main


import (
	"StudyMonorepo/Command"
	"fmt"
	"strings"
	"bufio"
	"os"
)

var CommandMap = map[string] func (string, []string) {
	"Command": Command.C.Execute,
	"Test": Command.T.Execute,
	"Create": Command.Creat.Execute,
}


func Execute(cmd string, key string ,args []string) {
	if fn, ok := CommandMap[cmd]; ok {
		fn(key, args)
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
	Command.C.Init()
	Command.T.Init()
	Command.Creat.Init()

	for {
		scanner := bufio.NewScanner(os.Stdin)

		
		var input string

		fmt.Print("Enter command: ")
		scanner.Scan()
		input = scanner.Text()
		Execute(ParseCommand(input))
	}


}

