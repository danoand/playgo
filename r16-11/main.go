package main

import "fmt"

var (
	queue1 = []int{1, 2, 3}
	queue2 = []int{1, 2}
)

// Function serve gets the element that is first in a queue
func serve(inQ []int) (retVal int, retQ []int) {
	// Are there elements in the queue?
	if len(inQ) == 0 {
		// The queue is empty return with a nil value
		return
	}

	// There are elements in the queue, pop off the first element
	retVal = inQ[0]
	inQ = inQ[1:len(inQ)]

	return
}

// Function getgetInLine puts a value in a queue
func getInLine(inVal int, inQ []int) {
	inQ = append(inQ, inVal)

	return
}

// Function testEqual tests if two queues are equal
//   Same number of elements, with the same values, in the same order
func testEqual(inQ1 []int, inQ2 []int) (retBool bool) {
	var v1, v2 int
	retBool = true

	// Test if queue lengths are not equal
	if len(inQ1) != len(inQ2) {
		// Queue lengths are not equal; return not equal indicator
		retBool = false
		return
	}

	// Test if queue length is zero
	if len(inQ1) == 0 {
		// Queue lengths are zero; queues are equal; return
		return
	}

	// Loop through an array
	for i := 0; i < len(inQ1); i++ {
		v1, inQ1 = serve(inQ1)
		v2, inQ2 = serve(inQ2)

		if v1 != v2 {
			retBool = false
			break
		}
	}

	return
}

func main() {
	// Test if queue1 and queue2 are equal
	tmpBool := testEqual(queue1, queue2)
	if tmpBool {
		fmt.Printf("\nQueue #1:\n%v\n\nis equal to\n\n%v\n", queue1, queue2)
	} else {
		fmt.Printf("\nQueue #1:\n%v\n\nis not equal to\n\n%v\n", queue1, queue2)
	}
}
