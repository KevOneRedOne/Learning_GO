package piscine

/**
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	a := 0
	b := 1
	piscine.Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
**/

func Swap(a *int, b *int) {

	var c int = *b //création d'une variable qui stocke la donnée du pointeur

	*b = *a //Transfert contenu de b vers a

	*a = c //Transfert contenu de a vers c, soit b

}

/** Ecrire une fonction qui échange le contenu de deux pointeurs**/
