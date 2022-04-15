package piscine

import "github.com/01-edu/z01"

/**
package main

import (
	"github.com/01-edu/z01"
	piscine ".."
)
func main() {
	piscine.PrintStr("Hello World!")
}
**/

func PrintStr(s string) {

	for length := range s {

		z01.PrintRune(rune(s[length]))

	}

}

//Ecrire une fonction qui affiche un par un les caractères de s string.
