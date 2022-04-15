package main

import (
	"fmt"

	Index ".."
)

func main() {

	fmt.Println(Index.Index("Hello!", "l"))   // 2
	fmt.Println(Index.Index("Salut!", "alu")) // 1
	fmt.Println(Index.Index("Ola!", "hOl"))   //-1
	fmt.Println(Index.Index("Coucou", "u"))   // 2
	fmt.Println(Index.Index("Ola!", "hOla"))  //

}
