package main

import (
	"flag"
	"fmt"

	"assignment/assignment1"
	"assignment/assignment2"
	"assignment/assignment3"
)

func main() {
	// Define a flag for the user input
	assignmentFlag := flag.String("assignment", "", "The assignment number to run (1 or 2)")

	// Parse the flags
	flag.Parse()

	// Check which assignment the user wants to run
	switch *assignmentFlag {
	case "1":
		assignment1.Assignment1()
	case "2":
		assignment2.Assignment2()
	case "3":
		assignment3.Assignment3()
	default:
		fmt.Println("Invalid assignment number")
	}
}
