package piscine

import "github.com/01-edu/z01"

func QuadB(x, y int) {

	for j := 1; j <= y; j++ { // création de la boucle y soit les colonnes

		for i := 1; i <= x; i++ { // création de la boucle x soit les lignes

			if j > 1 && j < y { // Création de la ligne du milieu

				if i > 1 && i < x {

					z01.PrintRune(' ')

				}

				if i == 1 || i == x {

					z01.PrintRune('*')

				}

				if i == x {

					z01.PrintRune('\n')
				}

			}

			if j == 1 { // Création ligne 1

				if i > 1 && i < x {

					z01.PrintRune('*')

				}

				if i == 1 {

					z01.PrintRune('/')

				}

				if i > 1 { // bloquer la premiere ligne pour pas la recommencer

					if i == x {

						z01.PrintRune(rune(92))

					}
				}

				if i == x {

					z01.PrintRune('\n')
				}

			}

			if j == y { // Création de la derniere ligne modifié

				if y > 1 { // Bloquer la derniere ligne pour pas la recommencer

					if i == 1 {

						z01.PrintRune(rune(92))

					}

					if i > 1 && i < x {

						z01.PrintRune('*')

					}

					if i > 1 {

						if i == x {

							z01.PrintRune('/')

						}

					}

				}

			}

		}
	}

}

// for initialisation ; condition ; itération {
/* le code qui sera répété */
