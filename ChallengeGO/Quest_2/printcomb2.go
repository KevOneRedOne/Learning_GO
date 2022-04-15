package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {

	for dizaines1 := '0'; dizaines1 <= '9'; dizaines1++ { // création de la 1ere boucle = 1er chiffre

		for unites1 := '0'; unites1 <= '9'; unites1++ { // création de la 1ere boucle = 2eme chiffre

			for dizaines2 := '0'; dizaines2 <= '9'; dizaines2++ { // création de la 2eme boucle = 1er chiffre

				for unites2 := '0'; unites2 <= '9'; unites2++ { // création de la 2eme boucle = 2eme chiffre

					if (dizaines1 < dizaines2) || (dizaines1 == dizaines2 && unites1 < unites2) {

						z01.PrintRune(dizaines1)
						z01.PrintRune(unites1)
						z01.PrintRune(' ')
						z01.PrintRune(dizaines2)
						z01.PrintRune(unites2)

						if dizaines1 == '9' && unites1 == '8' && dizaines2 == '9' && unites2 == '9' {

							z01.PrintRune('\n')

						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')

						}
					}
				}
			}
		}
	}
}

/** On veut écrire une fonction qui affiche dans un ordre croissant tous les nombres possibles
de deux nombres à deux chiffres.

Ces combinaisons sont séparés par un espace et finissent par une virgule.

Doit finir par 98 99
A la fin d'une combinaison, 00 99, elle doit redemarrer à 01 02**/
