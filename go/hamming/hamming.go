package hamming

import (
	"errors"
)
//Function for calculating Hamming Distance ....
func Distance(a, b string) (int, error) {
	//check if length of a and b are euqal if not return error
	if len(a) != len(b) {
		return 0,errors.New("length of DNA strand not equal")
	}
	//declaring out of block variables
	var HamDistance int
	//for loop over range of a and compare a && b ----->if not euqual increase counter
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			HamDistance ++
		}
	}
	//return vaue and error == nil
	return HamDistance, nil
}
