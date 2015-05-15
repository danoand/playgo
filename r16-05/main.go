package main

import "fmt"

// Type treeNode define the structure of a node including it's:
//   - own memory location
//   - parent's memory location
//   - value
//   - left node location (if any)
//   - right node location (if any)
type treeNode struct {
	MyPtr    *treeNode
	PtrPar   *treeNode
	Value    int
	PtrLeft  *treeNode
	PtrRight *treeNode
}

var (
	myRoot        = treeNode{Value: 20}
	myInsertData1 = []int{10, 30, 5, 15, 3, 7, 17}
	myInsertData2 = []int{30, 10, 15, 5, 17, 7, 3}

	myTree      []*treeNode
	myTreeTotal int

	wrkErr error
)

// Function appendInc append a new node to the collection array and
//    adds the node value to a total
func appendInc(inPtr *treeNode) {
	// Add the element's memory address to the collection array
	myTree = append(myTree, inPtr)

	// Add the node's value to a running total
	myTreeTotal = myTreeTotal + inPtr.Value

	return
}

// Function newNodeObj takes a value and the parent's memory location and
//    creates a new child node
func newNodeObj(inVal int, inPtrPar *treeNode) (retPtr *treeNode) {
	newNode := treeNode{PtrPar: inPtrPar, Value: inVal}
	newNode.MyPtr = &newNode
	appendInc(&newNode)

	retPtr = &newNode

	return
}

// Function insertNode takes a value and inserts it into a binary tree with a given root node
func insertNode(inVal int, inRoot *treeNode) (retErr error) {
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

func main() {
	myRoot.MyPtr = &myRoot
	appendInc(&myRoot)

	fmt.Printf("Inserting values into my existing tree %v\n\n", myTree)
	for _, v := range myInsertData2 {
		wrkErr = insertNode(v, &myRoot)
		if wrkErr != nil {
			fmt.Printf("Had an error processing: %v; See: %v\n", v, wrkErr)
		}
	}

	insertMe := 18

	fmt.Printf("\nInsert this node value: %v", insertMe)
	wrkErr = insertNode(insertMe, &myRoot)
	if wrkErr != nil {
		fmt.Printf("Had an error processing: %v; See: %v\n", insertMe, wrkErr)
	}

	fmt.Printf("\n\nMy new tree looks like:\n")
	for k, v := range myTree {
		fmt.Printf("Node %v is: %v\n", k, *v)
	}

	fmt.Printf("\n\nThe sum total of all of the node values are: %v\n\n", myTreeTotal)
}
