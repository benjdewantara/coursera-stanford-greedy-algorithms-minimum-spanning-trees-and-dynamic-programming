package WeightedEdgeArray

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type WeightedEdgeArray struct {
	NumNodes int
	RawArr   [][]int
}

func (w *WeightedEdgeArray) FindEdgeOfHeadAndTail(head int, tail int) []int {
	for i := 0; i < len(w.RawArr); i++ {
		headNode := w.RawArr[i][0]
		tailNode := w.RawArr[i][1]
		if headNode == head && tailNode == tail {
			return w.RawArr[i]
		} else if headNode == tail && tailNode == head {
			return w.RawArr[i]
		}
	}
	return nil
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
