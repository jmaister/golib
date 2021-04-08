package doublelinkedlist

// Double linked list node
type DLLNode struct {
	Prev  *DLLNode
	Next  *DLLNode
	Value interface{}
}

// Double linked list structure
type DoubleLinkedList struct {
	Head *DLLNode
	Tail *DLLNode
}

// Create a new and empty double linked list
func New() *DoubleLinkedList {
	return &DoubleLinkedList{
		Head: nil,
		Tail: nil,
	}
}

// Returns true if the list is empty
func (dll *DoubleLinkedList) IsEmpty() bool {
	return dll.Head == nil && dll.Tail == nil
}

// Insert a value after a node
func (dll *DoubleLinkedList) InsertAfter(node *DLLNode, value interface{}) *DLLNode {
	newNode := &DLLNode{
		Prev:  node,
		Value: value,
	}
	if node.Next == nil {
		newNode.Next = nil
		dll.Tail = newNode
	} else {
		newNode.Next = node.Next
		node.Next.Prev = newNode
	}
	node.Next = newNode

	return newNode
}

// Insert a value before a node
func (dll *DoubleLinkedList) InsertBefore(node *DLLNode, value interface{}) *DLLNode {
	newNode := &DLLNode{
		Next:  node,
		Value: value,
	}
	if node.Prev == nil {
		newNode.Prev = nil
		dll.Head = newNode
	} else {
		newNode.Prev = node.Next
		node.Prev.Next = newNode
	}
	node.Prev = newNode

	return newNode
}

// Insert a value as the new head
func (dll *DoubleLinkedList) InsertHead(value interface{}) *DLLNode {

	if dll.Head == nil {
		newNode := &DLLNode{
			Prev:  nil,
			Next:  nil,
			Value: value,
		}
		dll.Head = newNode
		dll.Tail = newNode
		return newNode
	} else {
		return dll.InsertBefore(dll.Head, value)
	}

}

// Insert a value as the new tail
func (dll *DoubleLinkedList) InsertTail(value interface{}) *DLLNode {
	if dll.Tail == nil {
		return dll.InsertHead(value)
	} else {
		return dll.InsertAfter(dll.Tail, value)
	}
}

// Remove a node from the list
func (dll *DoubleLinkedList) Remove(node *DLLNode) {
	if node.Prev == nil {
		dll.Head = node.Next
	} else {
		node.Prev.Next = node.Next
	}
	if node.Next == nil {
		dll.Tail = node.Prev
	} else {
		node.Next.Prev = node.Prev
	}

}
