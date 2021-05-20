package Main3

import (
	"../MwisEvaluator"
	"fmt"
)

func Main3() {
	mwisEvaluator := MwisEvaluator.MwisEvaluator{
		MwisResultCached: make(map[string][]int),
	}

	arrCheck := []int{1, 2, 3, 4, 17, 117, 517, 997}

	//arr := ReadTextfile("test_mwis1.txt")
	//arr := ReadTextfile("test_mwis2.txt")
	arr := MwisEvaluator.ReadTextfile("_790eb8b186eefb5b63d0bf38b5096873_mwis.txt")
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
