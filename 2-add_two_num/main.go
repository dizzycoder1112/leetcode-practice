package main

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewNode(val interface{}) *Node {
	n := new(Node)
	n.Value = val
	return n
}

func NewLinkedList() *LinkedList {
	return &LinkedList{size: 0}
}

func (l *LinkedList) Push(data interface{}) (*Node, error) {
	node := NewNode(data)
	if l.size == 0 {
		l.head = node
		l.tail = node
	} else {
		l.tail.Next = node
		node.Prev = l.tail
		l.tail = node
	}
	l.size++
	return node, nil
}

func (l *LinkedList) Pop() (*Node, error) {
	var previous *Node

	if l.head == l.tail {
		previous = l.head
		l.head = nil
		l.tail = nil
		return previous, nil
	}

	for current := l.head; current != l.tail; current = current.Next {
		previous = current
	}

	tail := previous.Next
	previous.Next = nil
	l.tail = previous
	l.size--
	return tail, nil
}

func (l *LinkedList) Get(index int) interface{} {
	current := l.head
	count := 0
	for current != nil {
		if count == index {
			return current
		}
		current = current.Next
		count++
	}
	return nil
}

// func (l *LinkedList) RemoveAtIndex(index int) error {
// 	if index < 0 || index > l.Length()-1 {
// 		return fmt.Errorf("index out of range")
// 	}

// 	switch index {
// 	case 0:
// 		if l.head != nil {
// 			l.head = l.head.Next
// 		}
// 	default:
// 		current := l.head
// 		count := 0
// 		for current != nil && count < index-1 {
// 			current = current.Next
// 			count++
// 		}
// 		current.Next = current.Next.Next
// 		if count == l.Length() {
// 			current
// 		}
// 	}

// 	return nil
// }

// func (l *LinkedList) AddAtIndex(index int, data interface{}) error {
// 	if index < 0 || index > l.Length() {
// 		return fmt.Errorf("index out of range")
// 	}
// 	node := NewNode(data)

// 	switch index {
// 	case 0:
// 		node.Next = l.head
// 		l.head = node

// 	case l.Length():
// 		l.tail.Next = node
// 		l.tail = node

// 	default:
// 		current := l.head
// 		count := 0
// 		for current != nil && count < index-1 {
// 			count++
// 			current = current.Next
// 		}

// 		node.Next = current.Next
// 		current.Next = node

// 	}

// 	return nil
// }

// func (l *LinkedList) Pop() *Node {

// }

// type LinkedList interface {
// 	ToString() string
// 	Length() int
// 	isEmpty() bool
// 	removeAtIndex(index int)
// 	addAtIndex(index int)
// }

func main() {
}
