package Main3

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Main3() {
	mwisEvaluator := MwisEvaluator{
		MwisResultCached: make(map[string][]int),
	}

	arrCheck := []int{1, 2, 3, 4, 17, 117, 517, 997}

	//arr := ReadTextfile("test_mwis1.txt")
	//arr := ReadTextfile("test_mwis2.txt")
	arr := ReadTextfile("_790eb8b186eefb5b63d0bf38b5096873_mwis.txt")
	arrMwis := mwisEvaluator.Mwis(arr)

	strs := ""

	for j := 0; j < len(arrCheck); j++ {
		add := ""

		vertex := j + 1
		for i := 0; i < len(arrMwis); i++ {
			if arrCheck[j] == vertex {
				add = "1"
				break
			}
		}

		if add == "" {
			add = "0"
		}

		strs = fmt.Sprintf("%s%s", strs, add)
	}

	fmt.Printf("%s\n", strs)
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

type MwisEvaluator struct {
	MwisResultCached map[string][]int
}

func (e *MwisEvaluator) Mwis(arr []int) []int {
	cachedResult := e.MwisResultCached[ArrUtility(arr).StringJoined()]
	if cachedResult != nil {
		return cachedResult
	}

	if len(arr) == 2 {
		if arr[0] < arr[1] {
			return arr[1:2]
		} else {
			return arr[0:1]
		}
	} else if len(arr) == 1 {
		return arr[0:1]
	}

	lastIndx := len(arr) - 1
	withoutVn := make([]int, 0)
	withoutVn = append(withoutVn, e.Mwis(arr[0:lastIndx])...)

	withVn := make([]int, 0)
	withVn = append(withVn, e.Mwis(arr[0:lastIndx-1])...)
	withVn = append(withVn, arr[lastIndx:]...)

	maxWithoutVn := Sum(withoutVn)
	maxWithVn := Sum(withVn)
	result := withoutVn

	if maxWithoutVn < maxWithVn {
		result = withVn
	}

	e.MwisResultCached[ArrUtility(arr).StringJoined()] = result

	return result
}

func Sum(arr []int) int {
	summed := 0
	for i := 0; i < len(arr); i++ {
		summed += arr[i]
	}
	return summed
}

func ReadTextfile(filepath string) []int {
	contentBytes, _ := ioutil.ReadFile(filepath)
	var arr []int

	for lineIndx, intStr := range strings.Split(string(contentBytes), "\n") {
		if arr == nil {
			numLen, _ := strconv.Atoi(intStr)
			arr = make([]int, numLen)
			continue
		}

		if intStr == "" {
			continue
		}

		num, _ := strconv.Atoi(intStr)
		arr[lineIndx-1] = num
	}

	return arr
}
