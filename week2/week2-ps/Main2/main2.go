package main2

import (
	"../EdgeBitmap"
	"../WeightedEdgeArray"
	"fmt"
)

func Main2() {
	edgeBitmap := EdgeBitmap.ReadProblem2Textfile("_fe8d0202cd20a808db6a4d5d06be62f4_clustering_big.txt")

	unionFinder := WeightedEdgeArray.UnionFinder{}
	unionFinder.Init(len(edgeBitmap.Arr))

	edgeBitmap.SortAndDetermineBitSumIndicesRange()

	for edgeCost := 1; edgeCost <= edgeBitmap.GetNumBitPerNode()/2; edgeCost++ {
		existingEdgeCost := edgeBitmap.BitSumIndicesRange[edgeCost]
		if existingEdgeCost == nil {
			continue
		}
	}

	fmt.Println("asdf")
}
