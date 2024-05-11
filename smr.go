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
		command.Execute(key, &Path ,args)
	} else {
		fmt.Println("Comando ", cmd, " não encontrado, digite help para ver os comandos disponíveis")
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

var Path string = "."

func main() {
	CommandMap["create"] = Command.CommandFactory(1)
	CommandMap["delete"] = Command.CommandFactory(2)
	CommandMap["test"] = Command.CommandFactory(3)
	CommandMap["help"] = Command.CommandFactory(4)
	CommandMap["monorepo"] = Command.CommandFactory(5)
	CommandMap["mono"] = CommandMap["monorepo"]


	for forever := true; forever;{
		scanner := bufio.NewScanner(os.Stdin)
		
		var input string

		fmt.Print(Path, " smr> ")
		scanner.Scan()
		input = scanner.Text()
		if input == "exit" {
			forever = false
			break
		}
		Execute(ParseCommand(input))
	}


}

