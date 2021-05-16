package main2

import (
	"../EdgeBitmap"
	"../WeightedEdgeArray"
	"fmt"
)

func Main2() {
	edgeBitmap := EdgeBitmap.ReadProblem2Textfile("_fe8d0202cd20a808db6a4d5d06be62f4_clustering_big.txt")

	unionFinder := WeightedEdgeArray.UnionFinder{}
	unionFinder.Init(edgeBitmap.Len())

	// fuse all nodes that have the same coordinate
	for _, nodes := range edgeBitmap.NodeBitDecimalArr {
		if len(nodes) > 1 {
			for nodeIdx := 0; nodeIdx <= len(nodes)-2; nodeIdx++ {
				nodeAIdx, nodeBIdx := nodes[nodeIdx], nodes[nodeIdx+1]
				nodeA, nodeB := nodeAIdx+1, nodeBIdx+1

				nodeALeader, nodeBLeader := unionFinder.GetNodeLeader(nodeA), unionFinder.GetNodeLeader(nodeB)

				nodeLeader, nodeFollower := nodeALeader, nodeBLeader
				if nodeFollower < nodeLeader {
					nodeLeader, nodeFollower = nodeFollower, nodeLeader
				}

				unionFinder.AssociateNodeWithLeader(nodeFollower, nodeLeader)
				unionFinder.MergeFollowers(nodeLeader, nodeFollower)
			}
		}
	}

	edgesWithDistanceOneTwo := make([][2]int, 0)

	// get all edges with distance 1
	for nodeIdx := 0; nodeIdx < edgeBitmap.Len(); nodeIdx++ {
		nodeBits := edgeBitmap.Arr[nodeIdx]
		nodesDist1 := edgeBitmap.GetNodesWithDistanceOne(nodeBits)
		if len(nodesDist1) > 0 {
			for _, node2 := range nodesDist1 {
				if node2 < nodeIdx {
					continue
				}
				edgesWithDistanceOneTwo = append(edgesWithDistanceOneTwo, [2]int{nodeIdx, node2})
			}
		}
	}

	// get all edges with distance 2
	for nodeIdx := 0; nodeIdx < edgeBitmap.Len(); nodeIdx++ {
		nodeBits := edgeBitmap.Arr[nodeIdx]
		nodesDist2 := edgeBitmap.GetNodesWithDistanceTwo(nodeBits)
		if len(nodesDist2) > 0 {
			for _, node2 := range nodesDist2 {
				if node2 < nodeIdx {
					continue
				}
				edgesWithDistanceOneTwo = append(edgesWithDistanceOneTwo, [2]int{nodeIdx, node2})
			}
		}
	}

	for _, edge := range edgesWithDistanceOneTwo {
		nodeAIdx, nodeBIdx := edge[0], edge[1]
		nodeABits, nodeBBits := edgeBitmap.Arr[nodeAIdx], edgeBitmap.Arr[nodeBIdx]
		distance := EdgeBitmap.GetDistanceBetween(nodeABits, nodeBBits)

		if distance > 2 {
			panic(fmt.Sprintf("Unexpected distance between nodeAIdx %d and nodeBIdx %d", nodeAIdx, nodeBIdx))
		}

		nodeA, nodeB := nodeAIdx+1, nodeBIdx+1
		nodeALeader, nodeBLeader := unionFinder.GetNodeLeader(nodeA), unionFinder.GetNodeLeader(nodeB)

		if nodeALeader == nodeBLeader {
			continue
		}

		nodeLeader, nodeFollower := nodeALeader, nodeBLeader
		if nodeBLeader < nodeALeader {
			nodeLeader, nodeFollower = nodeBLeader, nodeALeader
		}

		unionFinder.AssociateNodeWithLeader(nodeFollower, nodeLeader)
		unionFinder.MergeFollowers(nodeLeader, nodeFollower)
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
