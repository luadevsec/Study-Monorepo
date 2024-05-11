package Command

import (
	"StudyMonorepo/utils"
	"os"
	"time"
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

func (c *Tester) Inferno (args []string) {
	utils.CreateMonolist(".", "funcionando")
	time.Sleep(2 * time.Second)
	utils.CreateMonolist(".", "meu_monorepo")
	time.Sleep(2 * time.Second)
	utils.CreateMonolist(".", "principal")
	time.Sleep(2 * time.Second)
	utils.CreateMonolist(".", "javinha")
	time.Sleep(2 * time.Second)
	utils.CreateMonolist(".", "goLanguisse")

	time.Sleep(4 * time.Second)

	monos := utils.ReadMonolist(".")

	for num, mono := range monos {
		println(num, mono)
	}

	time.Sleep(10 * time.Second)
	os.Remove("Monolist.md")
}


