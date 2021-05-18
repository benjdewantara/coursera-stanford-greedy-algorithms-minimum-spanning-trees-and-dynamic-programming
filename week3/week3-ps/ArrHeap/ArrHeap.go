package ArrHeap

import (
    "io/ioutil"
    "strconv"
    "strings"
)

type ArrHeap struct {
    Arr []int
}

func (a *ArrHeap) Len() int {
    return len(a.Arr)
}

func (a *ArrHeap) Less(i, j int) bool {
    return a.Arr[i] < a.Arr[j]
}

func (a *ArrHeap) Swap(i, j int) {
    a.Arr[i], a.Arr[j] = a.Arr[j], a.Arr[i]
}

func (a *ArrHeap) Push(x interface{}) {
    xInt := x.(int)
    a.Arr = append(a.Arr, xInt)
}

func (a *ArrHeap) Pop() interface{} {
    oldLen := len(a.Arr)

    lastIndx := a.Arr[oldLen-1]
    lastElm := a.Arr[lastIndx-1]

    a.Arr = a.Arr[0 : oldLen-1]
    return lastElm
}

func ReadTextfile(filepath string) ArrHeap {
    contentBytes, _ := ioutil.ReadFile(filepath)
    var arrHeap = ArrHeap{}

    for lineIndx, intStr := range strings.Split(string(contentBytes), "\n") {
        if arrHeap.Arr == nil {
            numNodes, _ := strconv.Atoi(intStr)
            arrHeap.Arr = make([]int, numNodes)
            continue
        }

        if intStr == "" {
            continue
        }

        num, _ := strconv.Atoi(intStr)
        arrHeap.Arr[lineIndx-1] = num
    }

    return arrHeap
}
