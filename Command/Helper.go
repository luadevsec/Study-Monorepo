package Command

import (
)

type Helper struct {
	Command
}

func (c *Helper) CreatorHelp(args []string) {
	
}

func (c *Helper) DeleterHelp(args []string) {
	c.help(args)
}

func (c *Helper) TesterHelp(args []string) {
	c.help(args)
}

func (c *Helper) HelperHelp(args []string) {
	c.help(args)
}

