package main

/*
	L'informatique est la science du traitement de l'information.

	Instruction et variable.
*/

func main() {

	// Manipulation des variables
	// 1 - Déclaration de la variable
	// 2 - Assignation/Affectation d'une valeur à la variable

	// Pas d'espace dans l'identifieur/nom de la variable -> CamelCase
	// Ne pas utiliser d'accents
	//Le nom de votre variable ne peut pas débuter par un chiffre

	// 1- var nomDeLaVariable type = valeur

	// Déclaration           Affection/Assignation
	const pi = 3.14
	var etudiantPYC string = "John Smith"
	var age int = 19
	var classe rune = 'C'

	println(etudiantPYC)
	println(age)
	println(classe)

	// 2- var nomDeLaVariable type
	// Affectation implicite à une valeur par défaut

	var etudiantBYC string // ""
	var age2 int           // 0
	var classe2 rune       // ''

	print(etudiantBYC)
	print(age2)
	print(classe2)

	// 3- var nomDeLaVariable = valeur
	var ordinateur = "Dell"

	print(ordinateur)

	// 4- nomDeLaVariable := valeur
	// Inférence de type

	etudiant := "Lorine"
	entier := 15
	isRaining := false

	println(etudiant)
	println(isRaining)

	// Opérateurs d'assignation
	entier += 5 // entier = entier + 5
	println(entier)

	entier -= 5 // entier = entier - 5
	println(entier)

	entier *= 5 // entier = entier * 5
	println(entier)

	entier /= 5 // entier = entier / 5
	println(entier)

}
