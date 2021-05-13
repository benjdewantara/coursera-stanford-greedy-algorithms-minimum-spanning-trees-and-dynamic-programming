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

	maxSpacing := -1
	for wEdgeArray.Len() > 0 {
		head, tail, cost := wEdgeArray.ExtractTopmost()
		headLeader := wEdgeArray.UnionFinder.GetNodeLeader(head)
		tailLeader := wEdgeArray.UnionFinder.GetNodeLeader(tail)

		if headLeader == tailLeader {
			continue
		}

		if cost > maxSpacing {
			maxSpacing = cost
		}
	}

	fmt.Println(fmt.Sprintf("maxSpacing = %d", maxSpacing))
}
