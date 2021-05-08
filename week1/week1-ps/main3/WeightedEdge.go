package main3

type WeightedEdge struct {
    Tail        int
    Heads       []int
    Weights     []int
    MinEdgeCost float64
}

func (e WeightedEdge) IsEdgeNil() bool {
    return e.Weights == nil
}
