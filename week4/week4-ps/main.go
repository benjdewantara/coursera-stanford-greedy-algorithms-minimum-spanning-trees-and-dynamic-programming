package main

import (
	"./Knapsack"
	"fmt"
)

func main() {
	k := Knapsack.ReadTextfile("_6dfda29c18c77fd14511ba8964c2e265_knapsack1.txt")
	//k := Knapsack.ReadTextfile("test_knapsack1.txt")
	optimal := k.Optimal()
	fmt.Printf("optimal = %d\n", optimal)

	termOptimal1 := k.GetTerm(k.Len()-1, k.MaxCapacity)
	fmt.Printf("termOptimal1 = %d\n", termOptimal1)

	k2 := Knapsack.ReadTextfile("_6dfda29c18c77fd14511ba8964c2e265_knapsack_big.txt")
	termOptimal2 := k2.GetTerm(k2.Len()-1, k2.MaxCapacity)
	fmt.Printf("termOptimal2 = %d\n", termOptimal2)
}
