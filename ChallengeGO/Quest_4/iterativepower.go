package piscine

/**
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	fmt.Println(piscine.IterativePower(4, 3))
}
**/

func IterativePower(nb int, power int) int {

	if power < 0 || power > 20 { // on élimine les valeurs impossibles à print

		return 0

	}

	if power == 0 { // On a fonction factorielle de 0 égale 1 par convention

		return 1
	}

	result := 1

	for a := 1; a <= power; a++ {

		result = result * nb

	}
	return result

}

/**
Ecrire une fonction itérative qui renvoie le nombre "power" de l'int passé en paramètre.
Les nombres de "power" négatives renverront 0. Les débordements ne doivent pas être traités.
**/
