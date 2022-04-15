package main

import (
	"os"

	"github.com/01-edu/z01"
)

type even struct { //Création de la structure
	isEven   int
	printStr string
	even     int
}

func printStr(s string) {

	for r := range s {

		z01.PrintRune(rune(s[r]))

	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {

	yes := true
	no := false

	if nbr == 1 {

		return yes

	} else {

		return no

	}
}

func main() {

	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	lengthOfArg := 0
	tab := os.Args[1:]

	for range tab {

		lengthOfArg++
	}

	if lengthOfArg == 1 {

		printStr(EvenMsg)

	} else {

		printStr(OddMsg)
	}
}
