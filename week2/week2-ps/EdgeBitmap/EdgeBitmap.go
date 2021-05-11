package EdgeBitmap

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type EdgeBitmap struct {
	Arr                [][]int
	BitSum             []int
	BitSumIndicesRange map[int][]int
}

func (b *EdgeBitmap) Len() int {
	return len(b.Arr)
}

func (b *EdgeBitmap) Less(i, j int) bool {
	return b.BitSum[i] < b.BitSum[j]
}

func (b *EdgeBitmap) Swap(i, j int) {
	b.BitSum[i], b.BitSum[j] = b.BitSum[j], b.BitSum[i]
	b.Arr[i], b.Arr[j] = b.Arr[j], b.Arr[i]
}

func (b *EdgeBitmap) GetNumBitPerNode() int {
	return len(b.Arr[0])
}

func (b *EdgeBitmap) SortAndDetermineBitSumIndicesRange() {
	sort.Sort(b)

	currentBitSum := 0
	startIndx, endIndx := 0, 0
	isCapturingRange := false
	for endIndx < b.Len() {
		if !isCapturingRange {
			if currentBitSum < b.BitSum[endIndx] {
				currentBitSum++
				continue
			} else if currentBitSum == b.BitSum[endIndx] {
				isCapturingRange = true
			}
		} else {
			if currentBitSum < b.BitSum[endIndx] {
				endIndxDetermined := endIndx - 1
				if b.BitSumIndicesRange == nil {
					b.BitSumIndicesRange = make(map[int][]int)
				}
				b.BitSumIndicesRange[currentBitSum] = []int{startIndx, endIndxDetermined}
				currentBitSum++
				isCapturingRange = false
				startIndx = endIndx
				continue
			}
		}

		if !isCapturingRange {
			startIndx++
		}
		endIndx++
	}
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
