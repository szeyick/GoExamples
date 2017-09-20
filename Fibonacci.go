/**
	Write a function fib() that a takes an integer n and returns the nth fibonacci number.

	Let's say our fibonacci series is 0-indexed and starts with 0. So:

  	fib(0); // => 0
	fib(1); // => 1
	fib(2); // => 1
	fib(3); // => 2
	fib(4); // => 3

 */
package main

import (
	"fmt"
)

/**
	The program main.
 */
func main() {
	printFib()
}

/**
	Print the nth fibonacci.
 */
func printFib() {
	for i := 0; i < 10; i++ {
		fmt.Println("The", i, "th fibonacci number is: ", fib(i))
	}
}

/**
	Recursively find the nth fib number.
 */
func fib(nthFib int) int {
	if nthFib == 0 {
		return 0
	}
	if nthFib == 1 {
		return 1
	}
	return fib(nthFib - 1) + fib(nthFib - 2)
}

