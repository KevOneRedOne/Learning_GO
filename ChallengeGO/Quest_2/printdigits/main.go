package main

import "github.com/01-edu/z01"

func main() {

	for a := 48; /**0**/ a < 58; /**:**/ a++ {

		z01.PrintRune(rune(a))

	}
	z01.PrintRune('\n')

}

// Dans cet exercice, on doit afficher sur la même ligne et dans l'ordre croissant les nombres de 0 à 9 sur la même ligne
