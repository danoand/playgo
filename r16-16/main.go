package main

import (
	"fmt"
	"log"
)

var (
	myArray = []int{6, 8, 10, 12, -999, 14, 16, 3, 4, 5}
)

// Function findMe returns the minimum value in an array
func findMin(inArr []int) (retVal int) {
	var tmpVal int

	// Empty array? Return now.
	if len(inArr) == 0 {
		log.Printf("NOTE: the length of the passed array is: %v", len(inArr))
		return
	}

	// Iterate through a the array
	for i, v := range inArr {
		// Are we iterating on the first element?
		if i == 0 {
			tmpVal = v
		} else { // Not iterating on the first element
			// Is the current element less than our scratch value?
			if v < tmpVal {
				tmpVal = v
			}
		}
	}

	// Done iterating, tmpVal is the minimum value
	retVal = tmpVal

	return
}

func main() {
	fmt.Printf("Find the minimum value in:\n\n%v\n\n", myArray)
	minVal := findMin(myArray)
	fmt.Printf("The minimum value is:\n\n%v\n\n", minVal)
}
