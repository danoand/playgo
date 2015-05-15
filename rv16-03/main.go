package main

import "fmt"

var (
	potSubset1 = []string{"a", "b", "c"}
	potSubset2 = []string{"a", "b", "c", "z"}
	superset   = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	workmap = map[string]int{}
)

func genMap(inArr []string) {
	// Iterate over the array
	for k, v := range inArr {
		workmap[v] = k
	}

	return
}

func isSubset(inArr []string) (retBool bool) {
	retBool = true

	for _, v := range inArr {
		if _, ok := workmap[v]; !ok {
			// An proposed key is not present in the map
			retBool = false
			return
		}
	}

	return
}

func main() {
	fmt.Println("Start processing")
	genMap(superset)

	fmt.Printf("Is %v\n\na subset of\n\n%v\n\n%v\n\n*******************\n\n",
		potSubset1, superset, isSubset(potSubset1))

	fmt.Printf("Is %v\n\na subset of\n\n%v\n\n%v\n\n*******************\n\n",
		potSubset2, superset, isSubset(potSubset2))

	fmt.Println("End processing")
}
