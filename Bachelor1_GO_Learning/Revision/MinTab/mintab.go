package mintab

//MinTab is a function which print the minimum of an int onder an array
func MinTab(tab []int) int {
	if len(tab) == 0 {
		println("Array is empty")
		return -1
	} else {
		min := tab[0]
		for _, elt := range tab {
			if elt < min {
				min = elt
			}
		}
		return min
	}
}
