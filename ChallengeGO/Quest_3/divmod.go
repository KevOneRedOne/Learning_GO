package piscine

/**
import (
	"fmt"
	piscine ".."
)

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	piscine.DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
**/

func DivMod(a int, b int, div *int, mod *int) {

	*div = a / b
	*mod = a % b /** % signifie modulo, soit le reste d'une division**/

}

/** Ecrire une fonction qui :
- divise le int a et int b ;
- Le résultat de cette division sera stocké dans l'int pointeur div
- Le modulo de la division sera stocké dana l'int pointeur mod **/
