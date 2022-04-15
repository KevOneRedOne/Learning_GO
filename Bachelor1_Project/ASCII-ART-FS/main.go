package main

import (
	"bufio"
	"fmt"
	"os"
)

//ReadFile returns an array of string which is the same as the file (line = line)
func ReadFile(StylizedFile string) []string {
	var source []string
	file, _ := os.Open(StylizedFile)  // opens the .txt
	scanner := bufio.NewScanner(file) // scanner scans the file
	scanner.Split(bufio.ScanLines)    // set-up scanner preference to read the file line-by-line
	for scanner.Scan() {              // loop that performs a line-by-line scan on each new iteration
		if scanner.Text() != "" {
			source = append(source, scanner.Text()) // adds the value of scanner (that contains the characters from StylizedFile) to source
		}
	}
	file.Close() // closes the file
	return source
}

//Affichage prints the stylized characters
func Affichage(arguments []rune, txtfile []string) {
	for ligne := 0; ligne < 8; ligne++ {
		for index, char := range arguments {
			fmt.Print(txtfile[ligne+(int(char)-32)*8])
			if index == len(arguments)-1 && ligne != 7 {
				fmt.Println()
			}
		}
	}
}

func main() {
	var filepath string    //value which stocks the path of the file
	if len(os.Args) == 3 { // lenght of arguments is os.Args[0,1,2] = 3 => "main.go" "word" "font-style"
		filepath = os.Args[2]
		if filepath == "standard" || filepath == "shadow" || filepath == "thinkertoy" {
			filepath = filepath + ".txt"
		} else {
			fmt.Println("Error")
			os.Exit(0) //Stops running the program
		}
	} else if len(os.Args) == 2 { // lenght of arguments is os.Args[0,1] = 2 => "main.go" "word"
		filepath = "standard.txt"
	} else { // lenght of arguments is os.Args[0] = 1 => "main.go"
		fmt.Println("Error")
		os.Exit(0) //Stops running the program
	}
	txtfile := ReadFile(filepath)
	arguments := []rune(os.Args[1])
	for index := range arguments {
		if arguments[index] < 32 || arguments[index] > 126 { // Ascii-table
			fmt.Println("Error")
			os.Exit(0) // Stops running the program
		}
	}
	var start int
	for index := range arguments {
		if arguments[index] == '\\' && index != len(arguments)-1 { // condition to check if at the index of "arguments" the value is strictly equal to '\\' (which means a simple `\`)
			// and if the index is different from the length of "arguments"
			if arguments[index+1] == 'n' { // condition to verify if the character following '\' is a 'n'
				Affichage(arguments[start:index], txtfile)
				fmt.Println()
				start = index + 2
			}
		} else if index == len(arguments)-1 {
			Affichage(arguments[start:], txtfile)
		}
	}
}
