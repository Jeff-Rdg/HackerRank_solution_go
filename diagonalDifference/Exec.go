package diagonalDifference

import (
	"fmt"
	"math"
)

func Exec() {
	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	fmt.Println(diagonalDifference(arr))

}

func diagonalDifference(arr [][]int32) int32 {
	first := FirstDiagonal(arr)
	second := SecondDiagonal(arr)

	result := first - second
	resultAbs := math.Abs(float64(result))
	return int32(resultAbs)

}

func FirstDiagonal(arr [][]int32) int32 {
	var lastDiagonal int32
	for i, line := range arr {
		for j, value := range line {
			if i == j {
				lastDiagonal += value
			}
		}
	}
	return lastDiagonal
}

func SecondDiagonal(arr [][]int32) int32 {
	var lastDiagonal int32
	index := len(arr) - 1
	for _, line := range arr {
		for j, value := range line {
			if j == index {
				lastDiagonal += value
			}
		}
		index--
	}
	return lastDiagonal
}
