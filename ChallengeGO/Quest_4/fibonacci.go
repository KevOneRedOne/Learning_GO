package piscine

/**
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	arg1 := 4
	fmt.Println(piscine.Fibonacci(arg1))
}
**/

func Fibonacci(index int) int {

	if index < 0 { // on élimine les valeurs négatives

		return -1

	}
	if index == 0 { //La première variable commence à index 0

		return 0

	}

	if index == 1 { //La deuxième variable commence à index 1

		return 1

	}

	return Fibonacci(index-2) + Fibonacci(index-1)
}

/**
Ecrire une fonction recursive qui retourne la valeur de la séquence
fibonacci correspondant à l'index passé en paramètre.

la premiere variable est à l'index 0

la sequence commence par : 0, 1, 1, 2, 3 etc

Un index negatif retournera un -1

La fonction For est interdite.


¤ La suite de Fibonacci commence par des premiers termes, à savoir 0 et 1,
ensuite chaque terme successif est la somme des deux termes précédents ; soit :

0+1=1, 1+1=2, 1+2=3, 2+3=5, 3+5=8 etc..

la suite de Fibonacci peut etre defini avec les formules suivantes :
--> F0 = 0
--> F1 = 1
--> Fn = Fn-2 + Fn-1 pour n >= 2
**/
