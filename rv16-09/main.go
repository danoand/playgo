package main

import "fmt"

var (
	srcStack = []int{1, 2, 3}
	dstStack []int
)

func main() {
	// Iterate through a source stack (modeled as an array)
	var tmpElm int

	// Iterate through the source array while it has elements
	for len(srcStack) > 0 {
		fmt.Println("Processing srcStack:", srcStack)

		// Pop an element from the source stack
		tmpElm = srcStack[len(srcStack)-1]
		srcStack = srcStack[:len(srcStack)-1]

		// Push the popped value onto the destination stack
		dstStack = append(dstStack, tmpElm)

	}

	// Print out resulting slice
	fmt.Println("Printing out the resulting array:")
	for _, v := range dstStack {
		fmt.Printf("- %v\n", v)
	}

	fmt.Println("Confirm that the len(srcStack) is zero:", len(srcStack))

}
