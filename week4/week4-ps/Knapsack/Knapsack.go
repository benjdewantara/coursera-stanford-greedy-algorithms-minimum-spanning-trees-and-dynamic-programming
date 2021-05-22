package Knapsack

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Knapsack struct {
	MaxCapacity int
	Values      []int
	Weights     []int
}

func ReadTextfile(filepath string) Knapsack {
	k := Knapsack{}

	contentBytes, _ := ioutil.ReadFile(filepath)
	for lineIndx, intStr := range strings.Split(string(contentBytes), "\n") {
		if intStr == "" {
			continue
		}

		splitStr := strings.Split(intStr, " ")

		if lineIndx == 0 {
			maxCapacity, _ := strconv.Atoi(splitStr[0])
			numItems, _ := strconv.Atoi(splitStr[1])

			k.MaxCapacity = maxCapacity
			k.Values = make([]int, numItems)
			k.Weights = make([]int, numItems)

			continue
		}

		value, _ := strconv.Atoi(splitStr[0])
		weight, _ := strconv.Atoi(splitStr[1])

		k.Values[lineIndx-1] = value
		k.Weights[lineIndx-1] = weight
	}

	return k
}
