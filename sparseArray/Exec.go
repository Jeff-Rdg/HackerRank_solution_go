package sparseArray

import (
	"fmt"
	"strings"
)

func Exec() {
	arr1 := []string{"aba", "baba", "aba", "xzxb"}
	arr2 := []string{"aba", "xzxb", "ab"}

	result := matchingStrings(arr1, arr2)
	fmt.Println(result)
}
func matchingStrings(arr []string, queries []string) []int32 {
	var response []int32
	var result int32
	for _, q := range queries {
		for _, a := range arr {
			if strings.EqualFold(a, q) {
				result++
			}
		}
		response = append(response, result)
		result = 0
	}
	return response
}
