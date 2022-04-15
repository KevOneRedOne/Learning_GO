package piscine

/** on veut écrire la meme fonction que toupper.go sauf que la on veut changer la chaine de caractère
en minuscule**/

func ToLower(s string) string {

	var lettre string // création d'une variable string qui va stocker le resultat

	for _, l := range s {

		if l >= 65 /**'A'**/ && l <= 90 /**'Z'**/ {

			lettre += string(rune(l + 32 /**32 rang dans la table Ascii**/))

		} else {

			lettre += string(l)
		}

	}
	return lettre

}

// += additionne et attribue le résultat à la variable
