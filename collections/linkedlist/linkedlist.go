package linkedlist

type LLNode struct {
	Next  *LLNode
	Value interface{}
}

type LinkedList struct {
	Head *LLNode
	Tail *LLNode
}

// Create a new and empty linked list
func New() *LinkedList {
	return &LinkedList{
		Head: nil,
		Tail: nil,
	}
}

// Returns true if the list is empty
func (ll *LinkedList) IsEmpty() bool {
	return ll.Head == nil && ll.Tail == nil
}

// Append an element to the end of the list
func (ll *LinkedList) Append(value interface{}) *LinkedList {
	newNode := &LLNode{
		Next:  nil,
		Value: value,
	}

	if ll.IsEmpty() {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail.Next = newNode
	}
	ll.Tail = newNode
	return ll
}
