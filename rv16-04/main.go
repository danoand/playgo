package main

import "fmt"

var (
	arrProdID = []string{"Noah", "Jesse", "Noah", "Mary"}
	arrSales  = []int{4, 3, 5, 6}

	totHash = map[string]int{}
)

func main() {
	fmt.Println("Here are the players' scores:")
	for i := 0; i < len(arrProdID); i++ {
		fmt.Printf("%v: %v points\n", arrProdID[i], arrSales[i])
	}

	fmt.Printf("\n\n******************************\n\n")

	fmt.Println("Here are the players' total scores:")

	// Iterate through the Product ID array
	for k, v := range arrProdID {
		totHash[v] = totHash[v] + arrSales[k]
	}

	// Print out the totHash (which is an aggregated sales report)
	for key, val := range totHash {
		fmt.Printf("Product Id: %v, Sales: %v\n", key, val)
	}
}
