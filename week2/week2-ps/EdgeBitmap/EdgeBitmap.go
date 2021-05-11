package EdgeBitmap

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type EdgeBitmap struct {
	Arr    [][]int
	BitSum []int
}

func (b *EdgeBitmap) GetNumBitPerNode() int {
	return len(b.Arr[0])
}

func (b *EdgeBitmap) GetNumNodes() int {
	return len(b.Arr)
}

func ReadProblem2Textfile(filepath string) EdgeBitmap {
	contentBytes, _ := ioutil.ReadFile(filepath)

	edgeBitmap := EdgeBitmap{}
	var numNodes, numBitPerNode int

	for rowIndx, intStr := range strings.Split(string(contentBytes), "\n") {
		if numNodes == 0 {
			splitHeaderRowStr := strings.Split(intStr, " ")
			numNodes, _ = strconv.Atoi(splitHeaderRowStr[0])
			numBitPerNode, _ = strconv.Atoi(splitHeaderRowStr[1])

			edgeBitmap.Arr = make([][]int, numNodes)
			for i := 0; i < len(edgeBitmap.Arr); i++ {
				edgeBitmap.Arr[i] = make([]int, numBitPerNode)
			}

			edgeBitmap.BitSum = make([]int, numNodes)
			continue
		}

		if intStr == "" {
			continue
		}

		splitStr := strings.Split(intStr, " ")
		summed := 0
		for i := 0; i < len(splitStr); i++ {
			if splitStr[i] == "" {
				continue
			}

			bitUnit, _ := strconv.Atoi(splitStr[i])
			edgeBitmap.Arr[rowIndx-1][i] = bitUnit
			summed += bitUnit
		}

		edgeBitmap.BitSum[rowIndx-1] = summed
	}

	return edgeBitmap
}
