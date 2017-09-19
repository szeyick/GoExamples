/**
	Write a function that takes an integer flightLength (in minutes) and
	an array of integers movieLengths (in minutes) and returns a boolean
	indicating whether there are two numbers in movieLengths whose sum equals flightLength.

	When building your function:

	- Assume your users will watch exactly two movies
	- Don't make your users watch the same movie twice
	- Optimize for runtime over memory
 */
package main

import (
	"fmt"
)

/**
	The program main.
 */
func main() {
	var movieLengths = []int{1, 2, 3, 4, 5}
	fmt.Println("2 movies to equal flight length: ", hasEqualMovies(4, movieLengths))
	fmt.Println("2 movies to equal flight length: ", hasEqualMovies(25, movieLengths))
}

/**
	For each index, get the current movie length and subtract it from the flight
	length to get the value to find in the set. Store the current movie
	into a map once the comparison is done to avoid watching the same
	movie twice.
 */
func hasEqualMovies(flightLength int, movieLengths []int) bool {
	var watchedMoviesMap map[int]bool
	watchedMoviesMap = make(map[int]bool)

	for i := 0; i < len(movieLengths); i++ {
		timeLength := flightLength - movieLengths[i]
		if watchedMoviesMap[timeLength] {
			return true
		}
		watchedMoviesMap[movieLengths[i]] = true
	}
	return false;
}
