/**
	Implement a queue with 2 stacks. Your queue should have an enqueue and a
	dequeue function and it should be "first in first out" (FIFO).
 */
package main

import (
	"fmt"
)

var pushStack Stack
var popStack Stack
var queueSize int

/**
	The struct representing the stack.
 */
type Stack struct {
	size int
	container[] int
}

func main() {
	fmt.Println("An implementation of a queue with 2 stacks.")
	initQueue()
	push(1)
	push(2)
	push(3)
	push(4)
	//fmt.Println(getQueueSize())
	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(pop())
	push(5)
	//fmt.Println(getQueueSize())
	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(pop())

}

/**
	Initialise the stacks.
 */
func initQueue() {
	pushStack = Stack{size: 0, container:make([]int, 5, 5)}
	popStack = Stack{size: 0, container:make([]int, 5, 5)}
	queueSize = 0
}

/**
	Push a value onto the stack.
 */
func push(value int) {
	// Cap the stack limit to 5.
	if pushStack.size < 5 {
		insertIndex := pushStack.size
		pushStack.container[insertIndex] = value
		pushStack.size++
	}
}

/**
	Pop an item off the stack. We should only ever pop off the pop stack.
	We transfer items from the push stack over to the pop stack, only if it
	is empty, then pop.
 */
func pop() int {
	var poppedValue = -1
	if popStack.size == 0 {
		if pushStack.size > 0 {
			move()
		}
	}
	if popStack.size > 0 {
		indexToPop := popStack.size
		poppedValue = popStack.container[indexToPop]
		popStack.size = indexToPop - 1
	}
	return poppedValue
}

/**
	Transfer items from push stack to pop stack.
 */
func move() {
	var pushStackSize = pushStack.size
	for i := pushStackSize; i >= 0; i-- {
		insertIndex := popStack.size;
		popStack.container[insertIndex] = pushStack.container[i]
		popStack.size++
	}
	popStack.size = pushStack.size
	pushStack.size = 0
}

/**
	Return the size of the queue. It will be the sum of the contents from
	both stacks.
 */
func getQueueSize() int {
	return pushStack.size + popStack.size
}


