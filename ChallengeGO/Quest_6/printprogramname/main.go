package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	for index, a := range os.Args[0] {

		if index == 0 || index == 1 {

			continue

		}
		z01.PrintRune(a)

	}
	z01.PrintRune('\n')
}

/**Autre Méthode :

if !(index == 0 || index == 1) {
	z01.Printrune(a)
}

/**
on cherche à ecrire un programme qui renvoit le nom du programme !

langage go = langage compilé

Pour que le programme fonction il faut faire un go build, pour que Windows crée un fichier executable
**/
