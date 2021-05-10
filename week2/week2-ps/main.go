package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	arr := ReadProblem1Textfile("_fe8d0202cd20a808db6a4d5d06be62f4_clustering1.txt")
	if arr != nil {
		fmt.Println("arr populated")
	}
	fmt.Println("Hell on earth")
}

func ReadProblem1Textfile(filepath string) [][]int {
	contentBytes, _ := ioutil.ReadFile(filepath)
	var arr [][]int

	for _, intStr := range strings.Split(string(contentBytes), "\n") {
		if arr == nil {
			//numNodes, _ := strconv.Atoi(intStr)
			arr = make([][]int, 0)
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
        arr = append(arr, edge)
	}

	return arr
}
