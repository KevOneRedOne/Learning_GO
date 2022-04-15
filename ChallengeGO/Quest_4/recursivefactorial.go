package piscine

/**
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	arg := 4
	fmt.Println(piscine.RecursiveFactorial(arg))
}
**/

func RecursiveFactorial(nb int) int {

	if nb < 0 || nb > 20 { //on élimine les valeurs impossibles comme dans itérativefactorial

		return 0

	}

	if nb == 0 || nb == 1 { //Par convention, le factorielle de 0 est 1, et celui de 1 est 1

		return 1

	}

	return nb * RecursiveFactorial(nb-1) // on multiple nb par le nb en paramètre - 1
}

/**
Recursivefactorial :

Ecrire une fonction recursive qui retourne le nombre factoriel du int parametré.

Si erreur, alors retourne 0.

la boucle for est interdite dans la fonction.


valeur à supprimer  = en dessous de 0; au dessus de 20
à la place retourne 0
**/
