package piscine

/**
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	arg := 4
	fmt.Println(piscine.IterativeFactorial(arg))
}
**/

func IterativeFactorial(nb int) int {

	if nb < 0 || nb > 20 { //on supprime les valeurs impossibles : les negatives et les nombres supérieurs à 64 bits

		return 0

	}

	if nb == 0 { //La fontion factorielle de 0 est 1 par convention

		return 1

	}
	resultat := 1 //création d'une variable int à 1

	for a := 1; a <= nb; a++ { // création de la boucle itérative
		// si a est inferieur ou égale à nb alors l'itération commence
		resultat = resultat * a
	}
	return resultat
}

/**
¤ Ecrire une fonction iterative qui retourne le facteur du int déclaré

¤ si une Erreur retourne la valeur 0
¤ Avec boucle For + If


En gros, il faut ecrire une fonction qui te permet d'afficher l'expression factorielle
d'un nombre ; soit :

4! = 4 x 3 x 2 x 1 = 24

Et s'il y a une erreur, alors retourner la valeur 0


// une fonction factorielle (symbol : !) sert a multiplier tous les nombres précedents le
nombre choisi (inclu aussi) :

1! = 1
7! = 7 x 6 x 5 x 4 x 3 x 2 x 1 = 5040
**/
