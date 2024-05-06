package Command

import (
	
)

type Create struct {
	Command
}

var Creat = Create{
	Command: Command{
		Name:     "Create",
		Help_txt: []string{"create - create command", "help - show this help", "exit - exit the program"},
		Keymap:   map[string]func([]string){},
	},
}

func (c *Create) create(args []string) {
	println("Create command")
}

func (c *Create) Init() {
	c.Command.Keymap["create"] = c.create
	c.Command.Keymap["help"] = c.Command.help

}