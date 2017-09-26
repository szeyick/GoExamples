/**
	An implementation of Insertion Sort in Go.

	For insertion sort, the input array is split into 2 sections, the sorted
	part and the unsorted part. There are two loops with this algorithm, the outer loop
	iterates through the entire input and grows the sorted section each time. The inner loop
	represents the unsorted portion and it always takes the first item in the unsorted portion
	and compares it with the items in the sorted part to see if they need to be swapped.
 */
package main

import (
	"fmt"
	"time"
	"math/rand"
)

/**
	The program main.
 */
func main() {
	fmt.Println("An implementation of Insertion Sort in Go!")
	input := InitRandInput(50)
	PrintContents(input)
	fmt.Println()
	insertionSort(input)
	PrintContents(input)
}

/**
	Populate an array with random integers to use as input.
	Numbers are generated between 0 and 100
 */
func InitRandInput(size int) []int {
	var input []int = make([]int, size, size)
	seed := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(seed)
	for i := 0; i < size; i++ {
		input[i] = randomGenerator.Intn(100)
	}
	return input
}


/**
	Implementation of insertion sort in Go.
 */
func insertionSort(input []int) {
	for i := 0; i < len(input); i++ {
		for j := i; j > 0; j-- {
			currentValue := input[j]
			adjacentValue := input[j - 1]
			if currentValue < adjacentValue {
				Swap(j, j - 1, input)
			}
		}
	}
}

/**
	Swap the values from a and b from the input array.
 */
func Swap(a int, b int, input []int) {
	tmp := input[a]
	input[a] = input[b]
	input[b] = tmp
}

/**
	Print the contents of the array.
 */
func PrintContents(input []int) {
	for i := 0; i < len(input); i++ {
		fmt.Print(input[i], " ")
	}
}

