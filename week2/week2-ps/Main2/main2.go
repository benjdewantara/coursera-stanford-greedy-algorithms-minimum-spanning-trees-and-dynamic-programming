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

	for bitSumNodeA := 0; bitSumNodeA < edgeBitmap.GetNumBitPerNode()/2; bitSumNodeA++ {
		for bitSumNodeB := bitSumNodeA; bitSumNodeB < edgeBitmap.GetNumBitPerNode()/2; bitSumNodeB++ {
			bitSumNodeARange := edgeBitmap.BitSumIndicesRange[bitSumNodeA]
			bitSumNodeBRange := edgeBitmap.BitSumIndicesRange[bitSumNodeB]

			bitSumNodeAExists := bitSumNodeARange != nil
			bitSumNodeBExists := bitSumNodeBRange != nil

			if !bitSumNodeAExists {
				break
			} else if !bitSumNodeBExists {
				continue
			}

			nodeAStartIndx, nodeAEndIndx := bitSumNodeARange[0], bitSumNodeARange[1]
			nodeBStartIndx, nodeBEndIndx := bitSumNodeBRange[0], bitSumNodeBRange[1]
			for nodeAIdx := nodeAStartIndx; nodeAIdx <= nodeAEndIndx; nodeAIdx++ {
				for nodeBIdx := nodeBStartIndx; nodeBIdx <= nodeBEndIndx; nodeBIdx++ {
					if nodeAIdx == nodeBIdx {
						continue
					}


				}
			}
		}

		if edgeBitmap.BitSumIndicesRange[bitSumNodeA] == nil {
			continue
		}

		//bitSumNodeB1 := edgeCostTarget - bitSumNodeA
		//bitSumNodeB2 := edgeCostTarget + bitSumNodeA
/*		if edgeBitmap.BitSumIndicesRange[bitSumNodeB1] == nil {
			continue
		} else if bitSumNodeB1 > edgeBitmap.GetNumBitPerNode() {
			break
		}
*/
	}

	fmt.Println("asdf")
}
