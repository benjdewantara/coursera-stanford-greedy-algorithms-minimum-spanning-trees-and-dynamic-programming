package main1

import (
	"../WeightedEdgeArray"
	"fmt"
)

func Main1() {
	wEdgeArray := WeightedEdgeArray.ReadProblem1Textfile("_fe8d0202cd20a808db6a4d5d06be62f4_clustering1.txt")
	wEdgeArray.Sort()
	wEdgeArray.UnionFinder.Init(wEdgeArray.NumNodes)

	for len(wEdgeArray.UnionFinder.LeadersDistinct) > 4 {
		head, tail, _ := wEdgeArray.ExtractTopmost()

		headLeader := wEdgeArray.UnionFinder.GetNodeLeader(head)
		tailLeader := wEdgeArray.UnionFinder.GetNodeLeader(tail)

		if headLeader == tailLeader {
			continue
		}

		nodeLeader, nodeFollower := headLeader, tailLeader
		if tailLeader < headLeader {
			nodeLeader, nodeFollower = tailLeader, headLeader
		}
		wEdgeArray.UnionFinder.AssociateNodeWithLeader(nodeFollower, nodeLeader)
		wEdgeArray.UnionFinder.MergeFollowers(nodeLeader, nodeFollower)
	}

	minDistance := -1
	maxDistance := -1
	for wEdgeArray.Len() > 0 {
		head, tail, cost := wEdgeArray.ExtractTopmost()
		headLeader := wEdgeArray.UnionFinder.GetNodeLeader(head)
		tailLeader := wEdgeArray.UnionFinder.GetNodeLeader(tail)

		if headLeader == tailLeader {
			continue
		}

		if minDistance == -1 {
			minDistance = cost
		}

		if cost > maxDistance {
			maxDistance = cost
		}
	}

	fmt.Println(fmt.Sprintf("minDistance = %d", minDistance))
	fmt.Println(fmt.Sprintf("maxDistance = %d", maxDistance))
}
