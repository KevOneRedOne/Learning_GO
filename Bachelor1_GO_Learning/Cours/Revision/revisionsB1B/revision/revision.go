package revision

//
func StrLen(s string) int {

	length := 0
	for range s {
		length++
	}

	return length
}

// Affiche une chaine de caractère en utilisant PrintRune
func PrintStr(s string) {

}

func FirstRune(s string) {

}

func LastRune(s string) {

}

func Compare(a, b string) int {

	// On regarde la première lettre de chacun des mots
	// ...

	return 0
}

func Index(texte string, mot string) int {

	// 2- Calcul de la taille du mot que je cherche
	lengthMot := StrLen(mot)
	lengthTexte := StrLen(texte)

	if StrLen(mot) > StrLen(texte) {
		return -1
	} else if mot == texte {
		return 0
	} else {
		// 3- Parcourir le texte en formant des mots de la même taille que le mot que l'on cherche
		for i := 0; i <= lengthTexte-lengthMot; i++ {
			// 3.1 Formation d'un mot de la même taille que le mot que l'on cherche
			candidat := texte[i:(i + lengthMot)]
			// 3.2 Comparaison des deux mots
			if candidat == mot {
				return i
			}
		}

		return -1
	}
}

// Renvoie le minimum d'un tableau d'entiers
func MinTab(tab []int) int {

	return 0
}

// Renvoie le maximum d'un tableau d'entiers
func MaxTab(tab []int) int {

	return 0
}
