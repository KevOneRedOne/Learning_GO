package loop

import (
	"github.com/01-edu/z01"
)

// Principe : Une boucle for permet de répéter un block d'instructions un nombre de fois déterminé
// Syntaxe :
// for i := 0; i < 10; i++ {
// 	Block d'instructions
// }

// PrintAlphabetLower is a function that prints the latin alphabet letters in lowercase
func PrintAlphabetLower() {
	for dec := 97; dec <= 122; dec++ {
		z01.PrintRune(rune(dec))
	}
	// for char := 'a'; char <= 'z'; char++ {
	// 	z01.PrintRune(char)
	// }
}

// // PrintAlphabetUpper is a function that prints the latin alphabet letters in uppercase
// func PrintAlphabetUpper() {
// }

// // PrintDigits is a function that prints the digits from 0 to 9
// func PrintDigits() {
// }

// // PrintReverseDigits is a function that prints the digits from 9 to 0
// func PrintReverseDigits() {
// }

// // Sum is a function that computes the sum of the integers from 1 to n
// // 1 + 2 + 3 + 4 + ... + n
// func Sum(n int) int {
// }

// // Factorial is a function that computes the product of the integers from 1 to n
// // 1 * 2 * 3 * 4 * ... * n
// func Factorial(n int) int {
// }

// NumberVowels
