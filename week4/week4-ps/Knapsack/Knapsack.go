package Knapsack

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Knapsack struct {
	MaxCapacity int
	Values      []int
	Weights     []int
	Terms       map[string]int
}

func (k *Knapsack) Optimal() int {
	optimalValue := -1
	for i := 0; i < k.Len(); i++ {
		for x := 0; x <= k.MaxCapacity; x++ {
			term := k.GetTerm(i, x)
			if optimalValue == -1 {
				optimalValue = term
				continue
			}

			if optimalValue < term {
				optimalValue = term
			}
		}
	}

	return optimalValue
}

func TermKeyString(i int, x int) string {
	return fmt.Sprintf("%d,%d", i, x)
}

func (k *Knapsack) Len() int {
	return len(k.Weights)
}

func (k *Knapsack) GetTerm(i int, x int) int {
	term, exists := k.Terms[TermKeyString(i, x)]
	if exists {
		return term
	}

	if i < 0 {
		return 0
	}

	option1 := k.GetTerm(i-1, x)
	option2 := 0
	if k.Weights[i] <= x {
		option2 = k.GetTerm(i-1, x-k.Weights[i]) + k.Values[i]
	}
	maxOption := option1

	if maxOption < option2 {
		maxOption = option2
	}

	k.Terms[TermKeyString(i, x)] = maxOption
	return maxOption
}

func ReadTextfile(filepath string) Knapsack {
	k := Knapsack{}
	k.Terms = make(map[string]int)

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
