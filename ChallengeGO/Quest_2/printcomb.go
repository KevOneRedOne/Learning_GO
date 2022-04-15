package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {

	for centaines := '0'; centaines <= '9'; centaines++ {

		for dizaines := '0'; dizaines <= '9'; dizaines++ {

			for unites := '0'; unites <= '9'; unites++ {

				if centaines < dizaines && dizaines < unites {

					z01.PrintRune(centaines)
					z01.PrintRune(dizaines)
					z01.PrintRune(unites)

					if centaines == '7' && dizaines == '8' && unites == '9' {

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

/**Ecrire un programme qui affiche sur une seule ligne la combinaison de 3 nombres, sachant que
le premier nombre est inferieur au second, le second < au troisieme.

les differentes combinaisons sont séparées par une virrgule et un espace

000 et 999 ne sont pas admin

la combinaison fini par 789**/
