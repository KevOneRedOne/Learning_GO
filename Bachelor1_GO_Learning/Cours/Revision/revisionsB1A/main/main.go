package main

import (
	"fmt"
	"os"

	revision "../revision"
)

func main() {

	// tab := []int{5, 8, 7, 6, 5}
	// tab := []int{19, 8, 7, 6, 5}
	// tab := []int{5, 8, 7, 6, 20}

	// max := rev.MaxTab(tab)
	// fmt.Println(max)

	/* Utilisation de os.Args
	os.Args renvoie un slice []string avec en index 0 le nom du programme
	et les paramètres du programme à partir de l'indice 1
	 [NomDeLaFonction, "Vitaly", "Nom", "20"]
	*/

	programName := os.Args[0]
	fmt.Println(programName)

	programParams := os.Args[1:]
	fmt.Println(programParams)

	texte := "Je suis etudiant a Ynov Paris Campus."
	mot := "Nantes"

	fmt.Println(revision.Index(texte, mot))

}
