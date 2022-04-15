package main

import "fmt"

func main() {

	fmt.Println("----------- POINTERS -----------")
	// pointeur 	: Une variable qui stocke une adresse mémoire (Numéro d'une case en mémoire)

	// *int 		: Type pointeur vers un entier
	// *float32 	: Type pointeur vers un float32

	// &nombre 	: Permet de récupérer l'adresse de la variable nombre
	// *pNombre : Permet d'accéder à la valeur qui est stockée à l'adresse pNombre
	// Déclaration func Swap(pa *int, pb *int) { ... } ==> Appel: Swap(&a, &b)

	// Syntaxe :
	var nombre int64 = 150       // Déclaration d'une variable de type int + Affectation de la valeur 150
	var pNombre *int64 = &nombre // Déclaration + Affectation d'un pointeur (Obention de l'adresse de la case mémoire où est stockée la variable nombre)

	fmt.Println(nombre)   // Accès au contenu de la variable nombre en utilisant l'identifieur
	fmt.Println(*pNombre) // Accès au contenu de la variable nombre en utilisant l'adresse
	fmt.Printf("La variable nombre possède la valeur %v et est stockée à la case d'adresse (numéro) %v \n", nombre, pNombre)
	fmt.Printf("La variable nombre possède la valeur %v et est stockée à la case d'adresse (numéro) %v \n", *pNombre, &nombre)
	println()

	// Modification de la valeur associée à la variable nombre en utilisant son identifieur
	nombre = 37
	fmt.Println(nombre)   // Accès au contenu de la variable nombre en utilisant l'identifieur
	fmt.Println(*pNombre) // Accès au contenu de la variable nombre en utilisant l'adresse
	fmt.Printf("La variable nombre possède la valeur %v et est stockée à la case n° %v \n", nombre, pNombre)
	fmt.Printf("La variable nombre possède la valeur %v et est stockée à la case n° %v \n", *pNombre, &nombre)
	println()

	// Modification de la valeur associée à la variable nombre en utilisant son adresse
	*pNombre = 87
	fmt.Println(nombre)   // Accès au contenu de la variable nombre en utilisant l'identifieur
	fmt.Println(*pNombre) // Accès au contenu de la variable nombre en utilisant l'adresse
	fmt.Printf("La variable nombre possède la valeur %v et est stockée à la case n° %v \n", nombre, pNombre)
	fmt.Printf("La variable nombre possède la valeur %v et est stockée à la case n° %v \n", *pNombre, &nombre)
	println()

	fmt.Println("----------- ARRAY -----------")

	// ARRAY : La taille de l'array est statique : On ne peut ajouter de nouveaux élements
	// Syntaxe: [TAILLE]type{ , , }
	var arr = [5]int{25, 28, 37, 42, 526} // Les cases pour lequelles aucune valeur n'est renseignée sont initialisées
	// à la valeur par défaut du type donné après les crochets []int, dans ce cas 0.
	fmt.Println(len(arr))
	fmt.Println(arr)

	arr[2] = 205 // Modification de la valeur à l'indice 2: 37 devient 205
	fmt.Println(arr)

	arr[4] = 204 // Modification de la valeur à l'indice 4: 526 devient 205
	fmt.Println(arr)

	// Quelles sont les adresses des éléments de l'array ?
	for i := range arr {
		fmt.Println(&arr[i])
	}

	fmt.Println(arr[1:4])

	// Et si on voulait ajouter une nouvelle valeur dans notre array arr ?

	var newArr [6]int // On créé un nouveau tableau avec la taille souhaitée

	// On copie les éléments
	for index := range arr {
		newArr[index] = arr[index]
	}

	// On ajoute la nouvelle valeur
	newArr[5] = 10

	// Tous cela est assez fastidieux ==> On préfère utiliser des slices pour des tableaux à taille variable

	// Que se passe-t-il si on essaie d'ajouter une valeur en dehors de la taille du tableau ?
	// arr[7] = 152 -- invalid array index 7 (out of bounds for 7-element array)
	// ==> ERROR !!

	fmt.Println("----------- SLICE -----------")

	// SLICE : La taille du slice est dynamique
	var slice []int // Le slice est initialisé à nil

	if slice == nil {
		println("Slice is nil !")
	}

	slice = []int{59, 82, 436}
	fmt.Println(slice) // Affichage du slice à la création

	fmt.Println(len(slice)) // Affichage de la longueur à la création
	fmt.Println(cap(slice)) // Affichage de la capacité à la création

	// Afficher les valeurs du slice : Méthode 1
	for i := 0; i < len(slice); i++ {
		fmt.Printf("A l'indice %v, il y a la valeur %v. Son adresse est %v \n", i, slice[i], &slice[i])
	}
	println()

	// La capacité est insuffisante
	slice = append(slice, 85) // Ajout d'un élément dans le slice: append (reslice)
	fmt.Println(len(slice))   // Affichage de la longueur après ajout d'un élément
	fmt.Println(cap(slice))   // Affichage de la capacité après ajout d'un élément

	// Afficher les valeurs du slice : Méthode 2
	for index := range slice {
		fmt.Printf("A l'indice %v, il y la valeur %v. Son adresse est %v \n", index, slice[index], &slice[index])
	}
	println()

	// La capacité maximale n'est pas encore atteinte
	slice = append(slice, 15) // Ajout d'un élément dans le slice: append (reslice)
	fmt.Println(len(slice))   // Affichage de la longueur après ajout d'un élément
	fmt.Println(cap(slice))   // Affichage de la capacité après ajout d'un élément

	// Afficher les valeurs du slice : Méthode 3
	for index, value := range slice {
		fmt.Printf("A l'indice %v, il y la valeur %v. Son adresse est %v \n", index, value, &value)
	}
	println()

	slice = append(slice, 145, 310, 222) // Ajout de plusieurs éléments dans le slice: append (reslice)

	// Peut-on accéder à des valeurs d'un slice entre sa length et sa capacité ?
	// for i := 0; i < cap(slice); i++ {
	// 	fmt.Printf("A l'indice %v, il y la valeur %v. Son adresse est %v \n", i, slice[i], &slice[i])
	// }
	// NON ===> panic: runtime error: index out of range [4] with length 4

	// Utilisation du slice
	fmt.Println(slice)
	fmt.Println(slice[2:5]) // On récupère les éléments entre les indices 2 et 4
	fmt.Println(slice[:5])  // On récupère les éléments à partir de l'indice 0 (début) jusqu'à l'indice 4
	fmt.Println(slice[3:])  // On récupère les éléments à partir de l'indice 3 jusqu'à la fin
	fmt.Println(slice[:])   // On récupère tous les éléments

	fmt.Println("----------- STRING -----------")
	var prenom string = "émilie" // ^e => ê
	fmt.Println(prenom)
	fmt.Println(prenom[0])
	println()

	// Les objets de type string sont immuables (immutable): On ne peut pas modifier ses caratères
	// prenom[0] = 'E' // ERROR ==> cannot assign to prenom[0]

	for _, char := range prenom {
		fmt.Println(char)
	}
	println()

	prenomRunes := []rune(prenom)
	for _, r := range prenomRunes {
		fmt.Println(r)
	}
	println()

	prenomRunes[0] = 'E'
	prenom = string(prenomRunes)
	fmt.Println(prenom)
	println()

	fmt.Println("----------- STRUCT -----------")

	// A struct is a collection of fields (Un ensemble de champs/propriétés)

	// Example 1

	type Date struct {
		Jour  int
		Mois  string
		Annee int
	}

	type EtatCivil struct {
		Prenom          string
		Nom             string
		DateDeNaissance Date
		Enfants         []string
	}

	var dateOfBirth Date = Date{1, "Octobre", 2001} // dateOfBirth := Date{1, "Octobre", 2001}
	fmt.Println(dateOfBirth)

	dateOfBirth.Mois = "Avril"
	fmt.Println(dateOfBirth)

	monEtatCivil := EtatCivil{"John", "Smith", dateOfBirth, []string{"Margueritte", "Alice", "Bernard"}}
	fmt.Println(monEtatCivil)

	var dates []Date                                 // Array of dates
	dates = append(dates, Date{24, "Octobre", 2015}) // Assigning dates[0] to a value

	// Listes chainees
	fmt.Println("----------- LINKED LISTS -----------")

	type NodeL struct {
		Data int // interface{} pour type générique
		Next *NodeL
	}

	type List struct {
		Head *NodeL
		// Tail *NodeL // tail = *(List.Head).Next
	}

	var node1 NodeL
	var node2 NodeL
	var node3 NodeL
	var node4 NodeL

	// Initialisation du node 1
	node1.Data = 59
	node1.Next = &node2

	// Initialisation du node 2
	node2.Data = 82
	node2.Next = &node3

	// Initialisation du node 3
	node3.Data = 436
	node3.Next = &node4

	// Initialisation du node 4
	node4.Data = 85
	node4.Next = nil

	// Affiche les éléments de la liste
	pNode := &node1
	for pNode != nil {
		fmt.Println(*pNode)
		pNode = (*pNode).Next
	}

}
