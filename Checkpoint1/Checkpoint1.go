package main

import (
	"fmt"
	"math"
)

func main(){
	const s1 string = "declared a string"
	s2:="also declared a string"
	fmt.Println(s1)
	fmt.Println(s2)

	tempvar:=5/5
	fmt.Println(tempvar)

	fmt.Println(math.Sin(1))


	for tempvar=0;tempvar<=20;tempvar++{
		fmt.Println(tempvar)
	}

	var tempvar1,tempvar2,tempvar3 int = 0,1,0
	for tempvar1<=100{
		tempvar1=tempvar2+tempvar3
		fmt.Println(tempvar1)
		tempvar3=tempvar2
		tempvar2=tempvar1
	}

}