package hamming

import (
	"fmt"
)

func Distance(a, b string) (int, error) {
	//check if length of a and b are euqal if not return error
	if len(a) != len(b) {
		return 0, fmt.Errorf("length of DNA strand not equal")
	}
	//declaring out of block variables
	var Ham_Distance int
	//for loop over range of a and compare a && b ----->if not euqual increase counter
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			Ham_Distance += 1
		}
	}
	//return vaue and error == nil
	return Ham_Distance, nil
}
