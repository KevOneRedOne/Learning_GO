package main

import "github.com/01-edu/z01"

func main() {

	for l := 122; /**z**/ l > 96; /**rang avant lettre a**/ l-- {

		z01.PrintRune(rune(l))
	}
	z01.PrintRune('\n')
}

// Dans cet exercice, on veut afficher l'alphabet à l'envers
