package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	for index, elem := range os.Args {

		if index != 0 {

			for _, word := range elem {

				z01.PrintRune(word)
			}
			z01.PrintRune('\n')

		}

	}

}
