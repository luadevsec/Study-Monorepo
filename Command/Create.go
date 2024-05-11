package Command

import (
	"StudyMonorepo/utils"
)

type Create struct {
	Command
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
	utils.CreateMonolist(".", name)
	for _, folder := range src {
		utils.CreateFolder(name, folder)
	}
}

func (c *Create) Component(args []string) {}

func (c *Create) Homework(args []string) {}

func (c *Create) SmallProject(args []string) {}