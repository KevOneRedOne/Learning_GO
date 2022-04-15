package piscine

// NodeL is a struct
type NodeL struct {
	Data interface{} //
	Next *NodeL
}

// List is a struct
type List struct {
	Head *NodeL // Address of the First element of the list
	Tail *NodeL // Address of the Last element of the list
}

// ListPushBack is a function that prints an int in a string base passed in parameters.
func ListPushBack(l *List, data interface{}) {

	// With Tail

	nodeToInsert := NodeL{Data: data} //creation of a variable which take the structure NodeL

	if l.Head == nil { // If the first node is equal nil, affect to the head to the tail of the list the address of the node to insert

		l.Head = &nodeToInsert
		l.Tail = &nodeToInsert

	} else { // If the first node isn't equal nil

		l.Tail.Next = &nodeToInsert // Affect to the tail of the list NodeL the address of the next node to insert
		l.Tail = &nodeToInsert      // Update the tail of the list
		// l.Tail = l.Tail.Next // update the tail of the list

	}
	/**
	//Second case without Tail
	if l.Head == nil {

		node := NodeL{data, nil}
		l.Head = &node

	} else {

		node := l.Head // we're on the first node of the list

		for node.Next != nil {

			node = node.Next

		}
		node.Next = &NodeL{data, nil}
	}
	**/

}
