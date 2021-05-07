package main3

type WeightedEdge struct {
    Tail    int
    Heads   []int
    Weights []int
}

func (e WeightedEdge) IsEdgeNil() bool {
    return e.Weights == nil
}
