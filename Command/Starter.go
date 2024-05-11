package Command

import (
	"StudyMonorepo/utils"
	"os/exec"
)

type Starter struct {
	Command
}

func (c *Starter) Start(path *string, args []string) {
	temp := *path
	if len(args) > 0 {
		temp = args[0]
	}

	if utils.Contains(utils.ReadMonolist("."), temp) {
		*path = temp
	}

	cmd := exec.Command("code", "./"+*path)
	cmd.Start()
}