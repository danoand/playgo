package main

import "fmt"

// Type graphNode define the structure of a node including it's:
//   - own memory location
//   - parent's memory location
//   - value
//   - left node location (if any)
//   - right node location (if any)
type graphNode struct {
	MyPtr    *graphNode
	PtrPar   *graphNode
	Value    int
	PtrLeft  *graphNode
	PtrRight *graphNode
}

var (
	myRoot        = graphNode{Value: 20}
	myInsertData1 = []int{10}
	myInsertData2 = []int{30, 10, 15, 5, 17, 7, 3}

	myGraph        []*graphNode
	inspectedNodes []*graphNode
	foundNodes     []*graphNode
	manageNodes    []*graphNode

	wrkErr error
)

func storeNode(inPtr *graphNode) {

	return
}

// Function storeGraph appends a new graph node to an array which
//    can be used downstream to more easily print the graph
func storeGraph(inPtr *graphNode) {
	// Add the element's memory address to the collection array
	myGraph = append(myGraph, inPtr)

	return
}

// Function newNodeObj takes a value and the parent's memory location and
//    creates a new child node
func newNodeObj(inVal int, inPtrPar *graphNode) (retPtr *graphNode) {
	newNode := graphNode{PtrPar: inPtrPar, Value: inVal}
	newNode.MyPtr = &newNode
	storeGraph(&newNode)

	retPtr = &newNode

	return
}

// Function insertNode takes a value and inserts it into a binary tree with a given root node
func insertNode(inVal int, inRoot *graphNode) (retErr error) {
	// Set the current node for processing
	curNode := inRoot
	procFlag := true
	itrCtr := 0

	// Iterate processing (while flag is true)
	for procFlag {
		if itrCtr > 500000 {
			// Break the loop if is appears to be infinite
			procFlag = false
		}

		// Check if the insert value is less than the current node value
		if inVal < curNode.Value {
			// Yes it is, is does the current node have a "left node"?
			if curNode.PtrLeft == nil {
				// No left node, create a new node
				tmpLeftNode := newNodeObj(inVal, curNode)
				curNode.PtrLeft = tmpLeftNode

				// Done with the insert
				procFlag = false
			} else {
				// A left node is present, set the current node to the left node and iterate again
				curNode = curNode.PtrLeft
			}
		} else {
			// inVal is greater than the current node value
			if curNode.PtrRight == nil {
				// No right node, create a new node
				tmpRightNode := newNodeObj(inVal, curNode)
				curNode.PtrRight = tmpRightNode

				// Done with the insert
				procFlag = false
			} else {
				// A left node is present, set the current node to the left node and iterate again
				curNode = curNode.PtrRight
			}
		}

		itrCtr++
	}

	return
}

// Function beenInspected returns an indicator if a node has been inspected or not
func beenInspected(inNode *graphNode) (retFlag bool) {
	// Is the pointer (memory address) in the repository of inspected nodes?
	var inInsptd bool

	// Determine if node has already been inspected
	for i := 0; i < len(inspectedNodes); i++ {
		if inNode == inspectedNodes[i] {
			inInsptd = true
		}
	}

	switch {
	// If the passed pointer is nil, treat that as a node that has been inspected
	case inNode == nil:
		retFlag = true
	case inInsptd:
		retFlag = true
	}

	return
}

// Function inspectNode inspects if a node is set to value
func inspectNode(inVal int, inNode *graphNode) (retFlag bool) {
	if inNode.Value == inVal {
		foundNodes = append(foundNodes, inNode)
	}

	return
}

// Function searchNode searches a node for a value
//   If the node has branches it will invoke itself to search those first (left then right)
//   before inspecting itself for the passed value
func searchNode(inVal int, inNode *graphNode) {
	fmt.Printf("Just inside searchNode with inVal: %v, inNode: %v\n", inVal, inNode)

	// Determine if the node has been inspected
	tmpIns := beenInspected(inNode)
	tmpLft := inNode.PtrLeft
	tmpRig := inNode.PtrRight
	tmpLIn := beenInspected(tmpLft)
	tmpRIn := beenInspected(tmpRig)

	fmt.Printf("-----\nInside searchNode inNode: %v\ntmpIns: %v\ntmpLft: %v\ntmpRig: %v\ntmpLIn: %v\ntmpRIn: %v\n-----\n\n",
		inNode, tmpIns, tmpLft, tmpRig, tmpLIn, tmpRIn)

	switch {
	case tmpIns == true:
		// Node has already been searched return
		return
	case tmpIns == false:
		// Node has not been searched
		switch {
		case tmpLft == nil && tmpRig == nil:
			// Node does not have any children, inspect the node's value
			inspectNode(inVal, inNode)
			// Node has been inspected, add to our array of inspected nodes
			inspectedNodes = append(inspectedNodes, inNode)

		case tmpLIn && tmpRIn:
			// Both left and right children have been inspected - inspect current node
			inspectNode(inVal, inNode)
			// Node has been inspected, add to our array of inspected nodes
			inspectedNodes = append(inspectedNodes, inNode)

		case tmpLft != nil && !tmpLIn:
			// Node has a left child and it has not yet been inspected - search left child
			searchNode(inVal, tmpLft)

		case tmpLft == nil || tmpLIn:
			// Node does not have a left child or left child has already been inspected
			//   - search the right node
			searchNode(inVal, tmpRig)
		}
	}

	return
}

func popManage() (retVal *graphNode) {
	// "Pop" or return the first element in the array (stack) if it exists
	if len(manageNodes) == 0 {
		// No element in the stack, simply return
		return
	}

	// Fetch first element
	retVal = manageNodes[0]
	// Remove first element
	manageNodes = manageNodes[1:len(manageNodes)]

	return
}

func searchBreadth(inVal int, inNode *graphNode) {
	fmt.Printf("Currently searching %v which has value: %v\n", inNode.MyPtr, inNode.Value)

	// If inNode = nil, simply return no further action to take
	if inNode == nil {
		return
	}

	// Inspect passed node
	inspectNode(inVal, inNode)

	// If the node has a left child node then push the node onto the search stack (manageNodes)
	if inNode.PtrLeft != nil {
		manageNodes = append(manageNodes, inNode.PtrLeft)
	}

	// If the node has a right child node then push the node onto the search stack (manageNodes)
	if inNode.PtrRight != nil {
		manageNodes = append(manageNodes, inNode.PtrRight)
	}

	// Pop the next item off of the search stack and search it
	if len(manageNodes) != 0 {
		searchBreadth(inVal, popManage())
	}

	return
}

func main() {
	myRoot.MyPtr = &myRoot
	storeGraph(&myRoot)

	fmt.Printf("Inserting values into my existing tree %v\n\n", myGraph)
	for _, v := range myInsertData2 {
		wrkErr = insertNode(v, &myRoot)
		if wrkErr != nil {
			fmt.Printf("Had an error processing: %v; See: %v\n", v, wrkErr)
		}
	}

	fmt.Printf("\n\nMy new tree looks like:\n")
	for k, v := range myGraph {
		fmt.Printf("Node %v is: %v\n", k, *v)
	}

	searchVal := 3

	// Trigger the deep search
	searchBreadth(searchVal, &myRoot)
	fmt.Printf("Found %v in these node(s):\n\n%v\n\n", searchVal, foundNodes)

}
