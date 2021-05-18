package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

func main() {
    arr := ReadTextfile("_eed1bd08e2fa58bbe94b24c06a20dcdb_huffman.txt")

    if arr != nil {
    }

    fmt.Println("Hell on earth")
}

func ReadTextfile(filepath string) *[]int {
    contentBytes, _ := ioutil.ReadFile(filepath)
    var arr []int

    for lineIndx, intStr := range strings.Split(string(contentBytes), "\n") {
        if arr == nil {
            numNodes, _ := strconv.Atoi(intStr)
            arr = make([]int, numNodes)
            continue
        }

        if intStr == "" {
            continue
        }

        num, _ := strconv.Atoi(intStr)
        arr[lineIndx-1] = num
    }

    return &arr
}
