package Command

import (
	"fmt"
)


type Command struct {
	Name string
	Help_txt []string
	Keymap map[string]func([]string)
}

func (c *Command) Execute(cmd string, args []string) {
	if fn, ok:= c.Keymap[cmd]; ok {
		fn(args)
	} else {
		fmt.Println("Command not found")
	}
}

func (c *Command) help (args []string) {
	for _, v := range c.Help_txt {
		fmt.Println(v)
	}	
}


var C = Command{
	Name: "Command",
	Help_txt: []string{"help - show this help", "exit - exit the program"},
	Keymap: map[string]func([]string){
		
	},
}
