package piscine

import (
	"github.com/01-edu/z01"
)

// func main() {
// piscine.QuadC(5,3)
//}

func QuadD(x, y int) {
	for b := 1; b <= y; b++ { // création de la boucle des colonnes

		for a := 1; a <= x; a++ { // création de la boucle des lignes

			if b == 1 { // création de la première ligne

				if a > 1 && a < x {
					z01.PrintRune(rune(66)) //print B
				}
				if a == 1 {
					z01.PrintRune(rune(65)) // Print A
				}
				if a > 1 {

					if a == x {
						z01.PrintRune(rune(67)) // Print C
					}
				}
				if a == x {
					z01.PrintRune('\n')
				}
			}

			if b > 1 && b < y { // création de la ligne du milieu

				if a > 1 && a < x {
					z01.PrintRune(' ') // print space
				}
				if a == 1 || a == x {
					z01.PrintRune(rune(66)) //print B
				}
				if a == x {
					z01.PrintRune('\n')
				}
			}
			if b == y { // création dernière ligne

				if y > 1 {

					if a > 1 && a < x {
						z01.PrintRune(rune(66)) // print B
					}
					if a == 1 {
						z01.PrintRune(rune(65)) // print A
					}
					if x > 1 {

						if a == x {
							z01.PrintRune(rune(67)) //print C
						}
					}
				}
			}
		}
	}
}
