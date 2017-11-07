/**
	The MatchingParenthesis.

	Sometimes (when I nest them (my parentheticals) too much (like this (and this))) they get confusing.

	Write a function that, given a sentence like the one above, along with the position of an opening
	parenthesis, finds the corresponding closing parenthesis.

	Example: if the example string above is input with the number 10 (position of the first parenthesis),
	the output should be 79 (position of the last parenthesis).
 */
package main

import (
	"fmt"
	"github.com/szeyick/GoExamples/Stack"
)

/**
	The application main.
 */
func main() {
	findParanthesisPosition()
}

/**
	Function to find the position of the matching parenthesis.
 */
func findParanthesisPosition() {
	stack := Stack.CreateStack(100)
	input := getInput()
	var startPosition int = 10
	var matchingPosition = startPosition
	stack.Push(string(input[startPosition]))

	for i := startPosition + 1; !stack.IsEmpty() && i < len(input); i++ {
		value := input[i]
		matchingPosition++
		if value == '(' {
			stack.Push(string(value))
		}
		if value == ')' {
			stack.Pop()
		}
	}
	fmt.Println("Matching parenthesis position: ", matchingPosition)
}

/**
	Return an input string.
 */
func getInput() (string) {
	return "Sometimes (when I nest them (my parentheticals) too much (like this (and this))) " +
		"they get confusing."
}