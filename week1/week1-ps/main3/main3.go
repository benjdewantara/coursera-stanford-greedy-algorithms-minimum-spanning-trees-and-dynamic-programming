package main3

import (
    "container/heap"
    "fmt"
)

func Main3() {
    g := Graph{}
    g.PopulateFromFile("_d4f3531eac1d289525141e95a2fea52f_edges.txt")

    //g.RefreshMinimumEdgeCostAll(1)

    nodeSinkOrigin := 1

    gHeap := ConvertToGraphHeap(&g)
    gHeap.InitMinimumEdgeCostAll(1)
    gHeap.RefreshMinEdgeCostsOfIndex(nodeSinkOrigin, nodeSinkOrigin)
    heap.Init(&gHeap)

    summed := 0

    for gHeap.Len() > 0 {
        e := heap.Pop(&gHeap).(WeightedEdge)

        summed += int(e.MinEdgeCost)

        gHeap.RefreshMinEdgeCostsOfIndex(e.Tail, nodeSinkOrigin)
        heap.Init(&gHeap)
    }

    fmt.Println(fmt.Sprintf("summed = %d", summed))
}
