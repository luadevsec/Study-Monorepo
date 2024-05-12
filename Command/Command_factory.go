package Command

import (
)

type ICommand interface {
	Execute(string, *string, []string) string
}

var helpFL = " help command**\n<> - required\t[] - optional\n"

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
	case 6:
		return doStarter()
	case 7:
		return doLister()
	}

	return nil
}

func doHelper() *Helper {
	help_txt := []string{
		"** Helper" + helpFL,
		"- help \n\tlist all commands",
		"- help <command> \n\tshow help for a command",
		"- create \n\tcreate projects, homeworks, tests, components in the monorepo",
		"- delete \n\tremove something out os smr system",
		"- monorepo \n\tenter, exit, list monorepos",
		"- list \n\tlist all files and directories in the current monorepo",
		"- start \n\tstart vscode with the monorepo",
		"- test \n\ttest the smr system [not implemented yet]",
		
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


func doCreator() *Create {
	help_txt := []string{
		"** Create" + helpFL,
		"- monorepo <Name> \n\t\tcreate a monorepo",
		"- project <Name> \n\t\tcreate a project",
		"- homework <Name || Name.ext> [Group] \n\t\tcreate a homework",
		"- test <Name || Name.ext> \n\t\tcreate a test",
		"- component <Name.ext> <Group> \n\t\tcreate a component",
		
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
	product.Keymap["homework"] = product.Homework
	product.Keymap["test"] = product.Test
	product.Keymap["component"] = product.Component

	return product
}

func doDeleter() *Deleter {
	help_txt := []string{
		"** Delete" + helpFL,
		"- monorepo <Name> \n\t\tdelete a monorepo",
		"- project <Name> \n\t\tdelete a project",
		"- homework <Name || Name.ext> [Group] \n\t\tdelete a homework",
		"- test <Name || Name.ext> \n\t\tdelete a test",
		"- component <Name.ext> <Group> \n\t\tdelete a component",
		
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
	product.Keymap["project"] = product.Project
	product.Keymap["homework"] = product.Homework
	product.Keymap["test"] = product.Test
	product.Keymap["component"] = product.Component
	
	return product
}

func doTester() *Tester {
	help_txt := []string{
		"** Test" + helpFL,
		"- Cfolder \n\t\tcreate folder",
		"- Dfolder \n\t\tdelete folder",
		"- Cfile \n\t\tcreate file",
		"- Dfile \n\t\tdelete file",
		"- inferno \n\t\tcreate an inferno",
		
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


func doMonorepo() *Monorepo {
	help_txt := []string{
		"** Monorepo" + helpFL,
		"- enter <Name> \n\t\tenter a monorepo",
		"- exit \n\t\texit the current monorepo",
		"- list \n\t\tlist all monorepos",
		
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

func doStarter() *Starter {
	help_txt := []string{
		"** Starter" + helpFL,
		"- start \n\t\tstart vscode in the current monorepo",
		
	}

	product := &Starter{
		Command: Command{
			Name:     "Starter",
			Help_txt: help_txt,
			Keymap:   make(map[string]func(*string, []string)),
		},
	}

	product.Keymap["help"] = product.Command.help
	product.Keymap[""] = product.Start
	product.Keymap["start"] = product.Start
	return product
}

func doLister() *Lister {
	help_txt := []string{
		"** List" + helpFL,
		"- list \n\t\tlist all files and directories in the current monorepo",
		
	}

	product := &Lister{
		Command: Command{
			Name:     "List",
			Help_txt: help_txt,
			Keymap:   make(map[string]func(*string, []string)),
		},
	}

	product.Keymap["help"] = product.Command.help
	product.Keymap[""] = product.List
	product.Keymap["list"] = product.List
	return product
}
