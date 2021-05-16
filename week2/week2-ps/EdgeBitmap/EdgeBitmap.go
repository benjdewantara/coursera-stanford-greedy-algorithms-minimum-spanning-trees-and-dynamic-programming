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
	NodeBitDecimalArr  map[uint][]int
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

func (b *EdgeBitmap) Sort() {
	sort.Sort(b)
}

func (b *EdgeBitmap) DetermineBitSumIndicesRange() {
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
			} else if endIndx == b.Len()-1 {
				b.BitSumIndicesRange[currentBitSum] = []int{startIndx, endIndx}
			}
		}

		if !isCapturingRange {
			startIndx++
		}
		endIndx++
	}
}

func (b *EdgeBitmap) GetNodesWithDistanceOne(nodeBits []int) []int {
	nodes := make([]int, 0)

	baseTen := convertNodeBitsToDecimal(nodeBits)
	for i := 0; i < b.GetNumBitPerNode(); i++ {
		pow := (len(nodeBits) - 1) - i
		term := uint(1 << pow)

		baseTenShifted := baseTen ^ term
		existingDistancedNodes := b.NodeBitDecimalArr[baseTenShifted]
		if existingDistancedNodes != nil {
			nodes = append(nodes, existingDistancedNodes...)
		}
	}

	return nodes
}

func (b *EdgeBitmap) GetNodesWithDistanceTwo(nodeBits []int) []int {
	nodes := make([]int, 0)

	baseTen := convertNodeBitsToDecimal(nodeBits)
	for i := 0; i < b.GetNumBitPerNode(); i++ {
		for j := i + 1; j < b.GetNumBitPerNode(); j++ {
			pow := (len(nodeBits) - 1) - i
			term := uint(1 << pow)

			pow = (len(nodeBits) - 1) - j
			term += uint(1 << pow)
			baseTenShifted := baseTen ^ term

			existingDistancedNodes := b.NodeBitDecimalArr[baseTenShifted]
			if existingDistancedNodes != nil {
				nodes = append(nodes, existingDistancedNodes...)
			}
		}
	}

	return nodes
}

func (b *EdgeBitmap) GetNodesWithDistanceThree(nodeBits []int) []int {
	nodes := make([]int, 0)

	baseTen := convertNodeBitsToDecimal(nodeBits)
	for i := 0; i < b.GetNumBitPerNode(); i++ {
		for j := i + 1; j < b.GetNumBitPerNode(); j++ {
			for k := j + 1; k < b.GetNumBitPerNode(); k++ {
				pow := (len(nodeBits) - 1) - i
				term := uint(1 << pow)

				pow = (len(nodeBits) - 1) - j
				term += uint(1 << pow)

				pow = (len(nodeBits) - 1) - k
				term += uint(1 << pow)
				baseTenShifted := baseTen ^ term

				existingDistancedNodes := b.NodeBitDecimalArr[baseTenShifted]
				if existingDistancedNodes != nil {
					nodes = append(nodes, existingDistancedNodes...)
				}
			}
		}
	}

	return nodes
}

func convertNodeBitsToDecimal(nodeBits []int) uint {
	var num uint = 0
	for i := 0; i < len(nodeBits); i++ {
		if nodeBits[i] == 1 {
			pow := (len(nodeBits) - 1) - i
			term := uint(1 << pow)
			num += term
		}
	}
	return num
}

func ReadProblem2Textfile(filepath string) EdgeBitmap {
	contentBytes, _ := ioutil.ReadFile(filepath)

	edgeBitmap := EdgeBitmap{}
	var numNodes, numBitPerNode int
	edgeBitmap.NodeBitDecimalArr = make(map[uint][]int)

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
		var summedBitDecimal uint = 0
		for i := 0; i < len(splitStr); i++ {
			if splitStr[i] == "" {
				continue
			}

			bitUnit, _ := strconv.Atoi(splitStr[i])
			edgeBitmap.Arr[rowIndx-1][i] = bitUnit
			summed += bitUnit

			if bitUnit == 1 {
				var pow = uint(23 - i)
				var term = uint(1 << pow)
				summedBitDecimal += term
			}
		}

		nodeIndx := rowIndx - 1
		edgeBitmap.BitSum[nodeIndx] = summed
		nodesWithSummedBitDecimal := edgeBitmap.NodeBitDecimalArr[summedBitDecimal]
		if nodesWithSummedBitDecimal == nil {
			nodesWithSummedBitDecimal = make([]int, 0)
		}
		edgeBitmap.NodeBitDecimalArr[summedBitDecimal] = append(nodesWithSummedBitDecimal, nodeIndx)
	}

	return edgeBitmap
}

func GetDistanceBetween(nodeBitsA []int, nodeBitsB []int) int {
	summed := 0
	for i := 0; i < len(nodeBitsA); i++ {
		if nodeBitsA[i] != nodeBitsB[i] {
			summed++
		}
	}
	return summed
}
