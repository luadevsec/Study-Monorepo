package Command

import (
	"StudyMonorepo/utils"
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

func (c *Create) Init() {
	c.Command.Keymap["help"] = c.Command.help
	c.Command.Keymap["monorepo"] = c.Monorepo
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

	utils.CreateFolder(".", name)
	for _, folder := range src {
		utils.CreateFolder(name, folder)
	}
}

func (c *Create) Component(args []string) {}

func (c *Create) Homework(args []string) {}

func (c *Create) SmallProject(args []string) {}