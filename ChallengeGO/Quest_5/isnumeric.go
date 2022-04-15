package piscine

/**Écrire une fonction qui renvoie true si la chaîne passée en paramètre ne contient que
des caractères numériques, et qui renvoie false dans le cas contraire.**/

func IsNumeric(s string) bool {

	tableau := []rune(s)

	for _, a := range tableau {

		if a < 48 || a > 58 {

			return false

		}
	}
	return true
}
