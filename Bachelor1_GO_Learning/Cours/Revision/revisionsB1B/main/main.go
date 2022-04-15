package main

import (
	"fmt"
	"os"

	rev "../revision"
)

func main() {

	programName := os.Args[0]
	args := os.Args[1:]

	fmt.Println(programName)
	fmt.Println(args)
	rev.StrLen("12345")

}
