/**
	The KthToLastNode.

	This application class contains an implementation of an algorithm to find
	the kth to last node in an linked list.
 */
package main

import (
	"github.com/szeyick/GoExamples/LinkedList"
	"fmt"
)

/**
	The application main.
 */
func main() {
	head := LinkedList.CreateListWithValues(10)
	k := 5
	kthToLastNode := findKthToLastSlow(&head, k)
	printResult(kthToLastNode, k) // Print 6

	kthToLastNode = findKthToLastFast(&head, k)
	printResult(kthToLastNode, k) // Print 6

}

/**
	A function that will find the kth to last node in a list slowly.

	The slow way to find this is to loop through to see how many nodes are
	in the list, then subtract k from that which will resolve in the kth to
	last node when we iterate through the list again by the difference.
 */
func findKthToLastSlow (node *LinkedList.Node, k int) (*LinkedList.Node) {
	var nodeToFind = node.GetSize() - k
	currentNode := node
	for i := 0; i < nodeToFind; i++ {
		currentNode = currentNode.Next
	}
	return currentNode
}

/**
	A function that will find the kth to last node.

	This function will build the stick that is size k, then slide it along
	until the right end of the stick reaches the end. The left end of the stick
	in this instance is the kth last. This way is slightly faster as it does not require the
	initial call to find the size which takes O(n) time.
 */
func findKthToLastFast (node *LinkedList.Node, k int) (*LinkedList.Node) {
	rightStick := node
	for i := 0; i < k; i++ {
		rightStick = rightStick.Next
	}

	// Now that we have the right end of the stick, the
	leftStick := node
	for rightStick != nil {
		leftStick = leftStick.Next
		rightStick = rightStick.Next
	}
	return leftStick
}

/**
	Print the result.
 */
func printResult(node *LinkedList.Node, k int) {
	fmt.Println("The ", k, "th last node is - ", node.Value)
}
