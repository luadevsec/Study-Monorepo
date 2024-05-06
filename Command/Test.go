package Command

import (
	"fmt"
)

type test struct {
	Command
}

func (t *test) test (args []string) {
	fmt.Println("Test command")
}



func (t *test) Init() {
	t.Command.Keymap["test"] = t.test
	t.Command.Keymap["help"] = t.Command.help

}



var T = test{
	Command: Command{
		Name:     "Test",
		Help_txt: []string{"test - test command", "help - show this help", "exit - exit the program"},
		Keymap:   map[string]func([]string){},
	},
}