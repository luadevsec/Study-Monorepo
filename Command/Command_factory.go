package Command

import (
)

type ICommand interface {
	Execute(string, []string)
}

func CommandFactory(who int) ICommand{
	switch who {
	case 1:
		return doCreator()
	case 2:
		return doDeleter()
	}
	return nil
}

func doCreator() *Create{
	help_txt := []string{
		"create - create command", 
		"help - show this help", 
		"exit - exit the program",
	}

	product := &Create{
		Command: Command{
			Name:     "Create",
			Help_txt: help_txt,
			Keymap:   map[string]func([]string){},
		},
	}

	product.Command.Keymap = map[string]func([]string){
		"help": product.Command.help,
		"monorepo": product.Monorepo,
	}

	return product
}

func doDeleter() *Deleter{
	help_txt := []string{
		"delete - delete command", 
		"help - show this help", 
		"exit - exit the program",
	}

	product := &Deleter{
		Command: Command{
			Name:     "Delete",
			Help_txt: help_txt,
			Keymap:   map[string]func([]string){},
		},
	}

	product.Command.Keymap = map[string]func([]string){
		"help": product.Command.help,
		"monorepo": product.Monorepo,
	}


	return product
}
