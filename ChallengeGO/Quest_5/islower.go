package piscine

/**
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	fmt.Println(piscine.IsLower("hello"))
	fmt.Println(piscine.IsLower("hello!"))

}
**/

func IsLower(s string) bool {

	tab := []rune(s)

	for _, i := range tab {

		if i < 97 || i > 122 {

			return false
		}
	}
	return true

}

/** Ecrire une fonction qui affiche "vrai" si le string en paramètre ne contient que des minuscules, et faux dans les cas contraires**/
