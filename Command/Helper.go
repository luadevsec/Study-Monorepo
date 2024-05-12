package Command

import (
)

type Helper struct {
	Command
}

func Help(command ICommand){
	command.Execute("help", nil ,nil)
}