/**
The LinkedListLoop.

This code is responsible for demonstrating the algorithm for detecting
cycles in a linked list with Go.

 */
package main

import (
	"fmt"
	"github.com/szeyick/GoExamples/LinkedList"
)

/**
	The application main.
 */
func main() {
	var loopNode = createLoop();
	var hasLoop = hasLoop(&loopNode)
	fmt.Println("Has Loop: ", hasLoop)
}

/**
	Create a linked list with a loop in it.
 */
func createLoop() (LinkedList.Node) {
	var node1 = LinkedList.Node{Value: 1}
	var node2 = LinkedList.Node{Value: 2}
	var node3 = LinkedList.Node{Value: 3}

	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node1

	return node1
}

/**
	Evaluate if a LinkedList has a loop
	return true if the list has a loop, false otherwise.
 */
func hasLoop(head *LinkedList.Node) (bool) {
	hasLoop := false
	var slowPointer *LinkedList.Node = head
	var fastPointer *LinkedList.Node = head.Next

	for slowPointer != fastPointer && (fastPointer != nil) {
		fastPointer = fastPointer.Next
		if fastPointer == nil {
			break
		}
		if slowPointer == fastPointer {
			hasLoop = true
			break
		} else {
			slowPointer = slowPointer.Next
			fastPointer = fastPointer.Next
		}
	}
	return hasLoop
}
