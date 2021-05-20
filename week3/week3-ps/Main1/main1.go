package Main1

import (
    "../ArrHeap"
    "container/heap"
    "fmt"
)

func Main1() {
    // a := ArrHeap.ReadTextfile("test_huffman1.txt")
    a := ArrHeap.ReadTextfile("_eed1bd08e2fa58bbe94b24c06a20dcdb_huffman.txt")

    heap.Init(&a)
    for a.Len() >= 2 {
        a.MergeTwoLeastWeight()
    }

    maxDepth := -1
    minDepth := -1

    for i := 0; i < len(a.Depths); i++ {
        if i == 0 {
            maxDepth, minDepth = a.Depths[i], a.Depths[i]
            continue
        }

        if a.Depths[i] < minDepth {
            minDepth = a.Depths[i]
        }

        if maxDepth < a.Depths[i] {
            maxDepth = a.Depths[i]
        }
    }

    fmt.Printf("maxDepth = %d\n", maxDepth)
    fmt.Printf("minDepth = %d\n", minDepth)
}
