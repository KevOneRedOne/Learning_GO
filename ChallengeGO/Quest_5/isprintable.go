package piscine

/**
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	fmt.Println(piscine.IsPrintable("Hello"))
	fmt.Println(piscine.IsPrintable("Hello\n"))

}
**/

func IsPrintable(s string) bool {

	tab := []rune(s)

	for _, a := range tab {

		if a >= 0 && a < 32 { // on elimine les non-caractères de la table ascii

			return false
		}
	}
	return true
}

/**
Write a function that returns true if the string passed in parameter only contains printable characters,
 and that returns false otherwise.
**/
