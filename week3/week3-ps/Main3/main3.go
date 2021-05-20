package Main3

import (
	"../MwisEvaluator"
	"fmt"
)

func Main3() {
	arrCheck := []int{1, 2, 3, 4, 17, 117, 517, 997}

	//mwisEvaluator := MwisEvaluator.ReadTextfile("test_mwis1.txt")
	//mwisEvaluator := MwisEvaluator.ReadTextfile("test_mwis2.txt")
	mwisEvaluator := MwisEvaluator.ReadTextfile("_790eb8b186eefb5b63d0bf38b5096873_mwis.txt")
	arrVerticesMwis := mwisEvaluator.MwisBegin()

	strs := ""
	for i := 0; i < len(arrCheck); i++ {
		add := "0"
		for j := 0; j < len(arrVerticesMwis); j++ {
			if arrVerticesMwis[j] == arrCheck[i] {
				add = "1"
				break
			}
		}
		strs = fmt.Sprintf("%s%s", strs, add)
	}

	fmt.Printf("%s\n", strs)
}
