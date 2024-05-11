package Command

import (
)

type ICommand interface {
	Execute(string, *string, []string) string
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
	case 5:
		return doMonorepo()
	}

	return nil
}


func doCreator() *Create {
	help_txt := []string{
		"monorepo - create a monorepo",
		"project - create a project",
		"help - show this help",
		"exit - exit the program",
	}

	product := &Create{
		Command: Command{
			Name:     "Create",
			Help_txt: help_txt,
			Keymap:   make(map[string]func(*string, []string)),
		},
	}

	product.Keymap["help"] = product.Command.help
	product.Keymap["monorepo"] = product.Monorepo
	product.Keymap["project"] = product.Project
	return product
}

func doDeleter() *Deleter {
	help_txt := []string{
		"delete - delete command",
		"help - show this help",
		"exit - exit the program",
	}

	product := &Deleter{
		Command: Command{
			Name:     "Delete",
			Help_txt: help_txt,
			Keymap:   make(map[string]func(*string, []string)),
		},
	}

	product.Keymap["help"] = product.Command.help
	product.Keymap["monorepo"] = product.Monorepo
	return product
}

func doTester() *Tester {
	help_txt := []string{
		"Cfolder - create folder",
		"Dfolder - delete folder",
		"Cfile - create file",
		"Dfile - delete file",
		"inferno - create an inferno",
	}

	product := &Tester{
		Command: Command{
			Name:     "Test",
			Help_txt: help_txt,
			Keymap:   make(map[string]func(*string, []string)),
		},
	}

	product.Keymap["help"] = product.Command.help
	product.Keymap["Cfolder"] = product.CreateFolder
	product.Keymap["Dfolder"] = product.DeleteFolder
	product.Keymap["Cfile"] = product.CreateFile
	product.Keymap["Dfile"] = product.DeleteFile
	product.Keymap["inferno"] = product.Inferno
	return product
}

func doHelper() *Helper {
	help_txt := []string{
		"help - show this help",
		"create - create command",
		"delete - delete command",
		"test - test command",
		"exit - exit the program",
	}

	product := &Helper{
		Command: Command{
			Name:     "Helper",
			Help_txt: help_txt,
			Keymap:   make(map[string]func(*string, []string)),
		},
	}

	product.Keymap["help"] = product.Command.help
	product.Keymap[""] = product.Command.help
	return product
}

func doMonorepo() *Monorepo {
	help_txt := []string{
		"help - show this help",
		"list - list all monorepos",
		"enter - enter a monorepo",
		"exit - exit a monorepo",
	}

	product := &Monorepo{
		Command: Command{
			Name:     "Monorepo",
			Help_txt: help_txt,
			Keymap:   make(map[string]func(*string, []string)),
		},
	}

	product.Keymap["help"] = product.Command.help
	product.Keymap[""] = product.Command.help
	product.Keymap["enter"] = product.Enter
	product.Keymap["exit"] = product.Exit
	product.Keymap["list"] = product.List
	return product
}

