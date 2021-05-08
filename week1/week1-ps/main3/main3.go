package main3

import (
    "container/heap"
    "fmt"
)

func Main3() {
    g := Graph{}
    g.PopulateFromFile("_d4f3531eac1d289525141e95a2fea52f_edges.txt")

    g.RefreshMinimumEdgeCostAll(1)

    gHeap := ConvertToGraphHeap(&g)
    heap.Init(&gHeap)

    summed := 0

    for gHeap.Len() > 0 {
        e := heap.Pop(&gHeap).(WeightedEdge)
        fmt.Println(e)

        summed += int(e.MinEdgeCost)
        break
    }

    fmt.Println("Hell on earth")
}
