package utils

import (
	"os"
)

func CreateFolder(local string, name string) {
	os.Mkdir(local+"/"+name, 0777)
}