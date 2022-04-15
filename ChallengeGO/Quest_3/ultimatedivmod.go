package piscine

/**
import (
	"fmt"
	piscine ".."
)

func main() {
	a := 13
	b := 2
	piscine.UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
**/

func UltimateDivMod(a *int, b *int) {
	var pointer = *a / *b // Création de 2 variables qui vont stocker les opérations
	var pointer2 = *a % *b
	*a = pointer
	*b = pointer2

}

/**
- Ecrire une fonction qui divise le int a et le int b
- Le résultat de la division est stocké dans le pointer de a
- le modulo de la division est stocké dans le pointer de b **/
