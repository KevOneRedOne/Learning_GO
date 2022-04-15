package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListSize(l *List) int {

	var length int

	if l.Head == nil {

		return length

	} else {

		length = 1
		currentNode := l.Head

		for currentNode != nil {

			currentNode = currentNode.Next
			length++
		}
		return length
	}

}
