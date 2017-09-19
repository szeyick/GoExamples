package main

import "fmt"

func main() {
	var a int = 20   /* actual variable declaration */
	var ip *int      /* pointer variable declaration */

	ip = &a  /* store address of a in pointer variable*/

	fmt.Printf("Address of a variable: %x\n", &a  )

	// address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip )

	// access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip )

	// change the value of a through ip.
	*ip = 10
	fmt.Printf("Value of a by changing it through ip: %d\n", a)
	fmt.Printf("Value of ip: %d\n", *ip)
}