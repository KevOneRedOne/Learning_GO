package main

import (
	"fmt"

	maxtab ".."
)

func main() {

	tab := []int{50, 6, 8, 100}

	max := maxtab.MaxTab(tab)

	fmt.Println(max)
}
