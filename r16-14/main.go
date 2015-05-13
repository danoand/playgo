package main

import "fmt"

var (
	sortMe = []int{61, 64, 95, 19, 63, 46, 92, 87, 17, 22, 43, 59, 92, 72, 83, 18, 24, 50, 60, 90}
)

// Function insertVal inserts the current unsorted element within the sorted subset
//   in ascending sort order
func insertVal(k int, inVal int) {
	// noInsert is a bool flag indicating if an insertion was made
	noInsert := true
	var shuffleVal int
	var tmpVal int

	// Iterate through the sorted subset of the slice
	for i := 0; i < k+1; i++ {
		// If no insert has been done as yet, then check inVal against the current unsorted element
		if noInsert {
			tmpVal = sortMe[i]

			if inVal <= sortMe[i] {
				// inVal is less than current value, insert value in the current location
				sortMe[i] = inVal
				fmt.Printf("Just inserted %v at index: %v; my array is now:\n%v\n\n", inVal, i, sortMe)
				// Set a value that needs to shuffled one element in the array
				shuffleVal = tmpVal
				// Set the flag that an insert has been made
				noInsert = false
			}

		} else {
			// An insert has been made, we need to shuffle values down the array
			tmpVal = sortMe[i]
			sortMe[i] = shuffleVal
			shuffleVal = tmpVal
		}
	}

	return
}

// Function exSort triggers an insertion sort on an array (Go slice)
func exSort() {
	// Set up a variable to hold the index of the "unsorted" subset
	//    Start at 1 as the first element (i=0) is sorted by definition
	uIndex := 1
	// Set up a variable to indicate if the processing should continue (or cease)
	actFlag := true

	// Insert the current unsorted element within the "sorted" subset
	// Iteratively process the sort until completion
	for actFlag {
		// If we're past the last element in the array, stop iterating
		if uIndex >= len(sortMe) {
			actFlag = false
			break
		}

		// No finished yet, insert/sort the existing element
		insertVal(uIndex, sortMe[uIndex])

		// Increment the index - move to the next element in the unsorted array
		uIndex++
	}

	return
}

func main() {
	fmt.Printf("Execute an insertion sort on this array:\n\n%v\n\n", sortMe)
	// Execute the sort
	exSort()
	fmt.Printf("The sorted array is now:\n\n%v\n\n", sortMe)

}
