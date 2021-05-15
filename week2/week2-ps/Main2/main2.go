package main2

import (
	"../EdgeBitmap"
	"../WeightedEdgeArray"
	"fmt"
)

func Main2() {
	//DISTANCE_AT_MOST := 2

	edgeBitmap := EdgeBitmap.ReadProblem2Textfile("_fe8d0202cd20a808db6a4d5d06be62f4_clustering_big.txt")

	unionFinder := WeightedEdgeArray.UnionFinder{}
	unionFinder.Init(edgeBitmap.Len())

	edgeBitmap.Sort()
	edgeBitmap.DetermineBitSumIndicesRange()

	for targetCost := 1; targetCost <= 2; targetCost++ {
		for bitSumNodeA := 0; bitSumNodeA <= edgeBitmap.GetNumBitPerNode(); bitSumNodeA++ {
			for bitSumNodeB := bitSumNodeA; bitSumNodeB <= edgeBitmap.GetNumBitPerNode(); bitSumNodeB++ {
				if !CanTwoBitSumsProduceDistance(bitSumNodeA, bitSumNodeB, targetCost) {
					continue
				}

				fmt.Println(fmt.Sprintf("targetCost %d can be produced by %d\t%d",
					targetCost,
					bitSumNodeA,
					bitSumNodeB))

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
						nodeABits := edgeBitmap.Arr[nodeAIdx]
						nodeBBits := edgeBitmap.Arr[nodeBIdx]
						distance := EdgeBitmap.GetDistanceBetween(nodeABits, nodeBBits)

						if !(distance == targetCost) {
							continue
						}
						nodeA := nodeAIdx + 1
						nodeB := nodeBIdx + 1

						nodeALeader := unionFinder.GetNodeLeader(nodeA)
						nodeBLeader := unionFinder.GetNodeLeader(nodeB)

						if nodeALeader == nodeBLeader {
							continue
						}

						nodeLeader, nodeFollower := nodeALeader, nodeBLeader
						if nodeFollower < nodeLeader {
							nodeLeader, nodeFollower = nodeFollower, nodeLeader
						}
						unionFinder.AssociateNodeWithLeader(nodeFollower, nodeLeader)
						unionFinder.MergeFollowers(nodeLeader, nodeFollower)
					}
				}
			}
		}
	}

	fmt.Println(fmt.Sprintf("len(unionFinder.LeadersDistinct) = %d", len(unionFinder.LeadersDistinct)))
}

func CanTwoBitSumsProduceDistance(bitSumA int, bitSumB int, distanceTarget int) bool {
	bitSumSmall, bitSumLarge := bitSumA, bitSumB
	if bitSumLarge < bitSumSmall {
		bitSumSmall, bitSumLarge = bitSumLarge, bitSumSmall
	}

	for bitSum := 0; bitSum <= bitSumSmall; bitSum++ {
		bitSumRight := bitSum
		bitSumLeft := bitSumSmall - bitSumRight
		distance := (bitSumLarge - bitSumLeft) + bitSumRight
		if distance == distanceTarget {
			return true
		}
	}

	return false
}

func IsPossibleToFormDistanceAtMost(a int, b int, distanceAtMost int) bool {
	whenAdded := a+b <= distanceAtMost

	subtracted := a - b
	if subtracted < 0 {
		subtracted = -subtracted
	}
	whenSubtracted := subtracted <= distanceAtMost

	return whenAdded || whenSubtracted
}
