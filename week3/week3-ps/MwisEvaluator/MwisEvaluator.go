package MwisEvaluator

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type MwisEvaluator struct {
	MwisResultCached map[string][]int
	Weights          []int
	Vertices         []int
}

func (e *MwisEvaluator) Mwis(vertices []int) []int {
	cachedResult := e.MwisResultCached[ArrUtility(vertices).StringJoined()]
	if cachedResult != nil {
		return cachedResult
	}

	if len(vertices) == 2 {
		if e.GetWeightAtVertex(vertices[0]) < e.GetWeightAtVertex(vertices[1]) {
			return vertices[1:2]
		} else {
			return vertices[0:1]
		}
	} else if len(vertices) == 1 {
		return vertices[0:1]
	}

	lastIndx := len(vertices) - 1
	withoutVn := make([]int, 0)
	withoutVn = append(withoutVn, e.Mwis(vertices[0:lastIndx])...)

	withVn := make([]int, 0)
	withVn = append(withVn, e.Mwis(vertices[0:lastIndx-1])...)
	withVn = append(withVn, vertices[lastIndx:]...)

	maxWithoutVn := e.SumWhoseVertices(withoutVn)
	maxWithVn := e.SumWhoseVertices(withVn)
	result := withoutVn

	if maxWithoutVn < maxWithVn {
		result = withVn
	}

	e.MwisResultCached[ArrUtility(vertices).StringJoined()] = result

	return result
}

func (e *MwisEvaluator) GetWeightAtVertex(vertex int) int {
	return e.Weights[vertex-1]
}

func (e *MwisEvaluator) MwisBegin() []int {
	maxWeightIndependentSetResult := e.Mwis(e.Vertices)
	return maxWeightIndependentSetResult
}

func (e *MwisEvaluator) SumWhoseVertices(vertices []int) int {
	summed := 0
	for i := 0; i < len(vertices); i++ {
		summed += e.GetWeightAtVertex(vertices[i])
	}
	return summed
}

func ReadTextfile(filepath string) MwisEvaluator {
	contentBytes, _ := ioutil.ReadFile(filepath)
	mwisEvaluator := MwisEvaluator{
		MwisResultCached: make(map[string][]int),
	}

	for lineIndx, intStr := range strings.Split(string(contentBytes), "\n") {
		if mwisEvaluator.Weights == nil {
			numLen, _ := strconv.Atoi(intStr)
			mwisEvaluator.Weights = make([]int, numLen)
			mwisEvaluator.Vertices = make([]int, numLen)
			continue
		}

		if intStr == "" {
			continue
		}

		num, _ := strconv.Atoi(intStr)
		mwisEvaluator.Weights[lineIndx-1] = num
		mwisEvaluator.Vertices[lineIndx-1] = lineIndx
	}

	return mwisEvaluator
}

type ArrUtility []int

func (a ArrUtility) StringJoined() string {
	arr := ([]int)(a)
	strs := fmt.Sprintf("%d", arr[0])
	for i := 1; i < len(arr); i++ {
		strs = fmt.Sprintf("%s,%d", strs, arr[i])
	}
	return strs
}
