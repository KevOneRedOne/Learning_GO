package conditions

// Opérateurs booléens
// > < >= <= == !=

var entier int = 15

func EstPositif(entier int) {

	if entier >= 0 {
		print("L'entier est positif !")
	} else { // Dans tous les autres cas
		print("L'entier est négatif !")
	}

}

func EstPair(entier int) {

	if entier%2 == 0 {
		print("L'entier est pair !")
	} else {
		print("L'entier est impair !")
	}

}

// if ... else if ... else ...
func Compare(a int, b int) {

	if a < b {
		print("a inférieur strictement à b")
	} else if a > b {
		print("a supérieur strictement à b")
	} else { // a == b
		print("a est égal à b")
	}

}

// Opérateur logique et : &&
func IsDigit(char rune) bool {
	if char >= 48 && char <= 57 {
		return true
	} else {
		return false
	}

}

// Opérateur logique et : &&
func IsLower(char rune) bool {
	if char >= 97 && char <= 122 {
		return true
	} else {
		return false
	}
}

// Opérateur logique et : &&
func IsUpper(char rune) bool {
	if char >= 65 && char <= 90 {
		return true
	} else {
		return false
	}
}

// Opérateurs logiques ou : ||
func IsAlphaNum(char rune) bool {
	if IsDigit(char) || IsLower(char) || IsUpper(char) {
		return true
	} else {
		return false
	}

}

// Exercice: Ecrire la fonction Sort qui compare les entiers a, b et c et les affiche dans l'ordre croissant
// en les séparant par des espaces
// Sort(5,3,1) -> 1 3 5
// Sort(1,7,4) -> 1 4 7
// Sort(1,2,3) -> 1 2 3

func Sort(a int, b int, c int) {

}
