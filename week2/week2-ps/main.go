package main

import (
	"./WeightedEdgeArray"
	"fmt"
)

func main() {
	weightedEdgeArray := WeightedEdgeArray.ReadProblem1Textfile("_fe8d0202cd20a808db6a4d5d06be62f4_clustering1.txt")
	weightedEdgeArray.Sort()

	if weightedEdgeArray.RawArr != nil {
		fmt.Println("arr populated")
	}
	fmt.Println("Hell on earth")
}
