package WeightedEdgeArray

type UnionFinder struct {
	Leaders         []int
	LeadersDistinct []int
}

func (f *UnionFinder) Init(numNodes int) {
	if f.Leaders == nil {
		f.Leaders = make([]int, numNodes)
	}

	for i := 0; i < len(f.Leaders); i++ {
		node := i + 1
		f.Leaders[i] = node
		f.insertToLeadersDistinct(node)
	}
}

func (f *UnionFinder) AssociateNodeWithLeader(nodeFollower int, nodeLeader int) {
	f.Leaders[nodeLeader-1] = nodeLeader
	f.Leaders[nodeFollower-1] = nodeLeader
	if nodeFollower != nodeLeader {
		f.removeFromLeadersDistinct(nodeFollower)
	}
}

func (f *UnionFinder) NodeHasLeader(nodeFollower int) bool {
	return f.Leaders[nodeFollower-1] > 0
}

func (f *UnionFinder) GetNodeLeader(nodeFollower int) int {
	return f.Leaders[nodeFollower-1]
}

func (f *UnionFinder) MergeFollowers(leader1 int, leader2 int) {
	leader := leader1
	leaderSecondary := leader2
	if leader2 < leader1 {
		leader = leader2
		leaderSecondary = leader1
	}

	for i := 0; i < len(f.Leaders); i++ {
		if f.Leaders[i] == leaderSecondary {
			f.Leaders[i] = leader
		}
	}
}

func (f *UnionFinder) insertToLeadersDistinct(node int) {
	if f.LeadersDistinct == nil {
		f.LeadersDistinct = make([]int, 0)
		f.LeadersDistinct = append(f.LeadersDistinct, node)
		return
	}

	f.LeadersDistinct = append(f.LeadersDistinct, node)
	indxToInsert := - 1
	for i := 0; i < len(f.LeadersDistinct)-1; i++ {
		if node <= f.LeadersDistinct[i+1] {
			indxToInsert = i + 1
			break
		}
	}

	if indxToInsert > -1 {
		length := len(f.LeadersDistinct)
		remaining := f.LeadersDistinct[indxToInsert : length-1]

		f.LeadersDistinct = append(f.LeadersDistinct[0:indxToInsert], f.LeadersDistinct[length-1])
		f.LeadersDistinct = append(f.LeadersDistinct, remaining...)
	}
}

func (f *UnionFinder) removeFromLeadersDistinct(nodeFollower int) {
	indxToRemove := -1
	for i := 0; i < len(f.LeadersDistinct); i++ {
		if f.LeadersDistinct[i] == nodeFollower {
			indxToRemove = i
			break
		}
	}

	if indxToRemove > -1 {
		f.LeadersDistinct = append(f.LeadersDistinct[0:indxToRemove], f.LeadersDistinct[indxToRemove+1:]...)
	}
}
