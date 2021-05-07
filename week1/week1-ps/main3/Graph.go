package main3

import (
    "io/ioutil"
    "strconv"
    "strings"
)

type WeightedEdge struct {
    Tail    int
    Heads   []int
    Weights []int
}

func (e WeightedEdge) IsEdgeNil() bool {
    return e.Weights == nil
}

type Graph struct {
    Edges []WeightedEdge
}

func (g *Graph) GetEdgeAt(node int) *WeightedEdge {
    return &g.Edges[node-1]
}

func (g *Graph) PopulateFromFile(filepath string) {
    contentBytes, _ := ioutil.ReadFile(filepath)
    for lineIndx, lineStr := range strings.Split(string(contentBytes), "\n") {
        if lineIndx == 0 {
            strSplit := strings.Split(lineStr, " ")
            numVertices, _ := strconv.Atoi(strSplit[0])
            _, _ = strconv.Atoi(strSplit[1])
            g.Edges = make([]WeightedEdge, numVertices)
            continue
        }

        if lineStr == "" {
            continue
        }

        lineStrSplit := strings.Split(lineStr, " ")

        tail, _ := strconv.Atoi(lineStrSplit[0])

        var edgeTail *WeightedEdge
        if g.Edges[tail-1].IsEdgeNil() {
            edgeTail = &WeightedEdge{
                Tail:    tail,
                Heads:   make([]int, 0),
                Weights: make([]int, 0),
            }
            g.Edges[tail-1] = *edgeTail
        }
        edgeTail = &g.Edges[tail-1]

        head, _ := strconv.Atoi(lineStrSplit[1])
        var edgeHead *WeightedEdge
        if g.Edges[head-1].IsEdgeNil() {
            edgeHead = &WeightedEdge{
                Tail:    head,
                Heads:   make([]int, 0),
                Weights: make([]int, 0),
            }
            g.Edges[head-1] = *edgeHead
        }
        edgeHead = &g.Edges[head-1]

        weight, _ := strconv.Atoi(lineStrSplit[2])

        edgeTail.Heads = append(edgeTail.Heads, head)
        edgeTail.Weights = append(edgeTail.Weights, weight)

        edgeHead.Heads = append(edgeHead.Heads, tail)
        edgeHead.Weights = append(edgeHead.Weights, weight)
    }
}
