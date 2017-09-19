/**
	The BinarySearch.

	This file contains a recursive implementation of a binary search written in Go. It
	is written just as a means to get used to using slices (dynamic arrays), recursion
	and function calls for Go.

	An interview question revolving around a binary search would be as follows:

	Suppose we had an array of n integers sorted in ascending order. How quickly could
	we check if a given integer is in the array?
 */
package main

import (
	"fmt"
)

/**
	The program main.
 */
func main() {
	fmt.Println("An implementation of a binary search.")
	inputArray := initArray(10)
	printArray(inputArray)


	fmt.Println("The value 4 is found in the array: ", binarySearch(inputArray, 4, 0, len(inputArray)))
	fmt.Println("The value 10 is found in the array: ", binarySearch(inputArray, 9, 0, len(inputArray)))
	fmt.Println("The value 45 is found in the array: ", binarySearch(inputArray, 45, 0, len(inputArray)))

	doSomething(inputArray)
	printArray(inputArray)
}

/**
	Directly change an index of the array.
 */
func doSomething(inputArray [] int) {
	inputArray[2] = 1
}

/**
	Print the contents of the array to the console.
 */
func printArray(inputArray []int) {
	for i := 0; i < len(inputArray); i++ {
		fmt.Print(inputArray[i], " ")
	}
	fmt.Println()
}

/**
	Initialise the contents of the array.
 */
func initArray(size int) []int {
	var inputArray [] int
	inputArray = make([]int, size, size)
	for i := 0; i < size; i++ {
		inputArray[i] = i + 1;
	}
	return inputArray
}

/**
	A recursive implementation of a binary search in Go.
 */
func binarySearch(inputArray []int, valueToFind int, low int, high int) bool {
	if low < high {
		var middleIndex int = (high + low) / 2
		middleValue := inputArray[middleIndex]
		if valueToFind == middleValue {
			return true
		}
		if valueToFind < middleValue {
			return binarySearch(inputArray, valueToFind, low, middleValue)
		} else {
			return binarySearch(inputArray, valueToFind, middleValue + 1, high)
		}
	}
	return false
}