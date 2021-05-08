package main3

type GraphHeap struct {
    g       Graph
    indices []int
}

func ConvertToGraphHeap(g *Graph) GraphHeap {
    gh := GraphHeap{
        g:       *g,
        indices: make([]int, 0),
    }

    for i := 0; i < len(gh.g.Edges); i++ {
        if gh.g.Edges[i].MinEdgeCost == 0 {
            continue
        }
        gh.indices = append(gh.indices, gh.g.Edges[i].Tail)
    }

    return gh
}

func (gh *GraphHeap) Len() int {
    return len(gh.indices)
}

func (gh *GraphHeap) Less(i, j int) bool {
    nodeIIdx := gh.indices[i]
    nodeJIdx := gh.indices[j]
    return gh.g.Edges[nodeIIdx-1].MinEdgeCost < gh.g.Edges[nodeJIdx-1].MinEdgeCost
}

func (gh *GraphHeap) Swap(i, j int) {
    gh.indices[i], gh.indices[j] = gh.indices[j], gh.indices[i]
}

func (gh *GraphHeap) Push(x interface{}) {
    edge := x.(WeightedEdge)
    gh.indices = append(gh.indices, edge.Tail)
}

func (gh *GraphHeap) Pop() interface{} {
    oldLen := len(gh.indices)

    lastIndx := gh.indices[oldLen-1]
    lastEdge := gh.g.Edges[lastIndx-1]

    gh.indices = gh.indices[0 : oldLen-1]
    return lastEdge
}
