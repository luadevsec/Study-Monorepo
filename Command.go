package main

import (
	"fmt"
)


type Command struct {
	CommandName string;
	CommandMap map[string] func([]string);
	
}


func (c Command) help(args []string) {
	help := []string{
		"implement help",
		"no seu comando",
		"implemente o help",
		"para mostrar a ajuda do comando",

	}

	for _, h := range help {
		fmt.Println(h)
	}

}


func (c Command) Execute(key string, args []string) {
	if _, ok := c.CommandMap[key]; ok {
		c.CommandMap[key](args);
	} else {
		fmt.Println("Command not found")

	}
}

type test struct {
	Command
}

func (c test) help(args []string) {
	help := []string{
		"implement help",
		"help do comando test",
	}

	for _, h := range help {
		fmt.Println(h)
	}

}

func main() {
	c := test{Command: Command{CommandName: "test", CommandMap: make(map[string]func([]string))}}
	c.CommandMap["test"] = func(args []string) {
		fmt.Println("test")
	}
	c.CommandMap["help"] = c.help
	c.Execute("test", nil)
	c.Execute("help", nil)
	c.Execute("notfound", nil)
}