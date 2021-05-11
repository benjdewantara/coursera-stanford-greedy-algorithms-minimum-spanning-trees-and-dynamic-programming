package main1

import (
	"../WeightedEdgeArray"
	"fmt"
)

func Main1() {
	wEdgeArray := WeightedEdgeArray.ReadProblem1Textfile("_fe8d0202cd20a808db6a4d5d06be62f4_clustering1.txt")
	wEdgeArray.Sort()
	wEdgeArray.UnionFinder.Init(wEdgeArray.Len())

	for wEdgeArray.Len() > (4 - 1) {
		head, tail, _ := wEdgeArray.ExtractTopmost()

		headHasLeader := wEdgeArray.UnionFinder.NodeHasLeader(head)
		tailHasLeader := wEdgeArray.UnionFinder.NodeHasLeader(tail)

		headLeader := wEdgeArray.UnionFinder.GetNodeLeader(head)
		tailLeader := wEdgeArray.UnionFinder.GetNodeLeader(tail)

		if headHasLeader && tailHasLeader {
			if headLeader != tailLeader {
				wEdgeArray.UnionFinder.MergeFollowers(headLeader, tailLeader)
			}
		} else if headHasLeader {
			nodeLeader := wEdgeArray.UnionFinder.GetNodeLeader(head)
			wEdgeArray.UnionFinder.AssociateNodeWithLeader(tail, nodeLeader)
		} else if tailHasLeader {
			nodeLeader := wEdgeArray.UnionFinder.GetNodeLeader(tail)
			wEdgeArray.UnionFinder.AssociateNodeWithLeader(head, nodeLeader)
		} else {
			wEdgeArray.UnionFinder.AssociateNodeWithLeader(head, tail)
		}
	}

	maxEdgeCost := -1
	for wEdgeArray.Len() > 0 {
		_, _, cost := wEdgeArray.ExtractTopmost()
		if cost > maxEdgeCost {
			maxEdgeCost = cost
		}
	}

	fmt.Println(fmt.Sprintf("maxEdgeCost = %d", maxEdgeCost))
}
