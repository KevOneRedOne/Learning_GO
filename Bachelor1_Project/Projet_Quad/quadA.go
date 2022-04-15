package piscine

import "github.com/01-edu/z01"

//package main

//import piscine ".."

//func main() {
//	piscine.QuadA(5,3)
//}

//QuadA is..
func QuadA(x, y int) {

	for j := 1; j <= y; j++ {

		for i := 1; i <= x; i++ {

			if j > 1 && j < y {

				if i > 1 && i < x {

					z01.PrintRune(' ')

				}

				if i == 1 || i == x {

					z01.PrintRune('|')

				}

				if i == x {

					z01.PrintRune('\n')
				}

			}

			if j == 1 || j == y {

				if i > 1 && i < x {

					z01.PrintRune('-')

				}

				if i == 1 || i == x {

					z01.PrintRune('o')

					if i == x {

						z01.PrintRune('\n')
					}
				}

			}

		}
	}

}

// for initialisation ; condition ; itération {
/* le code qui sera répété */
