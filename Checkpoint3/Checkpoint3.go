package main

import "fmt"

func main(){
	fmt.Println("covering functions, if statements")
	output:= factorial(10)
	fmt.Println("10! =  ",output)

	fmt.Println("covering multiple return values")
	addition, subtraction, multiplication, division:=calculator(2,2) 
	fmt.Println("addition", addition)
	fmt.Println("subtraction", subtraction)
	fmt.Println("multiplication", multiplication)
	fmt.Println("division", division)

	fmt.Println("covering maps")
	fmt.Println("frequency ", frequency())
	fmt.Println("frequency hot", frequency()["hot"])
	fmt.Println("frequency meme", frequency()["meme"])

	fmt.Println("covering range")
	rangeFunction()
}


func factorial(input int)(int){
	if input==1 {
		return input
	}

	return (input*factorial(input-1))
}


func calculator(a,b int)(int,int,int,int){
	return a+b,a-b,a*b,a/b
}


func frequency()map[string]int{
	m:=make(map[string]int)

	m["meme"]=5
	m["word"]=2
	m["hot"]=3

	return m
}



func rangeFunction(){
	var a = "a string of letters"
	fmt.Println("index 5: ", a[5])
	
	for i,v :=range a{
		fmt.Println("index: ", i)
		fmt.Println("value: ", v, "	",string(v) )
	}
}
