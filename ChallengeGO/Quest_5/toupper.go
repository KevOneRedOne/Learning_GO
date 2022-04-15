package piscine

//on veut écrire une fonction qui transforme la chaine de caractère en majuscule

func ToUpper(s string) string {

	var lettre string

	for _, l := range s {

		if l >= 97 /**a**/ && l <= 122 /**z**/ {

			lettre += string(rune(l - 32))

		} else {
			lettre += string(l)
		}
	}
	return lettre
}

// += additionne et attribue le resultat à la variable

/** string(rune(l - 32)), car si l, qui est le rang de s, est superieur ou égale à 'a' et
inferieur ou égale à 'z', alors on enlève à l, 32 rang de la table Ascii (97-32 = 65, soit 'A').
Puis, avec string(l), on recommence à 'A' **/
