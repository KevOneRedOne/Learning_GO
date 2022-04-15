package main

import "github.com/01-edu/z01"

func main() {

	for l := 65; l < 91; l++ { // je definie l comme la valeur dans le tableau ASCII correspondant à la lettre a

		z01.PrintRune(rune(l)) // je transforme l en rune pour pouvoir print l

	}
	z01.PrintRune('\n')

}

// Dans cet exo, on nous demande d'afficher l'alphabet avec notre programme
