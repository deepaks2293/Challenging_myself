package main

import (
	"fmt"
	"strings"
)

func main() {
	f := toWeirdCase("This is a test Looks like you passed")
	fmt.Println(f)
}
func toWeirdCase(str string) string {
	// Your code here and happy coding!
	var EvenOddString string
	NewStr := strings.SplitAfter(str, " ")
	for i := 0; i < len(NewStr); i++ {
		for x, StrNew := range NewStr[i] {
			if x%2 == 0 {
				EvenOddString += strings.ToUpper(string(StrNew))
			} else {
				EvenOddString += strings.ToLower(string(StrNew))
			}
		}
	}
	return EvenOddString
}
