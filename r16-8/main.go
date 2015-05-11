package main

import "fmt"

type lstNode struct {
	AddrMy   *lstNode
	AddrPrev *lstNode
	AddrNext *lstNode
	Value    int
}

var (
	sourceData = []int{0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0}
	node1Ptr   *lstNode

	loopCtr  = 0
	loopThrd = 500000
	loopFlag = false
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
			node1Ptr = retAddr
		}

		// Set the previous node to the current node and iterate again
		prevNode = &tmpNode
	}

	return
}

// Function prepend inserts an existing node as the first element in the list
func prependNode(inPtr *lstNode) {
	// Save the Node's current precessor and successor node pointers
	predPtr := inPtr.AddrPrev
	succPtr := inPtr.AddrNext

	// Reference the predecessor and successor nodes to each other
	if succPtr != nil {
		// Prepending a node that is NOT currently the last in the list
		predPtr.AddrNext = succPtr.AddrMy
		succPtr.AddrPrev = predPtr.AddrMy
	} else {
		// Prepending a node that IS currently the last in the list
		predPtr.AddrNext = nil
	}

	// Reference the passed node as the first node in the linked list
	inPtr.AddrPrev = nil
	inPtr.AddrNext = node1Ptr.AddrMy
	node1Ptr.AddrPrev = inPtr.AddrMy

	// We have a new first node, set the node1Ptr to the passed node
	node1Ptr = inPtr

	return
}

// Function insNode inserts a new node in the correct order (sorted ascending)
func sortList(inPtr *lstNode) (retErr error) {
	loopCtr = loopCtr + 1

	// Flags that indicate if you are on the last element and have made an insertion
	befLastNode := true

	// Iterate through the list and find the insertion point
	// Start with the first node in the list
	curNode := inPtr

	// See if the current node has a successor (list has more than one node)
	if curNode.AddrNext == nil {
		// We are already at the last node
		befLastNode = false
	}

	// Iterate while we are before the last node
	for befLastNode && loopCtr < loopThrd {
		fmt.Println("Iterating... curent node is:", curNode)

		// Is current node's value > next node's value?
		if curNode.Value > curNode.AddrNext.Value {
			// Move the next node to the begining of the list
			prependNode(curNode.AddrNext)
		} else {
			// No move need set current node to the next node's address
			curNode = curNode.AddrNext
		}

		// Is current node the last node in the list?
		if curNode.AddrNext == nil {
			befLastNode = false
		}
	}

	if loopCtr > loopThrd {
		retErr = fmt.Errorf("Error: Experienced a runaway loop")
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

	// Sort the newly populated list
	sortErr := sortList(newListPtr)
	if sortErr != nil {
		fmt.Println("Error occurred -", sortErr)
	}

	// Print the linked list
	prtErr := prtList(node1Ptr)
	if prtErr != nil {
		fmt.Println("Error after attempting to print the linked list")
	}

}
