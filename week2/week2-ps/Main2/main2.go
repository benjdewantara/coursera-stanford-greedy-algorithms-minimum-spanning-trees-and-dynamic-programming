package main2

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Main2() {
	arr := ReadProblem2Textfile("_fe8d0202cd20a808db6a4d5d06be62f4_clustering_big.txt")
	if arr != nil {
	}
}

func ReadProblem2Textfile(filepath string) [][]int {
	contentBytes, _ := ioutil.ReadFile(filepath)

	var arr [][]int
	var numNodes, numBitPerNode int

	for rowIndx, intStr := range strings.Split(string(contentBytes), "\n") {
		if numNodes == 0 {
			splitHeaderRowStr := strings.Split(intStr, " ")
			numNodes, _ = strconv.Atoi(splitHeaderRowStr[0])
			numBitPerNode, _ = strconv.Atoi(splitHeaderRowStr[1])

			arr = make([][]int, numNodes)
			for i := 0; i < len(arr); i++ {
				arr[i] = make([]int, numBitPerNode)
			}

			continue
		}

		if intStr == "" {
			continue
		}

		splitStr := strings.Split(intStr, " ")
		for i := 0; i < len(splitStr); i++ {
			if splitStr[i] == "" {
				continue
			}

			arr[rowIndx-1][i], _ = strconv.Atoi(splitStr[i])
		}
	}

	return arr
}
