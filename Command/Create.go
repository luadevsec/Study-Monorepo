package Command

import (
	"StudyMonorepo/utils"
)

type Create struct {
	Command
}


func (c *Create) Monorepo(path *string, args []string) {
	src := []string{
		"components",
		"homeworks",
		"projects",
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

func (c *Create) Project (path *string, args []string) {
	name := "New_Project"
	src := []string{
		"model",
		"view",
		"controller",
		"test",
		"docs",
		"assets",

	}
	if len(args) > 0 {
		name = args[0]
	}

	utils.CreateFolder(*path+"/projects", name)
	for _, folder := range src {
		utils.CreateFolder(*path+"/projects/"+name, folder)
	}
}