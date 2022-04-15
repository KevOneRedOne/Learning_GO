package revision

func StrLen(s string) int {
	length := 0

	for range []rune(s) {
		length++
	}

	return length
}

func Index(texte string, mot string) int {

	// 2- Calcul de la taille du mot que je cherche
	lengthMot := StrLen(mot)
	lengthTexte := StrLen(texte)

	if lengthMot == 0 && lengthTexte == 0 {
		return 0
	} else {
		// 3- Parcourir le texte en formant des mots de la même taille que le mot que l'on cherche
		for i := 0; i < lengthTexte-lengthMot; i++ {
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

func PrintStr(s string) {
	// 1 - Convertir la chaine de caractère en slice de rune
	// 2 - Parcourir le slice en utilisant range
	// 3 - Affiche chacun des caractère avec PrintRune
}

func MaxTab(tab []int) int {

	if len(tab) == 0 {
		println("Le tableau est vide")
		return -1
	} else {
		max := tab[0]

		for i := 0; i < len(tab); i++ {
			if tab[i] > max {
				max = tab[i]
			}
		}

		return max

	}

}

func MinTab(tab []int) int {

	if len(tab) == 0 {
		println("Le tableau est vide")
		return -1
	} else {
		min := tab[0]
		for elt := range tab {
			if elt < min {
				min = elt
			}
		}

		return min
	}
}
