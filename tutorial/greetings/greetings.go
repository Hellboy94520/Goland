package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Checking input arg
	if name == "" {
		// Do not use capital letter or punctuation (ST1005)
		return "", errors.New("empty name")
	}

	// Return a greeting random message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	/* I am do not understanding how that function is used:
	 * - Try to comment it but nothing noticeable has been updated/changed
	 */
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	/* Defined a string array with 3 values
	 * Generate an integer between 0 and the array lenght generated
	 * Returning address content from integer generated
	 */
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you back %v !",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
