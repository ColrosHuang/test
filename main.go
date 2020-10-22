package main

import (
	"godemo/search"
	"log"
	"os"

	_ "godemo/matchers"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	var searchTerm = "president"
	if len(os.Args) > 1 {
		searchTerm = os.Args[1]
	}
	// Perform the search for the specified term.
	search.Run(searchTerm)
	//search.Run("hello")
}
