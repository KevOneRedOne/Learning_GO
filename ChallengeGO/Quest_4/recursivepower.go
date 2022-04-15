package piscine

/**
package main

import (
        "fmt"
        piscine ".."
)

func main() {
	fmt.Println(piscine.RecursivePower(4, 3))
}
**/

func RecursivePower(nb int, power int) int {

	if power < 0 { //on élimine les valeurs négatives

		return 0

	}

	if power == 0 {

		return 1

	}
	return nb * RecursivePower(nb, power-1) // on multiplie la valeur rentrée dans nb par la fonction recursive avec power - 1

}

/**
Ecrire une fonction récursive qui renvoie power de l'int passé en paramètre.
Les puissances négatives renverront 0. Les débordements ne doivent pas être traités.
La boucle For est interdit pour cet exercice.
**/
