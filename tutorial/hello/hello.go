package main

import (
	"fmt"
	// Requirement: 'go get rsc.io/quote' and generated go.sum file
	"rsc.io/quote"
)

func main() {
	fmt.Println("Quote package from module rsc.io:")
	fmt.Println("- quote.Go()	  : " + quote.Go())
	fmt.Println("- quote.Glass(): " + quote.Glass())
	fmt.Println("- quote.Hello(): " + quote.Hello())
	fmt.Println("- quote.Opt()  : " + quote.Opt())
}
