package revision

func MinTab(tab []int) int {

	return 0
}

func MaxTab(tab []int) int {

	return 0
}

func StrLen(s string) int {
	length := 0
	for range s {
		length++
	}
	return length
}

func Compare(a, b string) int {

	// On regarde la première lettre de chacun des mots
	// ...

	return 0
}

func PrintStr(s string) {

}

// texte:= "Je m'appelle Florian et je suis etudiant au Paris Ynov Campus."
// mot := "Patrick"

func Index(texte string, mot string) int {

	lengthTexte := len(texte)
	lengthMot := len(mot)

	if len(mot) > len(texte) {
		return -1
	} else {
		// Parcourir le texte
		for index := 0; index <= lengthTexte-lengthMot; index++ {
			// Forme un mot avec le même nombre de caractère que le mot que l'on cherche
			candidat := texte[index:(index + lengthMot)]
			// Compare le mot formé avec le mot que l'on cherche
			if candidat == mot {
				return index
			}
		}
		return -1
	}
}
