package main

import (
	"fmt"
	"log"

	/* Import my local module locally instead of remote :
	 * 'go mod edit -replace example.com/greetings=../greeting'
	 * and sync it on hello folder :
	 * 'go mod tidy'
	 */
	"example.com/greetings"
)

func main() {
	// Init Log
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greetings msg with empty string, expected error
	message, err := greetings.Hello("Gladys")
	if err != nil {
		// Fatal msg and leaving program
		log.Fatal(err)
	}

	// If not error, print msg
	fmt.Println(message)
}
