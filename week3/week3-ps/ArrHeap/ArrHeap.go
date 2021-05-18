package ArrHeap

import (
    "container/heap"
    "io/ioutil"
    "strconv"
    "strings"
)

type ArrHeap struct {
    ArrWeight []int
    ArrLabel  [][]int
    Depths    []int
}

func (a *ArrHeap) MergeTwoLeastWeight() {
    popped1 := heap.Pop(a).(WeightLabel)
    popped2 := heap.Pop(a).(WeightLabel)

    for i := 0; i < len(popped1.Label); i++ {
        a.Depths[popped1.Label[i]-1] += 1
    }

    for i := 0; i < len(popped2.Label); i++ {
        a.Depths[popped2.Label[i]-1] += 1
    }

    mergedNode := WeightLabel{
        Weight: popped1.Weight + popped2.Weight,
        Label:  append(popped1.Label, popped2.Label...),
    }

    heap.Push(a, mergedNode)
}

func (a *ArrHeap) Len() int {
    return len(a.ArrWeight)
}

func (a *ArrHeap) Less(i, j int) bool {
    return a.ArrWeight[i] < a.ArrWeight[j]
}

func (a *ArrHeap) Swap(i, j int) {
    a.ArrWeight[i], a.ArrWeight[j] = a.ArrWeight[j], a.ArrWeight[i]
    a.ArrLabel[i], a.ArrLabel[j] = a.ArrLabel[j], a.ArrLabel[i]
    a.Depths[i], a.Depths[j] = a.Depths[j], a.Depths[i]
}

func (a *ArrHeap) Push(x interface{}) {
    weightLabel := x.(WeightLabel)
    a.ArrWeight = append(a.ArrWeight, weightLabel.Weight)
    a.ArrLabel = append(a.ArrLabel, weightLabel.Label)
}

func (a *ArrHeap) Pop() interface{} {
    oldLen := len(a.ArrWeight)

    lastIndx := oldLen - 1
    lastElm := a.ArrWeight[lastIndx]
    lastLabel := a.ArrLabel[lastIndx]

    a.ArrWeight = a.ArrWeight[0 : oldLen-1]
    a.ArrLabel = a.ArrLabel[0 : oldLen-1]

    return WeightLabel{lastElm, lastLabel}
}

func ReadTextfile(filepath string) ArrHeap {
    contentBytes, _ := ioutil.ReadFile(filepath)
    var a = ArrHeap{}

    for lineIndx, intStr := range strings.Split(string(contentBytes), "\n") {
        if a.ArrWeight == nil {
            numNodes, _ := strconv.Atoi(intStr)
            a.ArrWeight = make([]int, numNodes)
            a.ArrLabel = make([][]int, numNodes)
            a.Depths = make([]int, numNodes)
            continue
        }

        if intStr == "" {
            continue
        }

        num, _ := strconv.Atoi(intStr)
        a.ArrWeight[lineIndx-1] = num
        a.ArrLabel[lineIndx-1] = []int{lineIndx}
        a.Depths[a.ArrLabel[lineIndx-1][0]-1] = 0
    }

    return a
}
