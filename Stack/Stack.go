/**
	The Stack package.

	This package contains an implementation of a Stack and its associated functions.

	This stack is used to store characters.
 */
package Stack

/**
	The data structure representing the stack.
 */
type container struct {
	array[] string
	index int
	size int
}

/**
	Push a value onto the stack.
 */
func (stack *container) Push(value string) {
	stack.array[stack.index] = value
	stack.index++
	stack.size++
}

/**
	Pop a value off the stack
 */
func (stack *container) Pop () (string) {
	var value string = "Item not found"
	if stack.size > 0 {
		value = stack.array[stack.index];
		stack.index--
		stack.size--
	}
	return value
}

/**
	Return the size of the stack.
 */
func (stack *container) Size() (int) {
	return stack.size
}

/**
	Return a look at the top most item without popping it off.
 */
func (stack *container) Peek() (string) {
	if stack.size == 0 {
		return ""
	}
	return stack.array[stack.size - 1]
}

/**
	Return true if the stack contains no items.
 */
func (stack *container) IsEmpty() (bool) {
	return stack.size == 0
}

/**
	Create a stack of a given size.
 */
func CreateStack(size int) (container) {
	return container{array:make([]string, size, size), index:0, size:0}
}