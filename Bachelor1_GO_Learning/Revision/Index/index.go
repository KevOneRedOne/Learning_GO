package index

//Index is a function which returns the index of the first character of s present in tofind
func Index(s string, toFind string) int {

	// 1 - Count the lentgh of a string with the function StrLen
	lentghS := StrLen(s)
	lentghtofind := StrLen(toFind)

	// 2 - Browse the string of character with the function
	for index := 0; index < lentghS-lentghtofind; index++ {

		// 3 - Compare the length of the string and the index
		if s[index:(index+lentghtofind)] == toFind {

			return index

		}
	}
	return -1

}

//StrLen return the length of a string
func StrLen(s string) int {

	length := 0

	for range s {

		length++

	}
	return length

}
