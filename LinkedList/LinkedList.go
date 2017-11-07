/**
	The LinkedList Package.

	This package provides a basic implementation of a LinkedList.

	It should be noted that this implementation only allows for integers to be stored.
 */
package LinkedList

import "fmt"
/**
	A data structure that represents a single
	node in a linked list.

	The properties are defined as follows:

	value - The value that the node holds.
	next - A pointer to the next node in the list.
 */
type Node struct {
	Value int
	Next *Node
}

/**
	Print the contents of the list.
 */
func (node *Node) Print() {
	currentNode := node
	for currentNode != nil {
		fmt.Print(currentNode.Value, " ")
		currentNode = currentNode.Next
	}
	fmt.Println()
}

/**
	Insert a value into the node.
 */
func (node *Node) Insert(valueTmp int) {
	if node.Next == nil {
		newNode := Node{Value:valueTmp}
		node.Next = &newNode
	} else {
		node.Next.Insert(valueTmp)
	}
}

/**
	Return the size of the LinkedList.
 */
func (node *Node) GetSize() (int) {
	var size int = 0
	currentNode := node
	for currentNode != nil {
		size++
		currentNode = currentNode.Next
	}
	return size
}

/**
	Initialise a linked list with incremental values.

	The properties are defined as follows -
	size - The size of the linked list.
 */
func CreateListWithValues(size int) (Node) {
	var node = Node{Value: 0}
	for i := 1; i < size; i++ {
		node.Insert(i)
	}
	return node
}
