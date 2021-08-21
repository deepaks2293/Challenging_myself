package twofer

func ShareWith(name string) string {
	var x string    //Defining variable result
	if name != "" { //checking if name is not empty
		x = name
	} else {
		x = "you"
	}
	result := "One for " + x + ", one for me."
	return result
}
