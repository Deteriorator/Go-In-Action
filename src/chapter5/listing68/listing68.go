/*
 * @program: Go-In-Action
 * @author: Leon
 * @create: 2020-09-15 22:58
 */
package main

import (
	"fmt"

	"chapter5/listing68/counters"
)

// main is the entry point for the application.
func main() {
	// Create a variable of the unexported type and initialize
	// the value to 10.
	counter := counters.New(10)

	// ./listing68.go:15: cannot refer to unexported name
	//                                         counters.alertCounter
	// ./listing68.go:15: undefined: counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
}
