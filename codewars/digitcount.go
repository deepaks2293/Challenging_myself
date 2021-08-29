package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	f := NbDig(25, 1)
	fmt.Println(f)
}
func NbDig(n int, d int) int {
	var squ string
	for i := 0; i <= n; i++ {
		squ += strconv.Itoa(i*i) + " "
	}
	var count int
	count = strings.Count(squ, strconv.Itoa(d))
	return count
}
