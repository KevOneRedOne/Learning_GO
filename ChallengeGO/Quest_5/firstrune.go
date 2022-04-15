package piscine

/**
package main

import (
	"github.com/01-edu/z01"

	piscine ".."
)

func main() {
	z01.PrintRune(piscine.FirstRune("Hello!"))
	z01.PrintRune(piscine.FirstRune("Salut!"))
	z01.PrintRune(piscine.FirstRune("Ola!"))
	z01.PrintRune('\n')
}
**/

func FirstRune(s string) rune {

	tab := []rune(s)

	return tab[0]

}

// Ecrire une fonction qui affice la première rune d'un string
