package pointers

// Swap sans pointeur
// a := 150, b := 15

func Swap(adressea *int, adresseb *int) {

	tmp := *adressea // *(824 633 966 440) ==> 150
	*adressea = *adresseb
	*adresseb = tmp

}

// StrRev is a program that returns the reversed string
func StrRev(s string) string {
	result := []rune(s)
	length := len(s)

	for i := 0; i <= length/2; i++ {
		result[i], result[length-1-i] = result[length-1-i], result[i]
	}

	return string(result)
}

// SliceRev is a program that returns a reversed version of the slice
func SliceRev(tab []int) {

	length := len(tab)

	for i := 0; i <= length/2; i++ {
		tab[i], tab[length-1-i] = tab[length-1-i], tab[i]
	}
	// fmt.Print("SliceRev : ")
	// fmt.Println(tab)

	// var result []int
	// for i := 0; i < length; i++ {
	// 	result = append(result, tab[length-1-i])
	// }
	// return result
}

// MaxTab([1,5,8,3,2]) ==> 8
func MaxTab(tab []int) int {

	return 0

}

func MinTab(tab []int) int {

	return 0

}

// SortTab sorts a slice using the Insertion Sort algorithm
func SortTab(tab []int) {

}
