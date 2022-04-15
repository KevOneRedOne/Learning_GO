package piscine

/**Écrivez une fonction qui prend un int min et un int max comme paramètres.
Cette fonction renvoie une tranche de int avec toutes les valeurs comprises entre min et max.
Min est inclus et max est exclu.
Si min est supérieur ou égal à max, une tranche nulle est renvoyée.
make n'est pas autorisé pour cet exercice.**/

/**resultat :
[5 6 7 8 9]
[]
**/

func AppendRange(min, max int) []int {

	var tab []int

	for i := min; i < max; i++ {

		tab = append(tab, i)
	}
	return tab
}
