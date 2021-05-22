package main

import (
	"./Knapsack"
	"fmt"
)

func main() {
	k := Knapsack.ReadTextfile("_6dfda29c18c77fd14511ba8964c2e265_knapsack1.txt")
	fmt.Println(k)
}
