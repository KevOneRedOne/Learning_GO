package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {

	test("Hello World")
	test("hjompG!!")
	test("Kévin") // é character is not found inside standard.txt = Error
	test("aaΓει") // Non-ASCII character = Error
	test("1a\" #FdwHywR&/()=")
	test("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	test("2hello 1there")
	test("")
	test("\"#$%&/()*+,-./")
	test("This\\nisbackslash") // Due to the language specificity of Golang, it is required to add another backslash so this test can work.
	test("Tab\\tf")            // Same as above
}

func test(input string) {

	out, _ := exec.Command("go", "run", "main.go", input).Output()

	fmt.Println(string(out))

}
