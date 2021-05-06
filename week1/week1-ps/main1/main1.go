package main1

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func Main1() {
	// number 1
	arr := ReadProblem1Textfile("_642c2ce8f3abe387bdff636d708cdb26_jobs.txt")

	sort.Sort(ArrSortByDifferenceThenWeight(arr))

	summedWrtDifference := 0
	summedLenWrtDifference := 0

	for _, arrElem := range arr {
		summedLenWrtDifference += arrElem[1]
		summedWrtDifference += arrElem[0] * summedLenWrtDifference
	}
	fmt.Println(fmt.Sprintf("summedWrtDifference = %d", summedWrtDifference))

	// number 2
	arr2 := ReadProblem1Textfile("_642c2ce8f3abe387bdff636d708cdb26_jobs.txt")
	sort.Sort(ArrSortByRatio(arr2))

	summedWrtRatio := 0
	summedLenWrtRatio := 0

	for _, arrElem := range arr2 {
		summedLenWrtRatio += arrElem[1]
		summedWrtRatio += arrElem[0] * summedLenWrtRatio
	}
	fmt.Println(fmt.Sprintf("summedWrtRatio = %d", summedWrtRatio))
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
		difference := weight - length

		arrIndx := lineIndx - 1
		arr[arrIndx] = make([]int, 3)
		arr[arrIndx][0] = weight
		arr[arrIndx][1] = length
		arr[arrIndx][2] = difference
	}

	return arr
}

type ArrSortByDifferenceThenWeight [][]int

func (a ArrSortByDifferenceThenWeight) Len() int {
	return len(a)
}

func (a ArrSortByDifferenceThenWeight) Less(i, j int) bool {
	if a[i][2] == a[j][2] {
		return a[i][0] > a[j][0]
	}

	return a[i][2] > a[j][2]
}

func (a ArrSortByDifferenceThenWeight) Swap(i, j int) {
	elemAtJ := make([]int, 3)
	elemAtJ = a[j]
	a[j] = a[i]
	a[i] = elemAtJ
}

type ArrSortByRatio [][]int

func (a ArrSortByRatio) Len() int {
	return len(a)
}

func (a ArrSortByRatio) Less(i, j int) bool {
	ratioI := a[i][0] / a[i][1]
	ratioJ := a[j][0] / a[j][1]

	return ratioI > ratioJ
}

func (a ArrSortByRatio) Swap(i, j int) {
	elemAtJ := make([]int, 3)
	elemAtJ = a[j]
	a[j] = a[i]
	a[i] = elemAtJ
}
