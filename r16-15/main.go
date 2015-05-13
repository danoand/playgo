package main

import "fmt"

var (
	myArr = []int{17, 18, 19, 22, 24, 43, 46, 50, 59, 60, 61, 63, 64, 72, 83, 87, 90, 92, 92, 95}
)

// Function findMe finds a value in an array (Go slice) and returns the position (index +1)
//    in the array
func findMe(inVal int, sIdx int, eIdx int) (retVal int, retErr error) {
	// Find the middle element (index) in the array
	tmplen := (eIdx - sIdx) + 1
	if tmplen <= 0 {
		// Error condition, have not found the element in the array
		retErr = fmt.Errorf("Can't find element %v in:\n\n%v", inVal, myArr)
		return
	}
	tmpIdx := sIdx + (tmplen / 2)

	// See if the search value is less than, equal to, or greater than the middle element
	switch {
	case inVal < myArr[tmpIdx]:
		// Call findMe on the "left" sub array
		retVal, retErr = findMe(inVal, sIdx, tmpIdx-1)
	case inVal > myArr[tmpIdx]:
		// Call findMe on the "right" sub array
		retVal, retErr = findMe(inVal, tmpIdx+1, eIdx)
	case inVal == myArr[tmpIdx]:
		retVal = tmpIdx + 1
	}

	return
}

func main() {
	findVal := 1000
	fmt.Printf("Find %v in array:\n\n%v\n\n", findVal, myArr)

	posVal, posErr := findMe(findVal, 0, len(myArr)-1)
	if posErr != nil {
		fmt.Printf("Got an error: %v\n\n", posErr)
	} else {
		fmt.Printf("The position of %v in array is: %v\n\n", findVal, posVal)
	}
}
