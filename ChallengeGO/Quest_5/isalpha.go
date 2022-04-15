package piscine

/**
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	fmt.Println(piscine.IsAlpha("Hello! How are you?"))
	fmt.Println(piscine.IsAlpha("HelloHowareyou"))
	fmt.Println(piscine.IsAlpha("What's this 4?"))
	fmt.Println(piscine.IsAlpha("Whatsthis4"))

}
**/

func IsAlpha(s string) bool {

	tab := []rune(s)

	for _, a := range tab {

		if (a < 48 /**0**/) || (a > 57 /**9**/ && a < 65 /**A**/) || (a > 90 /**Z**/ && a < 97 /**a**/) || (a > 122 /**z**/) {

			return false
		}
	}
	return true

}

/**
Write a function that returns true if the string passed in parameter only contains alphanumerical characters
or is empty, and that returns false otherwise.
**/
