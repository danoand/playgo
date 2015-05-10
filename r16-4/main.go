package main

import "fmt"

var (
	arrProdID = []int{211, 262, 211, 216}
	arrSales  = []int{4, 3, 5, 6}

	totHash = map[int]int{}
)

func main() {
	// Iterate through the Product ID array
	for k, v := range arrProdID {
		totHash[v] = totHash[v] + arrSales[k]
	}

	// Print out the totHash (which is an aggregated sales report)
	for key, val := range totHash {
		fmt.Printf("Product Id: %v, Sales: %v\n", key, val)
	}
}
