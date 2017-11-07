/**
	The LinkedListReversal.

	This application is responsible for demonstrating the reversal
	of a linked list.
 */
package main

import "github.com/szeyick/GoExamples/LinkedList"

/**
	The application main.
 */
func main() {
	node := LinkedList.CreateListWithValues(10)
	node.Print()
	reversedList := reverseList(&node)
	reversedList.Print()
}

/**
	Reverse the list.
 */
func reverseList(node *LinkedList.Node) (*LinkedList.Node) {
	currentNode := node
	var prevNode *LinkedList.Node
	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = prevNode

		prevNode = currentNode
		currentNode = nextNode
	}
	return prevNode
}
