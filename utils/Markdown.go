package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var monorepo = "- name\n"

func CreateMonolist (path string, name string) {
	name = strings.Replace(monorepo, "name", name, 1)
	monolist, err := os.OpenFile(path + "/" + "Monolist" + ".md", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)

	if err != nil {
		log.Fatal(err)
	}

	defer monolist.Close()

	_, err = monolist.WriteString(name)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Monolist created")
	}
}

func ReadMonolist (path string)  (monos []string) {
	
	monolist, err := os.Open(path + "/" + "Monolist" + ".md")

	if err != nil {
		log.Fatal(err)
	}

	defer monolist.Close()

	ler := bufio.NewScanner(monolist)

	for ler.Scan() {
		line := ler.Text()
		parts := strings.Split(line, "-")
		mono := strings.TrimSpace(parts[1])
		monos = append(monos, mono)
	}

	return
}

func DeleteMonorepo (path string, name string) {

	content := ReadMonolist(path)
	monolist, err := os.OpenFile(path + "/" + "Monolist" + ".md", os.O_WRONLY|os.O_TRUNC, 0777)

	if err != nil {
		log.Fatal(err)
	}

	defer monolist.Close()

	for _ , mono := range content {
		if mono != name {
			monolist.WriteString("- " + mono + "\n")
		}
	}
}
	
