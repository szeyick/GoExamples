/**
Maintain a reference to the largest element in a stack.

Provide an additional method, getMax() that returns the largest item
in the stack.

The stack will only contain integers.
 */
package main

import (
	"fmt"
	"time"
	"math/rand"
)

/**
	An implementation of the largest stack.
 */
type LargestStack struct {
	stack[] int
	largestItemStack[] int
	size int
	stackIndex int
	largestItemIndex int
}

/**
	The application main.
 */
func main() {
	var stack = LargestStack{stack:make([]int, 10, 10), largestItemStack:make([]int, 10,10),
		size:0, stackIndex:0, largestItemIndex:0}

		pushItems(&stack)

		fmt.Println("")
		fmt.Println("The largest item in the stack is: ", stack.getMax())
}

/**
	Add 10 items onto the stack.
 */
func pushItems(stack *LargestStack) {
	seed := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(seed)
	for i :=0; i < 10; i++ {
		var value int = randomGenerator.Intn(100)
		stack.push(value)
		fmt.Println(value)
	}
}

/**
	Return the largest item off the stack.
 */
func (container * LargestStack) getMax() (int) {
	return container.largestItemStack[container.largestItemIndex]
}

/**
	Method that will be part of the stack struct.
 */
func (container *LargestStack) push (value int) {
	container.stack[container.stackIndex] = value
	container.stackIndex++;

	if container.largestItemIndex == 0 || container.largestItemStack[container.largestItemIndex] < value {
		container.largestItemIndex++
		container.largestItemStack[container.largestItemIndex] = value
	}
}