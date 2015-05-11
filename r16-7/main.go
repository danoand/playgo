package main

import "fmt"

type lstNode struct {
	AddrMy   *lstNode
	AddrPrev *lstNode
	AddrNext *lstNode
	Value    int
}

var (
	sourceData = []int{2, 4, 6, 8, 10, 12, 14, 16}
	arrNodes   []lstNode
	node1Ptr   *lstNode
)

// Function popList creates a list from a data source array
func popList(inArr []int) (retAddr *lstNode) {
	var prevNode *lstNode

	// Iterate through the source data and create linked list nodes
	for i, v := range inArr {
		tmpNode := lstNode{Value: v}
		tmpNode.AddrMy = &tmpNode

		if i != 0 {
			// Processing a non-beginning node (i != 0), set the previous pointer
			prevNode.AddrNext = &tmpNode
			tmpNode.AddrPrev = prevNode.AddrMy
		} else {
			// Save the address of the first node in the linked list
			retAddr = &tmpNode
		}

		// Set the previous node to the current node and iterate again
		prevNode = &tmpNode
	}

	return
}

// Function insNode inserts a new node in the correct order (sorted ascending)
func insNode(inVal int, inPtr *lstNode) (retErr error) {
	tmpNode := lstNode{Value: inVal}
	tmpNode.AddrMy = &tmpNode

	// Flags that indicate if you are on the last element and have made an insertion
	befLastNode := true
	noUpdateYet := true

	// Iterate through the list and find the insertion point
	// Start with the first node in the list
	curNode := inPtr

	for befLastNode {
		fmt.Println("Iterating... curent node is:", curNode)

		// Are we iterating on the last node in the list?
		if curNode.AddrNext == nil {
			befLastNode = false
		}

		// Is the insert value less than the current node's value?
		if inVal < curNode.Value {
			// Found the insertion point, new node is inserted right before current node
			// Set new nodes previous and next pointers appropriately
			tmpNode.AddrPrev = curNode.AddrPrev
			tmpNode.AddrNext = curNode.AddrMy

			// Set the previous node's next pointer to the inserted node (confusing!)
			tmpNode.AddrPrev.AddrNext = tmpNode.AddrMy

			// Update the current node's previous pointer to reference the inserted node
			curNode.AddrPrev = tmpNode.AddrMy

			// Made a change so set the noUpdateYet to false
			noUpdateYet = false

			break
		}

		// Change the current node to the next node in the list and iterate again
		if befLastNode {
			curNode = curNode.AddrNext
		}
	}

	// See if an update has been made
	if noUpdateYet {
		// No update has been made even though we have iterated through the list -> error condition
		retErr = fmt.Errorf("ERROR - Iterated through the list but did not find an insertion point for %v", inVal)
	}

	return
}

// Function prtList prints out a linked list given a pointer to it's first node
func prtList(inNode *lstNode) (retErr error) {
	befLastNode := true
	curNode := inNode

	// Iterate through the list
	for befLastNode {
		// Are we iterating on the last node in the list?
		if curNode.AddrNext == nil {
			befLastNode = false
		}

		fmt.Printf("\n---\nAddress: %p\nPrevious: %p\nNext: %p\nValue: %v\n---\n",
			curNode.AddrMy, curNode.AddrPrev, curNode.AddrNext, curNode.Value)

		// Set current node to the next node and iterate again
		if befLastNode {
			curNode = curNode.AddrNext
		}

	}

	return
}

func main() {
	// Populate the list from the source data
	newListPtr := popList(sourceData)

	myInsert1 := 11
	// Insert the new value
	insErr := insNode(myInsert1, newListPtr)
	if insErr != nil {
		fmt.Println("Error attempting to insert a new node")
	}

	// Print the linked list
	prtErr := prtList(newListPtr)
	if prtErr != nil {
		fmt.Println("Error after attempting to print the linked list")
	}

}
