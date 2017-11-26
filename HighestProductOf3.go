/**

	The Highest Product of 3.

	Given an array of integers, find the highest product you can get from three of the integers.

	The input arrayOfInts will always have at least three integers.

	Solution: Use a greedy algorithm and hold the smallest and largest values found so far, along with
	finding the smallest product of 2 numbers, and the largest product of 2 numbers.

	At each iteration, we multiply the current value with the largest product and smallest products to see
	if we have a new highest product of 3. The reason that we need the smallest product of 2 is it is possible
	that the smallest possible product of 2 is -100 and the current value is 10 which flips the sign.
 */
package main

import (
	"github.com/szeyick/GoExamples/Utilities"
	"fmt"
)

var highestProduct int
var highestProductOf2 int
var lowestProductOf2 int
var highestValue int
var lowestValue int

/**
	The application main.
 */
func main() {
	arrayOfInts := Utilities.CreateRandomArrayOfIntegers(-10, 10, 10)
	highestProductOf3(arrayOfInts)
	fmt.Println()
	Utilities.PrintContents(arrayOfInts)
}

/**
	Find the highest product of 3.
 */
func highestProductOf3(arrayOfInts []int) {

	highestProductOf2 = arrayOfInts[0] * arrayOfInts[1]
	lowestProductOf2 = arrayOfInts[0] * arrayOfInts[1]
	lowestValue, highestValue = Utilities.Min(arrayOfInts[0], arrayOfInts[1])

	highestProductTmp := lowestProductOf2 * arrayOfInts[2]
	highestProductTmp2 := highestProductOf2 * arrayOfInts[2]

	if highestProductTmp > highestProductTmp2 {
		highestProduct = highestProductTmp
	} else {
		highestProductTmp = highestProductTmp2
	}

	for i := 2; i < len(arrayOfInts); i++ {
		currentValue := arrayOfInts[i]
		highestProductTmp := lowestProductOf2 * currentValue
		highestProductTmp2 := highestProductOf2 * currentValue

		highestProductTmp3 := highestProductTmp2
		if highestProductTmp > highestProductTmp2 {
			highestProduct = highestProductTmp
		}
		if highestProductTmp3 > highestProduct {
			highestProduct = highestProductTmp3
		}
		updateLowest2(currentValue)
		updateHighest2(currentValue)
	}
	fmt.Print(highestProduct)
}

/**
	Update the lowest values.
	Evaluate if the current value is the smallest encountered and also
	evaluate if we have a lowest product of 2.
 */
func updateLowest2(currentValue int) {
	lowestProductOf2Tmp := lowestValue * currentValue
	if lowestProductOf2Tmp < lowestProductOf2 {
		lowestProductOf2 = lowestProductOf2Tmp
	}
	if currentValue < lowestValue {
		lowestValue = currentValue
	}
}

/**
	Update the highest value
	Evaluate if the current value is the largest encountered and also
	evaluate if we have the highest product of 2.
 */
func updateHighest2(currentValue int) {
	highestProductOf2Tmp := highestValue * currentValue
	if highestProductOf2Tmp < highestProductOf2 {
		highestProductOf2 = highestProductOf2Tmp
	}
	if currentValue > highestValue {
		highestValue = currentValue
	}
}


