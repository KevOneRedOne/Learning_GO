package main

import (
	"fmt"
	"os"
)

func main() {

	if verifArgs() == true { //creation of a condition that checks if the function verifArgs return true to start the sudoku

		sudoku := creationArray()       //creation of the variable sudoku which will store the function creationArray
		sudoku = completeSudoku(sudoku) // modification of the variable sudoku by storing a new function : completeSudoku

		display(sudoku) //call of the function which display the sudoku grid

	}

}

func creationArray() [9][9]int {
	/** Creation of a function that return a double array of 9 row and 9 column.
	This function is use to collect the arguments of the program and store those arguments in a double array. **/

	arguments := os.Args[1:]
	/** os.Args collect and store in an array the arguments of the program.
	We recover with the index 1 which is the first argument , because index 0 is the name program **/

	var sudoku [9][9]int // Double array creation

	for a, elem := range arguments { // First loop for split the arguments
		// a is the index of the split of arguments

		for b, char := range elem { // Second loop for split every characters of the arguments
			// b is the index of the split of elem

			if char >= '0' && char <= '9' { // take only the number

				sudoku[a][b] = int(char) - 48
				/** Stored in the double array the character of every argument of the program
				We need to convert the characters, which is a rune type, to an int type. ( "int(char)")
				Besides, we removed the first 48 characters of the Ascii Table.**/
			}
		}
	}
	return sudoku //then, the function return the array

}

func display(sudoku [9][9]int) {
	/** Cration of the main function of the program whichi display the sudoku grid.
	This function have in parameter the variable sudoku of the function creationArray and print the sudoku grid.**/

	for row := 0; row < 9; row++ { //First loop that count the row of the sudoku

		for column := 0; column < 9; column++ { // Second loop that count the column of the sudoku

			fmt.Print(sudoku[row][column]) // fmt.Print allow to print somethings without a line break
			fmt.Print(" ")

		}
		fmt.Println() // fmt.Println() allow to print a line break
	}
}

func creationSquare(sudoku [9][9]int, row int, column int) [3][3]int {
	/** The function will create the 9 little blocks in the sudoku.
	This function have in parameter the sudoku of the function creationArray, and 2 int : row and column. **/

	var square [3][3]int // Creation of an double array of 3 rows and columns

	squareRow := 0 // Creation of 2 variables that will be used to delimit each block
	squareColumn := 0

	// Start the delimitation of the bloc in the lines
	if row >= 0 && row < 3 {

		squareRow = 0 // this is the name of the block in the lines

	}
	if row >= 3 && row < 6 {

		squareRow = 1
	}
	if row >= 6 && row < 9 {

		squareRow = 2
	}

	// We do the same to delimit the columns of the block.
	if column >= 0 && column < 3 {

		squareColumn = 0 //this is the name of the block in the columns
	}
	if column >= 3 && column < 6 {

		squareColumn = 1
	}
	if column >= 6 && column < 9 {

		squareColumn = 2
	}

	//Creation of two loops that will count the 9 blocks in the lines and columns
	for j := 0; j < 3; j++ {

		for k := 0; k < 3; k++ {

			square[j][k] = sudoku[(squareRow*3)+j][(squareColumn*3)+k]
			/**allocation of the lines and columns by the counters j and k in the square array.
			Then, we put the square array into the sudoku. We multiply the SquareRow and SquareColumn by 3, to have an double array of 9 box by blocks in the sudoku array.
			**/
		}

	}

	return square

}

func completeSudoku(sudoku [9][9]int) [9][9]int {
	/**Creation of the function that solve the sudoku grid with the backtrakking algorithm.
	This function have in parameter the suddoku of the function creationArray, and return an double array of int type. **/

	var tab [81][3]int
	/**creation of an double array which will store the row, column, and the numbers added in the array.
	It's a backup array with all the possible combination added in the array and solved the sudoku**/
	var x int = 0
	verif := true

	for row := 0; row < 9; row++ {

		for column := 0; column < 9; column++ {

			if sudoku[row][column] == 0 || verif == false {

				for num := 1; num <= 10; num++ {

					if num == 10 {

						for z := 0; z < 3; z++ {
							tab[x][z] = 0
						}

						sudoku[row][column] = 0

						x-- // to decrease by 1 the place of the argument // backtrakking

						row = tab[x][0]
						column = tab[x][1]
						num = tab[x][2]
						verif = false
					}

					if num < 10 {

						verif = verifParamSudoku(sudoku, num, row, column)
					}

					if verif == true {

						sudoku[row][column] = num

						tab[x][0] = row
						tab[x][1] = column // sauvegarde des differentes valeurs
						tab[x][2] = num
						x++ //to increase by 1 the place of the arguments // backtrakking

						break // stop the loop "for num"

					}

				}
			}
		}

	}
	return sudoku

}

func verifParamSudoku(sudoku [9][9]int, num int, row int, column int) bool {
	/** This function will check if the program can add a number into a cell. The function have in paramter the sudoku array,
	the row, column and num. Then the fonction display a boolean type (true or false)**/

	for i := 0; i < 9; i++ { // check in the lines and columns of the sudoku array, if there are a number.
		//the loop allow to count into the lines and columns
		if sudoku[i][column] == num {

			return false
		}
		if sudoku[row][i] == num {

			return false
		}
	}

	square := creationSquare(sudoku, row, column)

	/**call the square array of the function creation square to check in the 9 blocks if in the lines and columns there are a number.**/

	for a := 0; a < 3; a++ {

		for b := 0; b < 3; b++ {

			if square[a][b] == num {

				return false
			}
		}
	}

	return true

}

func verifArgs() bool {
	/**The function have nothing in parameter and return a boolean type.
	This function will check if the arguments of the program, or the sudoku, is available.**/

	arguments := os.Args[1:] // take the first arguments of the program, or the first line of the sudoku until the last by ":"

	if len(arguments) != 9 { //check if the program have 9 arguments

		fmt.Println("Error")
		return false
	}

	var row []rune // Array of rune type which will store the range of arguments.

	for index := range arguments { // Range will divide all the arguments one by one.

		row = []rune(arguments[index]) // store in a slice the arguments and their indexes.

		if len(row) != 9 { // check if in the arguments there are 9 characters.
			fmt.Println("Error")
			return false
		}
	}
	for a := range row { // Check if in the characters of the arguments, we have only numbers.
		if !((row[a] >= '0' && row[a] <= '9') || row[a] == '.') { // "!()" allow us to tell to the program = if we don't have numbers print error.

			fmt.Println("Error")
			return false
		}
	}

	sudoku := creationArray() // call the function creationArray and the array sudoku

	for a := 0; a < 9; a++ { // loop to count the lines

		for b := 0; b < 9; b++ { // loop to count the columns

			if sudoku[a][b] > 0 {

				num := sudoku[a][b]
				sudoku[a][b] = 0

				if verifParamSudoku(sudoku, num, a, b) == false {

					fmt.Println("Error")
					return false
				}

				sudoku[a][b] = num
			}
		}
	}

	return true
}
