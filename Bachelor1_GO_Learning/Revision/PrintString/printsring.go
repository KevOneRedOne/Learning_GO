package printstring

import "github.com/01-edu/z01"

//PrintString is a program that runs through a chain of characters and print it
func PrintString(s string) {

	tab := []rune(s)

	for element := range tab {

		z01.PrintRune(rune(tab[element]))

	}
}
