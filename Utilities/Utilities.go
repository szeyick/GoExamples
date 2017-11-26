/**
	The Utilities.
 */
package Utilities

import (
	"math/rand"
	"fmt"
)

// range specification, note that min <= max
type IntRange struct {
	min, max int
}

/**
	Print the contents of the array.
 */
func PrintContents(input []int) {
	for i := 0; i < len(input); i++ {
		fmt.Print(input[i], " ")
	}
}

/**
	Return an array of random integers
	@param min - The min range
	@param max - The max range
	@param n - The array size
 */
func CreateRandomArrayOfIntegers(min int, max int, n int) ([]int) {
	var array []int = make([]int, n, n)

	r := rand.New(rand.NewSource(55))
	ir := IntRange{min,max}
	for i := 0; i < n; i++ {
		//fmt.Println(ir.randomValue(r))
		array[i] = ir.randomValue((r))
	}
	return array
}

/**
	get next random value within the interval including min and max
 */
func (ir *IntRange) randomValue(r* rand.Rand) int {
	return r.Intn(ir.max - ir.min +1) + ir.min
}

/**
	return the minimum value first, then the other value
 */
func Min(value1 int, value2 int) (int, int) {
	if value1 < value2 {
		return value1, value2
	} else {
		return value2, value1
	}
}
