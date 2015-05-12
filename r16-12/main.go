package main

import "fmt"

var (
	queue1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	qErr   error
)

// Function delElm deletes a item (identified by number) from a queue
func delElm(iNum int, inQ []int) (retQ []int, retErr error) {
	// Handle error conditions
	if len(inQ) == 0 {
		// No elements in the queue, pass back an error
		retErr = fmt.Errorf("Error - no elements in the queue, cannot delete")
	}
	if iNum >= len(inQ) || iNum < 0 {
		// No elements in the queue, pass back an error
		retErr = fmt.Errorf("Error - attempting to delete a item that is not in the queue")
	}

	if retErr != nil {
		return
	}

	// Handle special conditions
	switch iNum {
	case 1:
		retQ = inQ[1:len(inQ)]
		return
	case len(inQ) - 1:
		retQ = inQ[0 : len(inQ)-1]
		return
	}

	// Handle general conditions
	tmpArr1 := inQ[0 : iNum-1]
	tmpArr2 := inQ[iNum:len(inQ)]

	retQ = append(tmpArr1, tmpArr2...)

	return
}

func main() {
	fmt.Printf("Remove the 13th item from the queue:\n%v\n\n", queue1)
	queue1, qErr = delElm(13, queue1)
	if qErr != nil {
		fmt.Println("Error deleting the item:", qErr)
	}
	fmt.Printf("The queue is now:\n%v\n\n", queue1)

}
