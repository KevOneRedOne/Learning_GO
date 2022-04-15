package piscine

/**
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	fmt.Println(piscine.Compare("Hello!", "Hello!"))
	fmt.Println(piscine.Compare("Salut!", "lut!"))
	fmt.Println(piscine.Compare("Ola!", "Ol"))
}

Ecrire une fonction qui compare deux strings ensembles
**/

/**resultat:
0
-1
1
**/

func Compare(a, b string) int {

	if a == b {
		return 0
	}

	if a > b {
		return 1
	}

	if a < b {
		return -1
	}
	return 0
}

/**
Compare renvoie un entier comparant deux chaînes lexicographiquement.
Le résultat sera 0 si a == b, -1 si a < b et +1 si a > b.
Compare est inclus uniquement pour la symétrie avec les octets du package.
Il est généralement plus clair et toujours plus rapide d'utiliser les opérateurs
de comparaison de chaînes intégrés ==, <,>, etc. **/
