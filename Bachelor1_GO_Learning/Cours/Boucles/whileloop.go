package loop

// Principe : Une boucle while permet de répéter un block d'instructions selon une condition donnée
// Principe : Une boucle for permet de répéter un block d'instructions un nombre de fois déterminé
// Syntaxe :
// i := 0;
// for  i < 10 {

// 	** Block d'instructions **
//  i+=5 				// Modification de la variable de boucle
// }
//

func CountWeeks() int {

	prix := 500
	semaine := 0

	economie := 0
	for economie < prix {
		economie += 20 //
		semaine++
	} // economie >= prix

	return semaine
}
