package ArrHeap

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

type ArrHeap struct {
    ArrWeight []int
    ArrLabel  []string
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
}

func (a *ArrHeap) Push(x interface{}) {
    weightLabel := x.(WeightLabel)
    a.ArrWeight = append(a.ArrWeight, weightLabel.Weight)
    a.ArrLabel = append(a.ArrLabel, weightLabel.Label)
}

func (a *ArrHeap) Pop() interface{} {
    oldLen := len(a.ArrWeight)

    lastIndx := a.ArrWeight[oldLen-1]
    lastElm := a.ArrWeight[lastIndx-1]
    lastLabel := a.ArrLabel[lastIndx-1]

    a.ArrWeight = a.ArrWeight[0 : oldLen-1]
    a.ArrLabel = a.ArrLabel[0 : oldLen-1]
    return WeightLabel{lastElm, lastLabel}
}

func ReadTextfile(filepath string) ArrHeap {
    contentBytes, _ := ioutil.ReadFile(filepath)
    var arrHeap = ArrHeap{}

    for lineIndx, intStr := range strings.Split(string(contentBytes), "\n") {
        if arrHeap.ArrWeight == nil {
            numNodes, _ := strconv.Atoi(intStr)
            arrHeap.ArrWeight = make([]int, numNodes)
            arrHeap.ArrLabel = make([]string, numNodes)
            continue
        }

        if intStr == "" {
            continue
        }

        num, _ := strconv.Atoi(intStr)
        arrHeap.ArrWeight[lineIndx-1] = num
        arrHeap.ArrLabel[lineIndx-1] = fmt.Sprintf("%d", lineIndx)
    }

    return arrHeap
}
