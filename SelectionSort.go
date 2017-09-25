/**
	This is an implementation of selection sort written in Go.
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
	fmt.Println("Selection Sort Written in Go.")
	input := initRandInput(50)
	printContents(input)
	fmt.Println()
	selectionSort(input)
	printContents(input)
}

/**
	Populate an array with random integers to use as input.
	Numbers are generated between 0 and 100
 */
func initRandInput(size int) []int {
	var input []int = make([]int, size, size)
	seed := rand.NewSource(time.Now().UnixNano())
	randomGenerator := rand.New(seed)
	for i := 0; i < size; i++ {
		input[i] = randomGenerator.Intn(100)
	}
	return input
}

/**
	Sort the input array using selection sort.
 */
func selectionSort(input []int) {
	for i := 0; i < len(input); i++ {
		largestIndex := i
		smallestValue := input[i]
		for j := i; j < len(input); j++ {
			currentValue := input[j]
			if currentValue < smallestValue {
				largestIndex = j
				smallestValue = currentValue
			}
		}
		swap(i, largestIndex, input)
	}
}

/**
	Swap the values from a and b from the input array.
 */
func swap(a int, b int, input []int) {
	tmp := input[a]
	input[a] = input[b]
	input[b] = tmp
}

/**
	Print the contents of the array.
 */
func printContents(input []int) {
	for i := 0; i < len(input); i++ {
		fmt.Print(input[i], " ")
	}
}