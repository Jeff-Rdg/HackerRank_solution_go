package main

import (
	"fmt"
	"github.com/jeff-rdg/hackerRankSolution/countSort1"
	"github.com/jeff-rdg/hackerRankSolution/diagonalDifference"
	"github.com/jeff-rdg/hackerRankSolution/flippingBits"
	"github.com/jeff-rdg/hackerRankSolution/lonelyInteger"
	"github.com/jeff-rdg/hackerRankSolution/pangrams"
	"github.com/jeff-rdg/hackerRankSolution/sparseArray"
	"github.com/jeff-rdg/hackerRankSolution/timeConversion"
)

func main() {
	fmt.Println("Informe o método: ")
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}
	SelectMethod(input)
}

func SelectMethod(value int) {
	switch value {
	case 1:
		timeConversion.Exec()
	case 2:
		sparseArray.Exec()
	case 3:
		lonelyInteger.Exec()
	case 4:
		flippingBits.Exec()
	case 5:
		diagonalDifference.Exec()
	case 6:
		countSort1.Exec()
	case 7:
		pangrams.Exec()

	}
}
