package pointers

// Swap sans pointeur
// a := 10, b := 15
func Swap(adressea *int, adresseb *int) {

	tmp := *adressea // *(824 633 966 440) ==> 150
	*adressea = *adresseb
	*adresseb = tmp

}
