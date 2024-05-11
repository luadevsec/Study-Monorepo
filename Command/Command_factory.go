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
	case 3:
		return doTester()
	case 4:
		return doHelper()
	}
	return nil
}

func doCreator() (product *Create){
	help_txt := []string{
		"create - create command", 
		"help - show this help", 
		"exit - exit the program",
	}

	product = &Create{
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
	return 
}

func doDeleter() (product *Deleter){
	help_txt := []string{
		"delete - delete command", 
		"help - show this help", 
		"exit - exit the program",
	}

	product = &Deleter{
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
	return 
}

func doTester() (product *Tester) {
	help_txt := []string{
		"Cfolder - create folder",
		"Dfolder - delete folder",
		"Cfile - create file",
		"Dfile - delete file",
		"inferno - create an inferno",
	}

	product = &Tester{
		Command: Command{
			Name:     "Test",
			Help_txt: help_txt,
			Keymap:   map[string]func([]string){},
		},
	}

	product.Command.Keymap = map[string]func([]string){
		"help": product.Command.help,
		"Cfolder": product.CreateFolder,
		"Dfolder": product.DeleteFolder,
		"Cfile": product.CreateFile,
		"Dfile": product.DeleteFile,
		"inferno": product.Inferno,
	}
	return 
}

func doHelper() (product *Helper) {
	help_txt := []string{
		"help - show this help", 
		"create - create command",
		"delete - delete command",
		"test - test command",
		"exit - exit the program",
	}

	product = &Helper{
		Command: Command{
			Name:     "Helper",
			Help_txt: help_txt,
			Keymap:   map[string]func([]string){},
		},
	}

	product.Command.Keymap = map[string]func([]string){
		"help": product.Command.help,
		"": product.Command.help,
		"create": doCreator().help,
		"delete": doDeleter().help,
		"test": doTester().help,
	}
	return 
}
