package piscine

import "github.com/01-edu/z01"

func Recurneg(n int) {
	if n == 0 {
		return
	} else {
		c := '0'
		for i := 0; i > n%10; i-- {
			c++
		}
		Recurneg(n / 10)
		z01.PrintRune(c)
	}
}

func Recurpos(n int) {
	if n == 0 {
		return
	} else {
		c := '0'
		for i := 0; i < n%10; i++ {
			c++
		}
		Recurpos(n / 10)
		z01.PrintRune(c)
	}
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune(45)
		Recurneg(n * -1)

	} else if n == 0 {
		z01.PrintRune(48)

	} else {
		Recurpos(n)
	}
}
