package timeConversion

import (
	"fmt"
	"strconv"
	"strings"
)

func Exec() {
	time := "12:05:45PM"
	result := TimeConversion(time)
	fmt.Println(result)
}

func TimeConversion(s string) string {
	var result string
	var arr []string
	var firstNumber string

	if strings.Contains(s, "PM") {
		result = strings.ReplaceAll(s, "PM", "")
		arr = strings.Split(result, ":")
		firstNumber = arr[0]
		if !strings.Contains(firstNumber, "12") {
			number, _ := strconv.Atoi(firstNumber)
			number += 12
			firstNumber = strconv.Itoa(number)
		}
	} else {
		result = strings.ReplaceAll(s, "AM", "")
		arr = strings.Split(result, ":")
		firstNumber = arr[0]

		if strings.Contains(firstNumber, "12") {
			firstNumber = "00"
		}
	}
	result = fmt.Sprintf("%v:%v:%v", firstNumber, arr[1], arr[2])

	return result
}
