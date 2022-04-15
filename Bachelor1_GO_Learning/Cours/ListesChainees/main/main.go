package main

import (
	"fmt"

	linkedList ".."
)

func SlicePushFront(slice []int, n int) []int {

	if slice == nil {
		return []int{n}
	} else {
		result := []int{n}
		return append(result, slice...)
	}
}

func SlicePushPosition(slice []int, n int, index int) []int {
	if index < 0 || index > len(slice) {
		fmt.Println("L'index n'est pas correcte")
		return nil
	} else {
		// slice = [1, 2, 3, 4, 5, 6]; n=15 index= 1
		// slice[:2] ==> [1,2]
		// slice[2:] ==> [3,4,5,6]
		// result <== [1,2,15]
		result := append(slice[:index], n)
		result = append(result, slice[index:]...)
		return result
	}
}

func main() {

	// Ressource pour comprendre les slices : https://blog.golang.org/slices-intro
	// Utilisation de make pour créer un slice
	var slice = make([]int, 8) // 8 représente la longueur, par défaut la capacité est initialisée à la longueur
	fmt.Println("slice: ", slice)

	var slice0 = make([]int, 8, 12) // 8 représente la longueur et 12 la capacité
	fmt.Println("slice0: ", slice0)

	slice0 = slice0[:10] // On ajoute deux cases au slice (les cases d'indices 8 et 9)
	fmt.Println()

	//
	var slice1 []int
	fmt.Println(slice1)

	// Ajouter des valeurs dans le slice ==> Append
	slice1 = append(slice1, 1, 2, 3, 4, 5, 6)
	fmt.Println(slice1)

	slice1 = SlicePushPosition(slice1, 15, 2)
	fmt.Println(slice1)

	// Utilisation d'un constructeur
	// On renseigne les éléments un à un
	var slice2 = []string{"Etudiant", "Paris", "Ynov"} // []string{}
	slice2 = append(slice2, "Campus")
	fmt.Println("slice2: ", slice2)

	var slice4 = []string{"a", "b", "c", "d"}
	fmt.Println("slice4: ", slice4)

	slice2 = append(slice2, slice4...)
	fmt.Println("slice2: ", slice2)

	slice2 = append(slice2, slice4[:3]...)
	fmt.Println("slice2: ", slice2)

	// On construit notre slice à partir d'un autre objet
	texte := "Paris Ynov Campus"
	var slice3 = []rune(texte) // []rune()
	fmt.Println(slice3)

	var slice5 []interface{}

	slice5 = append(slice5, 19, "Rayane", 17.5, []int{3, 4})
	fmt.Println("slice5: ", slice5)

	//slice5 = append(slice5, slice4...) ==> Erreur, problème dans les types
	for _, value := range slice4 {
		slice5 = append(slice5, value)
	}
	fmt.Println("slice5: ", slice5)

	// Tableau (slice/array) à deux dimensions
	// Syntaxte [][]type

	// [3]int = [. . .] - Les points représentent des entiers
	// [3]float = [. . .] - Les points réprésentent des flotants
	// [3]([5]int) = [ [. . . . .] [. . . . .] [. . . . .] ]
	/*
		[3]([5]int) = [
			[. . . . .]
			[. . . . .]
			[. . . . .]
		]
	*/
	var board [3][5]int

	board[0][2] = 1
	fmt.Println(board)

	/*
		[][]int = [
			[1 2 3]
			[4 5]
			[6 7 8 9]
			]
	*/

	var grid [][]int = [][]int{[]int{1, 2, 3}, []int{4, 5}, []int{6, 7, 8, 9}}
	fmt.Println(grid)
	grid[1] = append(grid[1], 10, 12, 14)
	fmt.Println(grid)

	slice6 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("slice6: ", slice6)

	// Création des noeuds de liste
	node6 := linkedList.NodeL{6, nil}
	node5 := linkedList.NodeL{5, &node6}
	node4 := linkedList.NodeL{4, &node5}
	node3 := linkedList.NodeL{3, &node4}
	node2 := linkedList.NodeL{Data: 2, Next: &node3}
	node1 := linkedList.NodeL{Data: 1, Next: &node2}

	// Création de la liste
	list := linkedList.List{&node1, &node6}
	linkedList.PrintList(list)

	link := linkedList.List{} // var link linkedList.List
	fmt.Println(link)

	for i := 0; i < 7; i++ {
		linkedList.ListPushBack(&link, i)
	}
	linkedList.PrintList(link)
	fmt.Println(linkedList.ListSizeNT(&link))
	fmt.Println(linkedList.ListSize(&link))

}
