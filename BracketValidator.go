/**
	The BracketValidator.

	This application is demonstrating the bracket validator question that
	often comes up during coding interviews.
 */
package main

import (
	"github.com/szeyick/GoExamples/Stack"
	"fmt"
)
/**
	The application main.
 */
func main() {
	parenthesis1 := "{[]()}"
	parenthesis2 := "{[(])}"
	parenthesis3 := "{[}"
	fmt.Println("Parenthesis: ", parenthesis1, "is valid - ", isValid(parenthesis1, len(parenthesis1)))
	fmt.Println("Parenthesis: ", parenthesis2, "is valid - ", isValid(parenthesis2, len(parenthesis2)))
	fmt.Println("Parenthesis: ", parenthesis3, "is valid - ", isValid(parenthesis3, len(parenthesis3)))
}

/**
	Given a set of parenthesis, evaluate if they are correctly balanced.
 */
func isValid(parenthesis string, length int) (bool) {
	stack := Stack.CreateStack(length)
	for i := 0; i < length; i++ {
		var c = parenthesis[i]
		if c == '(' || c == '{' || c == '[' {
			stack.Push(string(c))
		} else {
			// It has to be a closing brace.
			if !stack.IsEmpty() {
				if  (c == ')' && stack.Peek() == "(") ||
					(c == ']' && stack.Peek() == "[") ||
					c == '}' && stack.Peek() == "{"	{
					stack.Pop()
				} else {
					return false;
				}
			} else {
				return false;
			}
		}
	}
	return stack.IsEmpty()
}
