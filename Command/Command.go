package Command

import (
	"fmt"
)


type Command struct {
	Name string
	Help_txt []string
	Keymap map[string]func(*string, []string)
}

func (c *Command) Execute(cmd string, path *string, args []string) (void string){
	if fn, ok:= c.Keymap[cmd]; ok {
		fn(path, args)
	} else {
		fmt.Println("Command not found")
	}

	return
}

func (c *Command) help (path *string, args []string) {
	for _, v := range c.Help_txt {
		fmt.Println(v)
	}	
}

