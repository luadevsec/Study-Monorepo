package Command

import (
	"StudyMonorepo/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
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

func (c *Create) Homework (path *string, args []string) {
	name := "New_Homework"
	
	if len(args) > 0 {
		name = args[0]
	}

	group := ""
	if len(args) > 1 {
		group = args[1]
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("if group name is empty, the homework will be created in the default group")
		fmt.Print("Enter the group name: ")
		scanner.Scan()
		group = scanner.Text()
	}

	if group == "" {
		group = "default"
	}


	utils.CreateFolder(*path+"/homeworks", group)
	if strings.Contains(name, ".") {
		utils.CreateFile(*path+"/homeworks/"+group, name)
	} else {
		utils.CreateFolder(*path+"/homeworks/"+group, name)
	}
}

func (c *Create) Component (path *string, args []string) {
	name := "New_Component"
	
	if len(args) > 0 {
		name = args[0]
	}

	group := ""
	if len(args) > 1 {
		group = args[1]
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("if group name is empty, the component will be created in the utils group")
		fmt.Print("Enter the group name: ")
		scanner.Scan()
		group = scanner.Text()
	}

	if group == "" {
		group = "utils"
	}

	utils.CreateFolder(*path+"/components", group)

	for !strings.Contains(name, ".") {
		fmt.Println("Component name must have a file extension")
		fmt.Println("Example: New_Component.go")
		fmt.Print("Enter the component name: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		name = scanner.Text()
	}

	utils.CreateFolder(*path+"/components", group)
	utils.CreateFile(*path+"/components/"+group, name)

}

func (c *Create) Test (path *string, args []string) {
	name := "New_Test"
	
	if len(args) > 0 {
		name = args[0]
	}

	if strings.Contains(name, ".") {
		utils.CreateFile(*path+"/test-zone", name)
	} else {
		utils.CreateFolder(*path+"/test-zone", name)
	}
}
