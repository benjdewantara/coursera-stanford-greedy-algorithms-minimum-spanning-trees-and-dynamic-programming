package main3

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Main3() {

}

func ReadProblem3Textfile(filepath string) [][]int {
	contentBytes, _ := ioutil.ReadFile(filepath)
	var arr [][]int

	for lineIndx, intStr := range strings.Split(string(contentBytes), "\n") {
		if arr == nil {
			numLen, _ := strconv.Atoi(intStr)
			arr = make([][]int, numLen)
			continue
		}

		if intStr == "" {
			continue
		}

		splitStr := strings.Split(intStr, " ")
		weight, _ := strconv.Atoi(splitStr[0])
		length, _ := strconv.Atoi(splitStr[1])
		difference := weight - length

		arrIndx := lineIndx - 1
		arr[arrIndx] = make([]int, 3)
		arr[arrIndx][0] = weight
		arr[arrIndx][1] = length
		arr[arrIndx][2] = difference
	}

	return arr
}
