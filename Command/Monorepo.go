package Command

import (
	"StudyMonorepo/utils"
	"bufio"
	"fmt"
	"os"
)

type Monorepo struct {
	Command
}

func (c *Monorepo) Enter(path *string, args []string) {
	mono := args[0]
	if mono == "" {
		print("Enter the monorepo name: ")
		entrada := bufio.NewScanner(os.Stdin)
		input := entrada.Text()
		mono = input
	}

	monos := utils.ReadMonolist(".")
	if !utils.Contains(monos, mono) {
		fmt.Println("Monorepo not found")
		return
	}
	
	*path = mono
}

func (c *Monorepo) List(path *string, args []string) {
	monos := utils.ReadMonolist(".")
	for id, mono := range monos {
		fmt.Println(id, mono)
	}
}

func (c *Monorepo) Exit (path *string, args []string) {
	*path = "."
}
