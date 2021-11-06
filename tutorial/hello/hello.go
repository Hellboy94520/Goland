package main

import (
	"fmt"
	/* Import my local module locally instead of remote :
	 * 'go mod edit -replace example.com/greetings=../greeting'
	 * and sync it on hello folder :
	 * 'go mod tidy'
	 */
	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
