package WeightedEdgeArray

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type WeightedEdgeArray struct {
	NumNodes    int
	RawArr      [][]int
	UnionFinder UnionFinder
}

func (w *WeightedEdgeArray) Len() int {
	return len(w.RawArr)
}

func (w *WeightedEdgeArray) Sort() {
	sort.Sort(w)
}

func ReadProblem1Textfile(filepath string) WeightedEdgeArray {
	contentBytes, _ := ioutil.ReadFile(filepath)
	weightedEdgeArray := WeightedEdgeArray{}

	var numNodes int

	for _, intStr := range strings.Split(string(contentBytes), "\n") {
		if weightedEdgeArray.RawArr == nil {
			numNodes, _ = strconv.Atoi(intStr)
			weightedEdgeArray.NumNodes = numNodes
			weightedEdgeArray.RawArr = make([][]int, 0)
			continue
		}

		if intStr == "" {
			continue
		}

		splitStr := strings.Split(intStr, " ")
		node1, _ := strconv.Atoi(splitStr[0])
		node2, _ := strconv.Atoi(splitStr[1])
		cost, _ := strconv.Atoi(splitStr[2])

		edge := []int{node1, node2, cost}
		weightedEdgeArray.RawArr = append(weightedEdgeArray.RawArr, edge)
	}

	return weightedEdgeArray
}

func (w *WeightedEdgeArray) Less(i, j int) bool {
	return w.RawArr[i][2] < w.RawArr[j][2]
}

func (w *WeightedEdgeArray) Swap(i, j int) {
	w.RawArr[i], w.RawArr[j] = w.RawArr[j], w.RawArr[i]
}

func (w *WeightedEdgeArray) ExtractTopmost() (int, int, int) {
	topmost := w.RawArr[0]
	w.RawArr = w.RawArr[1:]
	return topmost[0], topmost[1], topmost[2]
}
