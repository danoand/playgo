package main

import "fmt"

var (
	srcStack = []int{1, 2, 3, 4, 5, 6, 7}
	dstStack = []int{}
)

func main() {
	// Iterate through a source stack (modeled as an array)
	var tmpElm int

	// Print the srcStack
	fmt.Println("srcStack is:")
	for i := len(srcStack) - 1; i > -1; i-- {
		fmt.Printf("- %v\n", srcStack[i])
	}

	// Iterate through the source array while it has elements
	for len(srcStack) > 0 {

		// Pop an element from the source stack
		tmpElm = srcStack[len(srcStack)-1]
		srcStack = srcStack[:len(srcStack)-1]

		// Push the popped value onto the destination stack
		if tmpElm%2 != 0 {
			dstStack = append(dstStack, tmpElm)
		}
	}

	// Iterate through the dstStack and pop each element and push back onto the srcStack
	for len(dstStack) > 0 {
		fmt.Println("Processing dstStack:", dstStack)

		// Pop an element from the source stack
		tmpElm = dstStack[len(dstStack)-1]
		dstStack = dstStack[:len(dstStack)-1]

		// And push it back on to the srcStack
		srcStack = append(srcStack, tmpElm)
	}

	// Print the srcStack
	fmt.Println("srcStack is now updated to:")
	for i := len(srcStack) - 1; i > -1; i-- {
		fmt.Printf("- %v\n", srcStack[i])
	}

}
