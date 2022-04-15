package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	lenght := 0

	for range os.Args {

		lenght++
	}

	for elem := range os.Args {

		if elem != lenght-1 {

			for _, word := range os.Args[lenght-1-elem] {

				z01.PrintRune(word)

			}
			z01.PrintRune('\n')
		}

	}
}
