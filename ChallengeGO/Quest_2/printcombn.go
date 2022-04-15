package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n < 1 || 9 < n {
		str := "Invalid input\n"
		for _, v := range str {
			z01.PrintRune(v)
		}
		return
	}
	var array [9]int
	arr := array[:n]
	last := n - 1
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	for {
		for i := 0; i <= last; i++ {
			z01.PrintRune(rune(arr[i] + '0'))
		}

		countMax := 0
		for i, v := range arr {
			if v == 10-n+i {
				countMax++
			}
		}
		if countMax == n {
			break
		}
		z01.PrintRune(',')
		z01.PrintRune(' ')
		arr[last-countMax]++
		for i := last - countMax; i < last; i++ {
			arr[i+1] = arr[i] + 1
		}
	}
	z01.PrintRune('\n')
}
