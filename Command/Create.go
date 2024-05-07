package Command

import "os"

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

func (c *Create) Init() {
	c.Command.Keymap["help"] = c.Command.help
	c.Command.Keymap["monorepo"] = c.Monorepo
}

func createFolder(local string, name string) {
	os.Mkdir(local+"/"+name, 0777)
}

func (c *Create) Monorepo(args []string) {
	src := []string{
		"components",
		"homeworks",
		"small-projects",
		"test-zone",
		"answer",
	}
	
	name := "Monorepo"
	if len(args) > 0 {
		name = args[0]
	}

	
	createFolder(".", name)
	for _, folder := range src {
		createFolder(name, folder)
	}
}