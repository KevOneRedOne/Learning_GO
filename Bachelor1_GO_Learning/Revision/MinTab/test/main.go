package main

import (
	"fmt"

	mintab ".."
)

func main() {

	tab := []int{5, 6, 7, 7, 8}
	// tab = [] int {1,4,5,0,7}

	min := mintab.MinTab(tab)
	fmt.Println(min)
}
