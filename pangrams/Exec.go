package pangrams

import (
	"fmt"
	"strings"
)

func Exec() {
	s := "We promptly judged antique ivory buckles for the next prize"
	fmt.Println(Pangrams(s))
}

func Pangrams(s string) string {
	s = strings.ToLower(s)
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "x", "y", "z"}
	var result bool

	for _, value := range alphabet {
		if !strings.Contains(s, value) {
			result = false
			break
		} else {
			result = true
		}
	}

	if result {
		return "pangram"
	} else {
		return "not pangram"
	}
}
