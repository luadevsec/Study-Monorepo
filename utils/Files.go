package utils

import (
	"os"
)

func CreateFolder(local string, name string) {
	os.Mkdir(local+"/"+name, 0777)
}

func DeleteFolder(local string, name string) {
	os.RemoveAll(local + "/" + name)
}

func CreateFile(local string, name string) {
	file, _ := os.Create(local + "/" + name)
	defer file.Close()
}

func DeleteFile(local string, name string) {
	os.Remove(local + "/" + name)
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}