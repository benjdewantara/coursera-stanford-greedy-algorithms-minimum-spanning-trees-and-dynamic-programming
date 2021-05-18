package Main3

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

func Main3() {
    arr := ReadTextfile("_790eb8b186eefb5b63d0bf38b5096873_mwis.txt")

    fmt.Printf("%d\n", arr[0])
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
