package Command

import (
	"fmt"
	"StudyMonorepo/utils"
)

type Deleter struct {
	Command
}



func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func (c *Deleter) Monorepo(args []string) {
	

	name := "Monorepo"
	if len(args) > 0 {
		name = args[0]
	}

	fmt.Print("Are you sure you want to delete ", name, "? (y/n): ")
	var input string
	fmt.Scanln(&input)
	acept := []string{"y", "Y", "yes", "Yes", "YES", "s", "S", "si", "Si", "SI", "sim", "Sim", "SIM"}
	if !contains(acept, input){
		fmt.Println("Aborted")
		return
	}

	utils.DeleteFolder(".", name)
}