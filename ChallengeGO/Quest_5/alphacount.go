package piscine

/**
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	s := "Hello 78 World!    4455 /"
	nb := piscine.AlphaCount(s)
	fmt.Println(nb)
}
**/

func AlphaCount(s string) int {

	counter := 0

	for _, index := range s {

		if (index >= 'A' && index <= 'Z') || (index >= 'a' && index <= 'z') {

			counter++

		}

	}
	return counter

}

/**
Ecrire une fonction qui compte les lettres d'un string et qui retourne le resultat.
Les caractères comptés sont uniquement des lettres latines, les espaces ou les autres caractères ne sont pas comptés.
**/
