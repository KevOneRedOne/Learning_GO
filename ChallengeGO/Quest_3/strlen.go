package piscine

/**
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	l := piscine.StrLen("Hello World!")
	fmt.Println(l)
}
**/

func StrLen(s string) int {

	var l int = 0

	for range s {

		l++

	}
	return l
}

/**Ecrire une fonction qui affiche le nombre de rune dans un string. Donc
une fonction qui affiche le nombre de caractère dans une chaine de
caractère**/
