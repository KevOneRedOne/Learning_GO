package linkedLists

import "fmt"

type NodeL struct {
	Data interface{} // Données contenue dans le noeud
	Next *NodeL      // Adresse du noeud suivant
}

type List struct {
	Head *NodeL // Adresse du premier noeud de la liste (Tête de liste)
	Tail *NodeL // Adresse du dernier noeud de la liste (Queue de liste)
}

func PrintList(l List) {
	// Cas de la liste vide
	if l.Head == nil {
		fmt.Println("List: []")
	} else {
		fmt.Print("List: [")
		// Cas liste non vide
		CurrentNode := l.Head
		for CurrentNode.Next != nil { // for node != l.Tail { // En utilisant Tail
			fmt.Print(CurrentNode, " ") // ou fmt.Print(node.Data, " ") si on veut afficher le noeud complet
			CurrentNode = CurrentNode.Next
		}
		// fmt.Print(node.Data) // fmt.Print(node.Data) si on veut afficher le noeud complet
		fmt.Print(CurrentNode)
		fmt.Println("]")
	}
}

func ListPushBackNT(l *List, data interface{}) {

	nodeToInsert := NodeL{} // NodeL{data, nil} // NodeL{Data: data}

	if l.Head == nil {
		l.Head = &nodeToInsert
	} else { // l.Head != nil
		currentNode := l.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		} // currentNode.Next == Nil : currentNode contient le dernier noeud de la liste
		currentNode.Next = &nodeToInsert
	}
}

func ListPushBack(l *List, data interface{}) {

	node := NodeL{data, nil} // Nouveau noeud, à ajouter en fin de liste
	if l.Head == nil {
		// La liste est vide (ne possède aucun noeud)
		l.Head = &node
		l.Tail = &node
	} else {
		// La liste possède au moins un noeud
		l.Tail.Next = &node
		l.Tail = &node // Mise à jour de la queue de liste
	}

}

func ListPushFront(l *List, data interface{}) {

	if l.Head == nil {
		// La liste est vide (ne possède aucun noeud)
		node := NodeL{data, nil}
		l.Head = &node
		l.Tail = &node
	} else {
		// La liste possède au moins un noeud
		node := NodeL{data, nil} // Nouveau noeud, ajouté en fin de liste
		l.Tail.Next = &node
		l.Tail = &node // Mise à jour de la queue de liste
	}

}

func ListSizeNT(l *List) int {

	var length int

	if l.Head == nil {
		return length
	} else {
		length = 1
		currentNode := l.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
			length++
		}
		return length
	}
}

func ListSize(l *List) int {

	var length int

	if l.Head == nil {
		return length
	} else {
		length = 1
		currentNode := l.Head
		for currentNode != l.Tail {
			currentNode = currentNode.Next
			length++
		}
		return length
	}
}
