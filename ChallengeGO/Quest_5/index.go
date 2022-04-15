package piscine

/**
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	fmt.Println(piscine.Index("Hello!", "l")) // 2
	fmt.Println(piscine.Index("Salut!", "alu"))// 1
	fmt.Println(piscine.Index("Ola!", "hOl")) //-1
}
**/

func Index(s string, toFind string) int {

	lentghS := StrLen(s)
	lentghtofind := StrLen(toFind)

	for i := 0; i < lentghS-lentghtofind; i++ {

		if s[i:i+lentghtofind] == toFind {

			return i

		}
	}
	return -1
}

/**
On veut ecrire une fonction qui fasse la meme chose que la fonction Index
c'est à dire une fonction qui retourne l'index du premier caractere de s present dans tofind

StrLen = fonction qui compte le nombre de caractère dans une chaine de caractère

**/
