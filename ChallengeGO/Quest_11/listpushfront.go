package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListPushFront(l *List, data interface{}) {

	// Without Tail
	if l.Head == nil {

		l.Head = &NodeL{data, nil}
		l.Tail = &NodeL{data, nil}

	} else {

		node := &NodeL{data, l.Head}
		l.Head = node
	}
}
