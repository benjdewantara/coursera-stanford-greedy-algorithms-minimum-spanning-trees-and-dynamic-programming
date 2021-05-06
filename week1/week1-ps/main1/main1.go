package main1

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Main1() {
	arr := ReadProblem1Textfile("_642c2ce8f3abe387bdff636d708cdb26_jobs.txt")
	fmt.Println(arr)

	fmt.Println("Hell on earth")
}

func ReadProblem1Textfile(filepath string) [][]int {
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

		arrIndx := lineIndx - 1
		arr[arrIndx] = make([]int, 2)
		arr[arrIndx][0] = weight
		arr[arrIndx][1] = length
	}

	return arr
}
