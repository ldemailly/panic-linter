package main

import "fmt"

func main() {
	if false {
		fmt.Println("Nested Hello, world") // Comment inside if.
		/* we want "panic without comment" but can't test that on line before... */
		// panic("this is bad")
	} else {
		panic("this is ok") // A comment on the same line makes panic() ok.
	}
}
