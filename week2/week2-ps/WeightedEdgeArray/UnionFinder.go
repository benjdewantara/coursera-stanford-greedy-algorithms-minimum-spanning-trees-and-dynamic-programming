package WeightedEdgeArray

type UnionFinder struct {
	Leaders []int
}

func (f *UnionFinder) Init(numNodes int) {
	if f.Leaders == nil {
		f.Leaders = make([]int, numNodes)
	}

	for i := 0; i < len(f.Leaders); i++ {
		f.Leaders[i] = -1
	}
}

func (f *UnionFinder) AssociateNodeWithLeader(nodeFollower int, nodeLeader int) {
	f.Leaders[nodeLeader-1] = nodeLeader
	f.Leaders[nodeFollower-1] = nodeLeader
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
