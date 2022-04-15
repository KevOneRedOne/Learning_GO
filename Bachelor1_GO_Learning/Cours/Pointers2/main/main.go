package main

import pointers ".."

// Scope - Portée des variables
func main() {

	a := 10
	b := 15
	println(a)
	println(b)
	println()
	pointers.Swap(&a, &b)
	println(a)
	println(b)
}
