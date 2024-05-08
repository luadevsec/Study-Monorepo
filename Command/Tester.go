package Command

import (
	"StudyMonorepo/utils"
)

type Tester struct {
	Command
}

func (c *Tester) CreateFolder (args []string) {
	name := "FolderTest"
	if len(args) > 0 {
		name = args[0]
	}
	utils.CreateFolder(".", name)
}

func (c *Tester) DeleteFolder (args []string) {
	name := "FolderTest"
	if len(args) > 0 {
		name = args[0]
	}
	utils.DeleteFolder(".", name)
}

func (c *Tester) CreateFile (args []string) {
	name := "FileTest"
	if len(args) > 0 {
		name = args[0]
	}
	utils.CreateFile(".", name)
	utils.CreateFile(".", name + ".txt")
	utils.CreateFile(".", name + ".md")
}

func (c *Tester) DeleteFile (args []string) {
	name := "FileTest"
	if len(args) > 0 {
		name = args[0]
	}
	utils.DeleteFile(".", name)
	utils.DeleteFile(".", name + ".txt")
	utils.DeleteFile(".", name + ".md")
}

