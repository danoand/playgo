package main

import "fmt"

var (
	myArrayEven = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	myArrayOdd  = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
)

// Function evenCheck checks that the passed array index
// is the first index of the array's "middle" elements
// Example: ("b", "c") in ["a", "b", "c", "d"]
func evenMiddleCheck(k int, inLen int) (retBool bool) {
	retBool = false
	if k == inLen/2-1 {
		retBool = true
	}

	return
}

func oddMiddleCheck(k int, inLen int) (retBool bool) {
	retBool = false
	if k == inLen/2 {
		retBool = true
	}

	return
}

func revArray(inArr []string) (retArr []string, retErr error) {
	flagOdd := false
	arrLen := len(inArr)
	var scratch string

	// Check whether the array is odd
	if arrLen%2 == 1 {
		fmt.Println("Confirmed the array has an odd number of elements.")
		flagOdd = true
	}

	// Set up a reverse index variable
	j := arrLen - 1

	// Iterate through the array
	for i := 0; i < len(inArr); i++ {
		// Check we have reached the array "middle"
		if flagOdd {
			// Iterating over an array with an odd number of elements
			if oddMiddleCheck(i, arrLen) {
				// Completed processing an array with an odd number of elements
				retArr = inArr
				return
			}
		} else {
			// Iterating over an array with an even number of elements
			if evenMiddleCheck(i, arrLen) {
				// Reached the middle element(s) of an array with an even number of elements
				// "Swap" the middle elements and we're done
				scratch = inArr[i]
				inArr[i] = inArr[i+1]
				inArr[i+1] = scratch

				// Completed processing an array with an even number of elements
				retArr = inArr
				return
			}
		}

		// Have not reached the middle, just "swap" the elements to reverse the array
		scratch = inArr[i]
		inArr[i] = inArr[j]
		inArr[j] = scratch

		// Step to the next element from the end of the array (towards the front)
		j = j - 1

	}

	// Iterated through the entire array.  If we have not returned as yet, we have an error situation
	retErr = fmt.Errorf("ERROR: Processed the array but did not find the middle")
	return
}

func main() {
	fmt.Println("Start processing")
	fmt.Printf("The reverse of array:\n\n%v\n\n is:\n\n", myArrayEven)
	revArr, revErr := revArray(myArrayEven)
	fmt.Printf("is:\n\n%v\n\n with error: %v\n\n", revArr, revErr)

	fmt.Println("*************")

	fmt.Printf("The reverse of array:\n\n%v\n\n is:\n\n", myArrayOdd)
	revArr, revErr = revArray(myArrayOdd)
	fmt.Printf("is:\n\n%v\n\n with error: %v", revArr, revErr)

	fmt.Println("End processing")
}
