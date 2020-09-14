/*
 * @program: Go-In-Action
 * @author: Leon
 * @create: 2020-09-14 17:19
 */
package main

import (
	"./chapter2/sample/search"
	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"log"
	"os"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
