package main3

import "math"

type GraphHeap struct {
    g       Graph
    indices []int
    //minEdgeCosts []float64
}

func ConvertToGraphHeap(g *Graph) GraphHeap {
    gh := GraphHeap{
        g: *g,
    }

    return gh
}

func (gh *GraphHeap) InitMinimumEdgeCostAll(
    nodeSink int) {

    gh.indices = make([]int, 0)

    for i := 0; i < len(gh.g.Edges); i++ {
        if i == nodeSink-1 {
            continue
        }

        gh.indices = append(gh.indices, gh.g.Edges[i].Tail)
        gh.g.Edges[i].MinEdgeCost = math.Inf(1)
    }
}

func (gh *GraphHeap) RefreshMinEdgeCostsOfIndex(
    nodeSinkNext int,
    nodeSinkOrigin int) {

    // nodeSinkNext is made to be nodeSinkOrigin
    edgeOfNext := &gh.g.Edges[nodeSinkNext-1]
    edgeOfOrigin := &gh.g.Edges[nodeSinkOrigin-1]

    // make the heads of nodeSinkNext become the heads of nodeSinkOrigin
    for i := 0; i < len(edgeOfNext.Heads); i++ {
        // do not do anything if nodeSinkNext equals nodeSinkOrigin
        if nodeSinkNext == nodeSinkOrigin {
            break
        }

        head := edgeOfNext.Heads[i]
        weight := edgeOfNext.Weights[i]

        if head == nodeSinkNext || head == nodeSinkOrigin {
            continue
        }

        edgeOfOrigin.Heads = append(edgeOfOrigin.Heads, head)
        edgeOfOrigin.Weights = append(edgeOfOrigin.Weights, weight)
    }

    // replace nodeSinkNext with nodeSinkOrigin in next heads
    // while doing so, also recalculate MinEdgeCost
    for i := 0; i < len(edgeOfNext.Heads); i++ {
        headNext := edgeOfNext.Heads[i]

        if headNext == nodeSinkOrigin {
            continue
        }

        var edgeOfHeadNext = &gh.g.Edges[headNext-1]

        for h := 0; h < len(edgeOfHeadNext.Heads); h++ {
            //recentlyAdded := false
            if edgeOfHeadNext.Heads[h] == nodeSinkNext {
                edgeOfHeadNext.Heads[h] = nodeSinkOrigin
                //recentlyAdded = true
            }

            if edgeOfHeadNext.Heads[h] == nodeSinkOrigin {
                if float64(edgeOfHeadNext.Weights[h]) < edgeOfHeadNext.MinEdgeCost {
                    edgeOfHeadNext.MinEdgeCost = float64(edgeOfHeadNext.Weights[h])
                }
            }
        }
    }
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
