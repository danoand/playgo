package main

import "fmt"

var (
	array1 = []int{1, 3, 5, 7, 9}
	array2 = []int{0, 2, 4, 6, 8, 10}

	newArr []int
)

// Function mergeMe merges together two sorted arrays
func mergeMe(inAr1 []int, inAr2 []int) {
	// Flag to drive the processing
	procFlag := true

	// Variable to track inAr1's index value
	ind1 := 0
	ind2 := 0

	for procFlag {
		fmt.Printf("In loop with ind1: %v, ind2: %v, and procFlag: %v\n", ind1, ind2, procFlag)
		// Check if we've exhausted the inAr1 array
		if ind1 < len(inAr1) {

			switch {

			case inAr2[ind2] > inAr1[ind1]:
				// inAr1's element is less than inAr2's element, append it (inAr1 element) to the new array
				newArr = append(newArr, inAr1[ind1])
				// Handled the current inAr1 element, move to the next element in inAr1
				ind1++
			case inAr2[ind2] < inAr1[ind1]:
				// inAr2's element is less than inAr2's element, append it (inAr2 element) to the new array
				newArr = append(newArr, inAr2[ind2])
				// Handled the current inAr2 element, move to the next element in inAr2
				ind2++
			case inAr2[ind2] == inAr1[ind1]:
				// inAr2's element is less than inAr2's element, append it (inAr2 element) to the new array
				newArr = append(newArr, inAr1[ind1])
				newArr = append(newArr, inAr2[ind2])
				// Handled the current inAr2 element, move to the next element in inAr2
				ind1++
				ind2++
			}

		} else {
			// No more inAr1 elements to check, append the inAr2 element to the new array
			newArr = append(newArr, inAr2[ind2])
			ind2++
		}

		// Check if we've exhausted inAr2
		if ind2 >= len(inAr2) {
			// Stop the iterations
			procFlag = false
		}
	}

	// Check to see if there are any remaining elements in inAr1
	//   if so, simply append those to the new array
	if ind1 < len(inAr1) {
		newArr = append(newArr, inAr1[ind1:len(inAr1)]...)
	}

	return
}

func main() {
	fmt.Printf("Merge these two arrays\n\n%v\n%v\n\n", array1, array2)
	mergeMe(array1, array2)
	fmt.Printf("The merged array is:\n%v", newArr)
}
