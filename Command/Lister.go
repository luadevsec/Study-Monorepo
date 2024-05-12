package Command

import (
	"fmt"
	"os/exec"
	"strings"
)

type Lister struct {
	Command
}

func (l *Lister) List(path *string, args []string) {
	if *path == "." {
		fmt.Println("you have to enter in a monorepo first")
		return
	}

	// Comando "tree" com seus argumentos
	cmd := exec.Command("tree", "/f", "/a", *path)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	outputStr := string(output)

	// filtra apenas a arvore de diretórios e arquivos
	pos := strings.Index(outputStr, "+")
	if pos == -1 {
		fmt.Println("Cabeçalho não encontrado")
		return
	}
	result := outputStr[pos:]

	
	
	fmt.Println("\n**" + *path + "** monorepo: ")
	fmt.Println(result)
}
