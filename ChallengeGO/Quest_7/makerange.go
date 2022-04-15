package piscine

/**Écrire une fonction qui prend un int min et un int max comme paramètres.
Cette fonction renvoie une tranche de int avec toutes les valeurs comprises entre min et max.
Min est inclus et max est exclu.
Si min est supérieur ou égal à max, une tranche nulle est renvoyée.
append n'est pas autorisé pour cet exercice.**/

/** Resultat :
[5 6 7 8 9]
[]
**/

//on veut le meme resultat que pour appendrange.go sauf qu'il faut utiliser Make

func MakeRange(min, max int) []int {

	var tab []int

	if min >= max {

		return nil

	}

	tab = make([]int, max-min)

	for a := min; a < max; a++ {

		tab[a-min] = a
	}
	return tab
}

// for initialisation ; condition; itération {}
