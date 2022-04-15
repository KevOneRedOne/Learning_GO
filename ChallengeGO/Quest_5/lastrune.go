package piscine

/**
import (
	"github.com/01-edu/z01"
	piscine "..")

package main() {

z01.PrintRune(piscine.LastRune("Hello!"))
z01.PrintRune(piscine.LastRune("Salut!"))
z01.PrintRune(piscine.LasstRune("Ola!"))
z01.PrintRune('\n')
}

resultat : !!!
**/

func LastRune(s string) rune {

	var lettre rune //décalaration de la variable lettre en rune

	for _, place := range s { // _, permet de ne pas déclarer toutes les variables

		lettre = place //le mot-clé range retourne à la fois l'index et la valeur de l'element itéré du tableau

	}
	return lettre
}

/** Le _, est un espace réservé qui signifie essentiellement
«Je ne me soucie pas de cette valeur de retour particulière».
**/

/** Autre méthode :

func LastRune(s string) rune {

	tab := []rune(s)
	a := 0

	for range tab {
		a++
	}
	return tab[a-1]

}
**/
