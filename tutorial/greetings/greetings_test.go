package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	/* Function to test Hello method with input arg
	 * Will will test if response contains at least the name defined as input
	 * Expected success
	 */
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	/* Function to test Hello method without input arg
	 * Expected failure
	 */
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
