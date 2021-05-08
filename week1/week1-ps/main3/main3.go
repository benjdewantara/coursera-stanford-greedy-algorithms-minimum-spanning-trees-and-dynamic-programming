package main3

import (
    "fmt"
)

func Main3() {
    g := Graph{}
    g.PopulateFromFile("_d4f3531eac1d289525141e95a2fea52f_edges.txt")

    g.RefreshMinimumEdgeCost(1)

    fmt.Println(g)
    fmt.Println("Hell on earth")
}
