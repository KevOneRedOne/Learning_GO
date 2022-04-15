package piscine

/**
import (
	piscine".."
	"fmt"
)

func main (){
	n := 0
	piscine.PointOne(&n)
	fmt.Println(n)
} **/

func PointOne(n *int) {
	*n = 1

}

/** Dans cet exercice sur les pointeurs dans le langage go, on veut écrire une fonction
qui prend un pointeur vers un int comme argument et donne à ce int la valeur 1.

Donc on doit utilier un pointeur comme un int et affecter à ce int la valeur de 1.**/
