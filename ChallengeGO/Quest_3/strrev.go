package piscine

/**
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	s := "Hello World!"
	s = piscine.StrRev(s)
	fmt.Println(s)
}
**/

func StrRev(s string) string {

	lettre := 0 //création d'une variable qui va compter

	for range s {

		lettre++
	}
	lettre--

	tab := []byte(s) //création d'un tableau , conversion s en byte

	for index := range s {

		tab[index] = s[lettre-index]

	}
	revers := string(tab) // reconvertir tab en string

	return revers

}

/** Ecrire une fonction qui inverse les caractères d'un string, soit
"!dlroW olleH"
La fonction retourne le s string**/

/**Autre Méthode :

func StrRev(s string) string {

	runes := []rune(s)
	len := Strlen(s)

	for i, j := 0, len-1; i < len/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
**/
