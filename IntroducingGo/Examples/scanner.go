package main

import "fmt"

func main(){
	fmt.Println(doubleNumber(readInput()))
}

func readInput() float32{
	fmt.Println("Enter a number:")
	var input float32
	fmt.Scanf("%f", &input)
	return input
}

func doubleNumber(x float32) float32{
	return x * 2
}