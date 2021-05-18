package main

import (
    "./ArrHeap"
    "container/heap"
    "fmt"
)

func main() {
    a := ArrHeap.ReadTextfile("_eed1bd08e2fa58bbe94b24c06a20dcdb_huffman.txt")

    heap.Init(&a)
    for a.Len() >= 2 {
        a.MergeTwoLeastWeight()
    }

    fmt.Println("Hell on earth")
}
