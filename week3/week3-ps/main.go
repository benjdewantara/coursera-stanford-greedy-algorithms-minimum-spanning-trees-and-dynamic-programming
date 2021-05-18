package main

import (
    "./ArrHeap"
    "container/heap"
    "fmt"
)

func main() {
    arrHeap := ArrHeap.ReadTextfile("_eed1bd08e2fa58bbe94b24c06a20dcdb_huffman.txt")

    heap.Init(&arrHeap)

    fmt.Println("Hell on earth")
}
