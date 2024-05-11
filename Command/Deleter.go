package Command

import (
	"StudyMonorepo/utils"
	"fmt"
	"strings"
)

type Deleter struct {
	Command
}

func (c *Deleter) Monorepo(path *string, args []string) {
	name := "Monorepo"
	if len(args) > 0 {
		name = args[0]
	}

	fmt.Print("Are you sure you want to delete ", name, "? (y/n): ")
	var input string
	fmt.Scanln(&input)
	acept := []string{"y", "Y", "yes", "Yes", "YES", "s", "S", "si", "Si", "SI", "sim", "Sim", "SIM"}
	if !utils.Contains(acept, input) {
		fmt.Println("Aborted")
		return
	}

	utils.DeleteFolder(".", name)
	utils.DeleteMonorepo(".", name)
}

func (c *Deleter) Project(path *string, args []string) {
	name := ""

	if len(args) > 0 {
		name = args[0]
	}

	if name == "" {
		fmt.Println("You must provide a project name")
		return
	}

	fmt.Print("Are you sure you want to delete ", name, "? (y/n): ")
	var input string
	fmt.Scanln(&input)
	acept := []string{"y", "Y", "yes", "Yes", "YES", "s", "S", "si", "Si", "SI", "sim", "Sim", "SIM"}
	if !utils.Contains(acept, input) {
		fmt.Println("Aborted")
		return
	}

	utils.DeleteFolder(*path+"/projects", name)
}

func (c *Deleter) Homework(path *string, args []string) {
	name := ""
	group := ""

	if len(args) > 0 {
		name = args[0]
	}

	if len(args) > 1 {
		group = args[1]
	} else {
		fmt.Print("Enter the group name: ")
		fmt.Scanln(&group)
	}

	if name == "" {
		fmt.Println("You must provide a homework name")
		return
	}

	if group == "" {
		fmt.Println("You must provide a group name")
		return
	}

	fmt.Print("Are you sure you want to delete ", name, " from group ", group, "? (y/n): ")
	var input string
	fmt.Scanln(&input)
	accept := []string{"y", "Y", "yes", "Yes", "YES", "s", "S", "si", "Si", "SI", "sim", "Sim", "SIM"}
	if !utils.Contains(accept, input) {
		fmt.Println("Aborted")
		return
	}

	utils.DeleteFolder(*path+"/homeworks/"+group, name)
}

func (c *Deleter) Test(path *string, args []string) {
	name := ""

	if len(args) > 0 {
		name = args[0]
	}

	if name == "" {
		fmt.Println("You must provide a test name")
		return
	}

	fmt.Print("Are you sure you want to delete ", name, "? (y/n): ")
	var input string
	fmt.Scanln(&input)
	accept := []string{"y", "Y", "yes", "Yes", "YES", "s", "S", "si", "Si", "SI", "sim", "Sim", "SIM"}
	if !utils.Contains(accept, input) {
		fmt.Println("Aborted")
		return
	}

	utils.DeleteFolder(*path+"/tests", name)
}

func (c *Deleter) Component(path *string, args []string) {
	name := ""

	if len(args) > 0 {
		name = args[0]
	}

	group := ""

	if len(args) > 1 {
		group = args[1]
	} else {
		fmt.Print("Enter the group name: ")
		fmt.Scanln(&group)
	}

	if name == "" {
		fmt.Println("You must provide a component name")
		return
	}

	for !strings.Contains(name, ".") {
		fmt.Println("Component name must have a file extension")
		fmt.Println("Example: New_Component.go")
		fmt.Print("Enter the component name: ")
		fmt.Scanln(&name)
	}

	fmt.Print("Are you sure you want to delete ", name, " of ", group, "? (y/n): ")
	var input string
	fmt.Scanln(&input)
	accept := []string{"y", "Y", "yes", "Yes", "YES", "s", "S", "si", "Si", "SI", "sim", "Sim", "SIM"}
	if !utils.Contains(accept, input) {
		fmt.Println("Aborted")
		return
	}

	utils.DeleteFolder(*path+"/components/"+group, name)
}
