package main
import (
	"fmt"
)
func main(){
	f:= Solution("abc")
	fmt.Println(f)
}

func Solution(str string) []string {
	if len(str)%2 == 0{
		str = str
	}else{
		str = str+"_"
	}
	var Chrac []string
	x:= 0
	for i:=2;i<=len(str);i+=2{
		Chrac=append(Chrac,string(str[x:i]))
		x =i 
	}
	return Chrac
}
