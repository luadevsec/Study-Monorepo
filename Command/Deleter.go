package Command

import (
	"fmt"
	"os"
)

type Deleter struct {
	Command
}

var Delete = Deleter{
	Command: Command{
		Name:     "Delete",
		Help_txt: []string{"delete - delete command", "help - show this help", "exit - exit the program"},
		Keymap:   map[string]func([]string){},
	},
}

func (c *Deleter) Init() {
	c.Command.Keymap["help"] = c.Command.help
	c.Command.Keymap["monorepo"] = c.Monorepo
}

func deleteFolder(local string, name string) {
	os.RemoveAll(local + "/" + name)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}


func (c *Deleter) Monorepo(args []string) {
	

	name := "Monorepo"
	if len(args) > 0 {
		name = args[0]
	}

	fmt.Print("Are you sure you want to delete ", name, "? (y/n): ")
	var input string
	fmt.Scanln(&input)
	acept := []string{"y", "Y", "yes", "Yes", "YES", "s", "S", "si", "Si", "SI", "sim", "Sim", "SIM"}
	if !contains(acept, input){
		fmt.Println("Aborted")
		return
	}

	deleteFolder(".", name)
}