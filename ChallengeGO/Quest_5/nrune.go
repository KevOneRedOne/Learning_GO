package piscine

/**
package main

import (
	"github.com/01-edu/z01"

	piscine ".."
)

func main() {
	z01.PrintRune(piscine.NRune("Hello!", 3))
	z01.PrintRune(piscine.NRune("Salut!", 2))
	z01.PrintRune(piscine.NRune("Bye!", -1))
	z01.PrintRune(piscine.NRune("Bye!", 5))
	z01.PrintRune(piscine.NRune("Ola!", 4))
	z01.PrintRune('\n')
}
**/

func NRune(s string, n int) rune {

	tab := []rune(s)
	counter := 0

	for range tab {

		counter++
	}

	if n > 0 && n <= counter {

		return tab[n-1]
	}
	return 0
}

//Ecrire une fonrcion quiu retourne la rune d'un string => soit la valeur de l'index
//Création d'un tableau ?
//Quand impossible retourner 0
