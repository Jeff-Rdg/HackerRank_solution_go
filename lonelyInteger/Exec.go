package lonelyInteger

import "fmt"

func Exec() {
	a := []int32{1, 2, 3, 4, 3, 2, 1}
	fmt.Println(LonelyInteger(a))
}

func LonelyInteger(a []int32) int32 {
	var count int
	var result int32
	for _, column := range a {
		for _, row := range a {
			if column == row {
				count++
			}
		}
		if count == 1 {
			result = column
			break
		}
		count = 0
	}
	return result
}
