package main

import (
	"fmt"
	"os"
)

func main() {

	// []int

	// 1- []type en tant que type de slice
	// var runes []rune

	// 2- En tant que constructeur (Construire un objet de type slice)
	// []rune{ , , , }
	// []rune()

	//s := "Paris Ynov Campus"

	// var runes []rune = []rune{'a', 'b', 'c'}
	// fmt.Println(runes)

	// sRunes := []rune(s)
	// fmt.Println(sRunes)

	// fmt.Println(revision.Index("Je m'appelle Alexis.", "Florian"))
	// fmt.Println(revision.Index("Je m'appelle Alexis.", "appelle"))
	// fmt.Println(revision.Index("Je m'appelle Alexis.", "Alexis"))

	// os.Args est de type []string

	// PrintProgramName
	fmt.Println(os.Args[0])

	// PrintParams
	fmt.Println(os.Args[1:])

}
