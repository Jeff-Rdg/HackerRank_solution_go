package flippingBits

import (
	"fmt"
	"strconv"
	"strings"
)

func Exec() {
	var a int64
	a = 2147483647

	fmt.Println(flippingBits(a))
}

func flippingBits(n int64) int64 {
	binary := strconv.FormatUint(uint64(n), 2)
	binary = fmt.Sprintf("%032s", binary)
	var arr []string

	for _, i := range binary {
		if string(i) == "0" {
			arr = append(arr, "1")
			continue
		}
		if string(i) == "1" {
			arr = append(arr, "0")
		}
	}

	strResult := strings.Join(arr, "")
	result, _ := strconv.ParseInt(strResult, 2, 64)
	return result
}
