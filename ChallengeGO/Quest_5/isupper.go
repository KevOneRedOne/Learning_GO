package piscine

/**
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	fmt.Println(piscine.IsUpper("HELLO"))
	fmt.Println(piscine.IsUpper("HELLO!"))
}
**/

func IsUpper(s string) bool {

	tab := []rune(s)

	for _, i := range tab {

		if i < 65 || i > 90 { // Voir la table AScii

			return false
		}
	}
	return true

}

/** Ecrire une fonction qui affiche "vrai" si le string en paramètre ne contient que des majuscules, et faux dans les cas contraires**/
