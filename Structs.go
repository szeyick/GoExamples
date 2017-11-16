package main

import (
	"fmt"
)

/**
 * A struct is a data structure that can be used
 * to store custom data.
 */
type Book struct {
	title string
	author string
	id int
}

/**
	The application main.
 */
func main() {
	fmt.Println("This is a demo of Structs in Go!\n")
	book := Book {"1984", "George Orwell", 1}
	printBookInfo(book)

}

/**
	Print the stored book information.
 */
func printBookInfo(book Book) {
	fmt.Println("Book Information")
	fmt.Println("Title:", book.title)
	fmt.Println("Author:", book.author)
	fmt.Println("ID:", book.id)
}