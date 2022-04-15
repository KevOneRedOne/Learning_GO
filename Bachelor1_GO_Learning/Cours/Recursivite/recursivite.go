package recursivite

// Fonction récursive : Une fonction qui s'appelle elle même
// Relation de récurrence : Une relation/expression qui permet de lier un objet de taille (n-1) à l'objet de même type de taille n
// U_0 = ...
// U_(n+1) = f(Un)

// sum(4) = 1 + 2 + 3 + 4 = 10
// sum(5) = (1 + 2 + 3 + 4) + 5

// sum(5) = sum(4) + 5
// sum(5) = sum(5-1) + 5
// ==> 5 --> n
// sum(n) = sum(n-1) + n

// RecursiveSum(n) = RecursiveSum(n-1) + n

// RecursiveSum est une fonction récursive qui calcule la somme des entiers de 1 à n
func RecursiveSum(n int) int {

	somme := -1
	// Cas de base | Cas initial
	if n == 1 {
		somme = 1 // RecursiveSum(1) = 1
	} else if n > 1 {
		somme = n + RecursiveSum(n-1) // RecursiveSum(n) = RecursiceSum(n-1) + n
	}

	return somme
}

// RecursiveFactorial(4) = 1 * 2 * 3 * 4
// RecursiveFactorial(5) = 1 * 2 * 3 * 4 * 5
//
// RecursiveFactorial(n) = RecursiveFactorial(n-1) * n

// RecursiveFactorial est une fonction récursive qui calcule la factorielle d'un nombre
func RecursiveFactorial(n int) int {

	factorielle := 1
	limit := 20

	if n == 0 || n == 1 {
		factorielle = 1
	} else if n > 1 && n <= limit {
		factorielle = RecursiveFactorial(n-1) * n // RecursiveFactorial(n) = RecursiveFactorial(n-1) * n
	}
	return factorielle
}

// RecursivePower(5, 2) = 5 * 5
// RecursivePower(5, 3) = RecursivePower(5,2) * 5
// RecursivePower(5, 3) = RecursivePower(5,3-1) * 5
// RecursivePower(nb, power) = RecursivePower(nb,power-1) * nb
// RecursivePower est une fonction récursive qui calcule la puissance d'un nombre
// PRECONDITION : nb > 0, power >= 0
// POSTCONDITION : RecursivePower(nb, power) >= 0
func RecursivePower(nb int, power int) int {

	result := -1

	if power == 0 {
		result = 1
	} else if power > 0 {
		result = RecursivePower(nb, power-1) * nb // RecursivePower(nb, power) = RecursivePower(nb,power-1) * nb
	}
	return result
}

// "abcde" = 'a' + "bcde"
// "bcde" = 'b' + "cde"
// "cde" = 'c' + "de"
// "de" = 'd' + "e"
// "e" = 'e' + ""
// RecursiveStrLen("") ==> 0

// RecursiveStrLen("abcde") = 1 + RecursiveStrLen("bcde")

// RecursiveStrLen est une fonction récursive qui calcule la taille d'un slice de rune
func RecursiveStrLen(runes []rune) int {

	length := -1

	if string(runes) == "" {
		length = 0
	} else {
		length = 1 + RecursiveStrLen(runes[1:])
	}

	return length
}

// // RecursiveStrLen est une fonction récursive qui calcule la taille d'une chaine de caractere
// func RecursiveStrLen(s string) int {

// 	length := -1
// 	runes := []rune(s)

// 	if s == "" {
// 		length = 0
// 	} else {
// 		length = 1 + RecursiveStrLen(string((runes[1:])))
// 	}

// 	return length

// }
