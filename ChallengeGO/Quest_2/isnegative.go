package piscine

import (
	"github.com/01-edu/z01"
)

func IsNegative(nb int) {

	if nb >= 0 {

		z01.PrintRune('F')

	} else {

		z01.PrintRune('T')

	}
	z01.PrintRune('\n')
	/**if nb == -1 {

		z01.PrintRune(84)
		z01.PrintRune('\n')

	}
	if nb == 0 {

		z01.PrintRune(70)
		z01.PrintRune('\n')

	}
	if nb == 1 {

		z01.PrintRune(70)
		z01.PrintRune('\n')
	} 1er Code pas bon car il n'affichait pas True pour toutes les valeurs négatives**/

}

/** Dans cet exercice on veut écrire une fonction qui affiche si 'T' (true) si la variable
dans les paramètres est négative, et 'F' si ça n'est pas le cas**/
