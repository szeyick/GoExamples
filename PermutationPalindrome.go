/**
	Write an efficient function that checks whether any permutation of an input string is a palindrome.

	You can assume the input string only contains lowercase letters.

	Examples:

	"civic" should return true
	"ivicc" should return true
	"civil" should return false
	"livci" should return false

	Any permutation would mean that at the end, the remaining characters should be a count of either 0 in
	the case of an even length string or 1 in the case of an odd case string.
 */
package main

import "fmt"

/**
	The program main.
 */
func main() {
	string1 := "civic"
	string2 := "ivicc"
	string3 := "civil"
	string4 := "livci"

	fmt.Println(string1, " is a permutation palindrome: " , isPermutationPalindrome(string1))
	fmt.Println(string2, " is a permutation palindrome: " , isPermutationPalindrome(string2))
	fmt.Println(string3, " is a permutation palindrome: " , isPermutationPalindrome(string3))
	fmt.Println(string4, " is a permutation palindrome: " , isPermutationPalindrome(string4))
}

/**
	@word - The word to process.
	@return true if a permutation of the word can be a palindrome, false otherwise.
 */
func isPermutationPalindrome (word string) (bool) {
	var alphabet[26] int
	for i := 0; i < len(word); i++ {
		var c = word[i] - 'a'
		if alphabet[c] == 0 {
			alphabet[c]++
		} else {
			alphabet[c]--
		}
	}

	count := 0;
	for i := 0; i < len(alphabet); i++  {
		if alphabet[i] > 0 {
			count++
		}
	}
	return count <= 1
}
