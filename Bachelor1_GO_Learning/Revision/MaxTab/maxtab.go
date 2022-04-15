package maxtab

//MaxTab is a function which print the maximum of an int onder an array
func MaxTab(tab []int) int {

	if len(tab) == 0 {

		println("The Array is empty")
		return -1

	} else {

		max := tab[0]

		for _, element := range tab {

			if element > max {

				max = element
			}

		}
		return max

	}

}
