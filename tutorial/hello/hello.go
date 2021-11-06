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

func fatalLog(err error) {
	if err != nil {
		// Fatal msg and leaving program
		log.Fatal(err)
	}
}

func main() {
	// Init Log
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greetings msg with empty string
	message, err := greetings.Hello("Gladys")
	fatalLog(err)
	fmt.Println(message)

	// Request several greetings msg
	names := []string{"Fladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	fatalLog(err)
	fmt.Println(messages)
}
