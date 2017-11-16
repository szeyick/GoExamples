/**
	The Top Scores.

	Each round, players receive a score between 0 and 100, which you use to rank them from highest to lowest. So far you're using an algorithm that sorts in O(n\lg{n})O(nlgn) time, but players are complaining that their rankings aren't updated fast enough. You need a faster sorting algorithm.

	Write a function that takes:

	1. an array of unsortedScores
	2. the highestPossibleScore in the game
	3. and returns a sorted array of scores in less than O(n\lg{n})O(nlgn) time.

	For Example:

	int[] unsortedScores = {37, 89, 41, 65, 91, 53};
	final int HIGHEST_POSSIBLE_SCORE = 100;

	int[] sortedScores = sortScores(unsortedScores, HIGHEST_POSSIBLE_SCORE);
	// sortedScores: [91, 89, 65, 53, 41, 37]
 */
package main

import "fmt"

const HIGHEST_POSSIBLE_STORE int = 100

/**
	The application main.
 */
func main() {
	var unsortedScores[] int = []int {37, 89, 41, 65, 91, 53}
	var sortedScores = sortScores(unsortedScores)
	printScores(sortedScores)
}

/**
	Use a bucket sort to sort the scores in O(n) time.
	@unsortedScores - The array of unsorted scores.
	@return the list of sorted scores.
 */
func sortScores(unsortedScores[] int) ([]int) {
	var sortedScores[] int = make([]int, len(unsortedScores))
	var possibleScores[HIGHEST_POSSIBLE_STORE] int

	// Store the scores into possible scores.
	for i := 0; i < len(unsortedScores); i++ {
		var scoreIndex = unsortedScores[i] - 1
		possibleScores[scoreIndex]++
	}

	var sortedScoresIndex = 0
	for i := len(possibleScores) - 1; i >= 0; i-- {
		var scoreCount = possibleScores[i] // Count of this score.
		for k := 0; k < scoreCount; k ++ {
			sortedScores[sortedScoresIndex] = i + 1
			sortedScoresIndex++ // Increment the storage index.
		}
	}
	return sortedScores
}

/**
	Print the scores.
 */
func printScores(scores[] int) {
	fmt.Print("Scores are: ")
	for i := 0; i < len(scores); i++ {
		fmt.Print(scores[i], " ")
	}
	fmt.Println()
}
